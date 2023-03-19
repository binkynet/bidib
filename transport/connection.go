package transport

import "github.com/binkynet/bidib"

// Connection is implemented by a specific transport type.
type Connection interface {
	// SendMessages encodes all given messages and sends them to the serial port.
	SendMessages(messages []bidib.Message, seqNum bidib.SequenceNumber) error
	// Close the connection
	Close() error
}
