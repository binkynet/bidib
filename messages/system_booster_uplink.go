package messages

import (
	"fmt"

	"github.com/binkynet/bidib"
)

// Booster state
type BstState struct {
	BaseMessage
	State bidib.BstState
}

func (m BstState) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{uint8(m.State)}
	bidib.EncodeMessage(write, bidib.MSG_BOOST_STAT, m.Address, seqNum, data)
}

func (m BstState) String() string {
	return fmt.Sprintf("%T addr=%s state=%s", m, m.Address, m.State)
}

func decodeBstState(addr bidib.Address, data []byte) (BstState, error) {
	var result BstState
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.State = bidib.BstState(data[0])
	return result, nil
}

// Booster diagnostics
type BstDiag struct {
	BaseMessage
	DiagI    byte
	DiagV    byte
	DiagTemp byte
}

const (
	BIDIB_BST_DIAG_I    = 0x00
	BIDIB_BST_DIAG_V    = 0x01
	BIDIB_BST_DIAG_TEMP = 0x02
)

func (m BstDiag) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{BIDIB_BST_DIAG_I, m.DiagI, BIDIB_BST_DIAG_V, m.DiagV, BIDIB_BST_DIAG_TEMP, m.DiagTemp}
	bidib.EncodeMessage(write, bidib.MSG_BOOST_DIAGNOSTIC, m.Address, seqNum, data)
}

func (m BstDiag) String() string {
	return fmt.Sprintf("%T addr=%s i=%02x v=%02x temp=%02x", m, m.Address, m.DiagI, m.DiagV, m.DiagTemp)
}

func decodeBstDiag(addr bidib.Address, data []byte) (BstDiag, error) {
	var result BstDiag
	result.Address = addr
	if len(data)%2 != 0 {
		return result, fmt.Errorf("invalid data length; got %d, expected multiple of 2", len(data))
	}
	for i := 0; i < len(data); i += 2 {
		switch data[i] {
		case BIDIB_BST_DIAG_I:
			result.DiagI = data[i+1]
		case BIDIB_BST_DIAG_V:
			result.DiagV = data[i+1]
		case BIDIB_BST_DIAG_TEMP:
			result.DiagTemp = data[i+1]
		default:
			return result, fmt.Errorf("unknown enum value %02x", data[i])
		}
	}
	return result, nil
}
