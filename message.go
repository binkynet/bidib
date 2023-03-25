package bidib

import (
	"fmt"
	"strconv"
)

// Abstract interface for Bidib messages
type Message interface {
	// Encode this message
	Encode(write func(uint8), seqNum SequenceNumber)
	String() string
}

// Type of message
type MessageType uint8

// String returns a human readable representation of a MessageType.
func (mt MessageType) String() string {
	return fmt.Sprintf("0x%02x", uint8(mt))
}

// Message sequence number
type SequenceNumber uint8

// String returns a human readable representation of a MessageType.
func (sn SequenceNumber) String() string {
	return strconv.Itoa(int(sn))
}

// Next returns the next sequence number.
func (sn SequenceNumber) Next() SequenceNumber {
	if sn == 255 {
		return 1
	}
	return sn + 1
}

// Reset the sequence number to 0.
func (sn *SequenceNumber) Reset() {
	*sn = 0
}

// MessageProcessor takes decoded raw message information and processes it.
type MessageProcessor func(MessageType, Address, SequenceNumber, []byte)

// SplitPackageAndProcessMessages splits a stream of message bytes into individual messages
// and processes those.
func SplitPackageAndProcessMessages(src []byte, proc MessageProcessor) error {
	for {
		// Decode first message
		mType, addr, seqNum, data, remaining, err := decodeMessage(src)
		if err != nil {
			return err
		}
		// Process message
		proc(mType, addr, seqNum, data)
		// Continue with next message (if any)
		if len(remaining) == 0 {
			return nil
		}
		src = remaining
	}
}

// decodeMessage decodes the given byte slice into the its raw message parts.
// Returns: type, address, seqNum, data, remaining, error
func decodeMessage(src []byte) (MessageType, Address, SequenceNumber, []byte, []byte, error) {
	// Fetch length (this excludes msgLength itself)
	msgLength := src[0]
	// Skip msgLength
	src = src[1:]
	// Check length
	if len(src) < int(msgLength) {
		return 0, Address{}, 0, nil, nil, fmt.Errorf("Invalid message length; got %d expected %d", len(src), msgLength)
	}
	remaining := src[msgLength:]

	// Fetch address
	addressIdx := 0
	var addr Address
	for {
		if src[0] == 0 {
			// We're done with address
			src = src[1:]
			msgLength--
			break
		}
		if addressIdx == 4 {
			// Address exists 4 bytes
			return 0, Address{}, 0, nil, nil, fmt.Errorf("Address exceeds 4 bytes")
		}
		addr[addressIdx] = src[0]
		addressIdx++
		src = src[1:]
		msgLength--
	}

	// Fetch message num
	seqNum := SequenceNumber(src[0])
	src = src[1:]
	msgLength--

	// Fetch message type
	mType := MessageType(src[0])
	src = src[1:]
	msgLength--

	// Fetch data
	data := src[:msgLength]

	return mType, addr, seqNum, data, remaining, nil
}

// Encode this message to the given writer.
func EncodeMessage(write func(uint8), mType MessageType, addr Address, seqNum SequenceNumber, data []byte) {
	addrLen := addr.GetLength()
	dataLen := len(data)
	// MsgLength
	msgLength := 1 /*type*/ + (addrLen + 1) + 1 /*msgNum*/ + uint8(dataLen)
	write(msgLength)
	// MsgAddr
	for i := 0; i < int(addrLen); i++ {
		write(addr[i])
	}
	write(0) // Terminating address
	// MsgNum
	write(uint8(seqNum))
	// MsgType
	write(uint8(mType))
	// Data
	for i := 0; i < dataLen; i++ {
		write(data[i])
	}
}
