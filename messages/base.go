package messages

import (
	"github.com/binkynet/bidib"
)

// BaseMessage is the basis for every message, containing an address.
type BaseMessage struct {
	Address bidib.Address
}

// GetAddress returns the address of the message
func (m BaseMessage) GetAddress() bidib.Address {
	return m.Address
}

// String converts an address into a readable string
func (m BaseMessage) String() string {
	return m.Address.String()
}
