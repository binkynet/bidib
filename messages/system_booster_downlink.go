package messages

import (
	"fmt"

	"github.com/binkynet/bidib"
)

// BoostOn
type BoostOn struct {
	BaseMessage
	CurrentNodeOnly bool
}

func (m BoostOn) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{0}
	if m.CurrentNodeOnly {
		data[0] = 1
	}
	bidib.EncodeMessage(write, bidib.MSG_BOOST_ON, m.Address, seqNum, data)
}

func (m BoostOn) String() string {
	return fmt.Sprintf("%T addr=%s current_node_only=%v", m, m.Address, m.CurrentNodeOnly)
}

func decodeBoostOn(addr bidib.Address, data []byte) (BoostOn, error) {
	var result BoostOn
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.CurrentNodeOnly = data[0] != 0
	return result, nil
}

// BoostOff
type BoostOff struct {
	BaseMessage
	CurrentNodeOnly bool
}

func (m BoostOff) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{0}
	if m.CurrentNodeOnly {
		data[0] = 1
	}
	bidib.EncodeMessage(write, bidib.MSG_BOOST_OFF, m.Address, seqNum, data)
}

func (m BoostOff) String() string {
	return fmt.Sprintf("%T addr=%s current_node_only=%v", m, m.Address, m.CurrentNodeOnly)
}

func decodeBoostOff(addr bidib.Address, data []byte) (BoostOff, error) {
	var result BoostOff
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.CurrentNodeOnly = data[0] != 0
	return result, nil
}
