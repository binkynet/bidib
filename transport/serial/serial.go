package serial

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/rs/zerolog"
	ts "github.com/tarm/serial"

	"github.com/binkynet/bidib"
	"github.com/binkynet/bidib/messages"
	"github.com/binkynet/bidib/transport"
)

// Serial port config
type Config struct {
	PortName string
}

// New constructs and opens a new serial port transport.
func New(cfg Config, log zerolog.Logger, processor bidib.MessageProcessor) (transport.Connection, error) {
	sc := &serialConnection{
		cfg:       cfg,
		log:       log,
		processor: processor,
	}
	if err := sc.open(); err != nil {
		return nil, err
	}
	return sc, nil
}

// serialConnection implements serial port transport.
type serialConnection struct {
	cfg       Config
	log       zerolog.Logger
	processor bidib.MessageProcessor
	port      io.Closer
	reader    io.Reader
	writer    io.Writer
	read      struct {
		mutex  sync.Mutex
		buffer [256]byte
	}
	write struct {
		mutex  sync.Mutex
		buffer [256]byte
	}
	running bool
}

// open the serial port connection and
func (sc *serialConnection) open() error {
	c := ts.Config{
		Name:     sc.cfg.PortName,
		StopBits: ts.Stop1,
		Parity:   ts.ParityNone,
	}
	var lastError error
	sysGetMagic := messages.SysGetMagic{}
	for _, baud := range []int{1000000, 115200, 19200} {
		sc.log.Debug().
			Int("baud", baud).
			Str("portName", c.Name).
			Msg("Try to open port")
		c.Baud = baud
		port, err := ts.OpenPort(&c)
		if err != nil {
			lastError = err
			time.Sleep(time.Millisecond * 10)
			continue
		}
		// Port can be opened, try sending MSG_SYS_GET_MAGIC
		sc.port = port
		sc.reader = port
		sc.writer = port
		if err := sc.SendMessages([]bidib.Message{sysGetMagic}, 0); err != nil {
			port.Close()
			lastError = err
			time.Sleep(time.Millisecond * 10)
			continue
		}
		// We were able to send a message
		sc.log.Debug().Msg("SendMessages succeeded")
		go sc.run()
		return nil
	}
	return lastError
}

func (sc *serialConnection) Close() error {
	sc.running = false
	return nil
}

// Run the serial connection until the connection is closed.
func (sc *serialConnection) run() {
	sc.running = true
	for {
		if !sc.running {
			sc.port.Close()
			return
		}
		sc.receivePacket()
	}
}

// SendMessages encodes all given messages and sends them to the serial port.
func (sc *serialConnection) SendMessages(messages []bidib.Message, seqNum bidib.SequenceNumber) error {
	sc.write.mutex.Lock()
	defer sc.write.mutex.Unlock()

	// Encode messages
	bufferIndex := 0
	buffer := sc.read.buffer[:]
	crc := uint8(0)

	write := func(data uint8) {
		crc = bidibCrcArray[data^crc]
		if data == bidib.BIDIB_PKT_MAGIC {
			buffer[bufferIndex] = bidib.BIDIB_PKT_ESCAPE
			bufferIndex++
			buffer[bufferIndex] = data ^ 0x20
			bufferIndex++
		} else {
			buffer[bufferIndex] = data
			bufferIndex++
		}
	}

	// Start with magic
	buffer[bufferIndex] = bidib.BIDIB_PKT_MAGIC
	bufferIndex++

	// Encode messages
	for _, m := range messages {
		m.Encode(write, seqNum)
		seqNum++
	}

	// Add CRC
	buffer[bufferIndex] = crc
	bufferIndex++

	// Add closing MAGIC
	buffer[bufferIndex] = bidib.BIDIB_PKT_MAGIC
	bufferIndex++

	// Write buffer to serial port
	startIndex := 0
	for {
		n, err := sc.writer.Write(buffer[startIndex:bufferIndex])
		if err != nil {
			return err
		}
		startIndex += n
		if startIndex == bufferIndex {
			// We're done
			return nil
		}
	}
}

// receivePacket tries to receive one or more messages from the serial port
// Received messages are sent to the mesage processor.
func (sc *serialConnection) receivePacket() {
	sc.read.mutex.Lock()
	defer sc.read.mutex.Unlock()

	bufferIndex := 0
	buffer := sc.read.buffer[:]
	escapeHot := false
	crc := uint8(0)
	tmp := [1]byte{}

	// Read the packet bytes
	for sc.running {
		var data uint8
		for {
			if n, err := sc.reader.Read(tmp[:]); err == nil && n == 1 {
				// We read a byte
				data = tmp[0]
				break
			} else {
				sc.log.Trace().Err(err).Int("n", n).Msg("Read failed")
			}
			if !sc.running {
				return
			}
			time.Sleep(time.Microsecond * 5)
		}

		//sc.log.Trace().Uint8("data", data).Msg("read data")
		if data == bidib.BIDIB_PKT_MAGIC {
			if bufferIndex == 0 {
				continue
			}
			break
		} else if data == bidib.BIDIB_PKT_ESCAPE {
			// Next byte is escaped
			escapeHot = true
		} else {
			// Put byte in buffer
			if escapeHot {
				data ^= 0x20
				escapeHot = false
			}
			buffer[bufferIndex] = data
			crc = bidibCrcArray[buffer[bufferIndex]^crc]
			bufferIndex++
		}
	}

	if !sc.running {
		return
	}

	if crc == 0x00 {
		//sc.log.Trace().Msg("CRC correct, split packet in messages")
		// Split packet in messages and process them
		bufferIndex--
		if err := bidib.SplitPackageAndProcessMessages(buffer[:bufferIndex], sc.processor); err != nil {
			sc.log.Warn().Err(err).Msg("failed to split messages")
			return
		}
	} else {
		sc.log.Warn().
			Str("packet", fmt.Sprintf("%0x", buffer[:bufferIndex])).
			Msg("CRC wrong, packet ignored")
	}
}
