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

// Current in mA
func (m BstDiag) Current() string {
	i := m.DiagI
	if i == 0 {
		return "0 mA"
	}
	if i >= 1 && i <= 15 {
		return fmt.Sprintf("%d mA", int(i))
	}
	if i >= 16 && i <= 63 {
		return fmt.Sprintf("%d mA", int(i-12)*4)
	}
	if i >= 64 && i <= 127 {
		return fmt.Sprintf("%d mA", int(i-51)*16)
	}
	if i >= 128 && i <= 191 {
		return fmt.Sprintf("%d mA", int(i-108)*64)
	}
	if i >= 192 && i <= 250 {
		return fmt.Sprintf("%d mA", int(i-171)*256)
	}
	if i >= 251 && i <= 253 {
		return "reserved"
	}
	if i == 254 {
		return "overcurrent"
	}
	return "unknown"
}

// Voltage
func (m BstDiag) Voltage() string {
	v := m.DiagV
	if v >= 0 && v <= 250 {
		return fmt.Sprintf("%d mV", int(v)*100)
	}
	if v >= 251 && v <= 254 {
		return "reserved"
	}
	return "unknown"
}

// Temperature
func (m BstDiag) Temperature() string {
	t := m.DiagTemp
	if t >= 0 && t <= 127 {
		return fmt.Sprintf("%d C", int(t))
	}
	if t >= 128 && t <= 225 {
		return "reserved"
	}
	return fmt.Sprintf("neg %d C", int(t))
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
