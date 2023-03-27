package messages

import (
	"fmt"

	"github.com/binkynet/bidib"
)

// CV-message, followed by 5 bytes: ADDRL, ADDRH, CVL, CVH, DAT
type BmCv struct {
	BaseMessage
	DccAddress uint16
	Cv         uint16
	Data       uint8
}

func (m BmCv) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{0, 0, 0, 0, 0}
	writeUint16(data, m.DccAddress)
	writeUint16(data[2:], m.Cv)
	data[4] = m.Data
	bidib.EncodeMessage(write, bidib.MSG_BM_CV, m.Address, seqNum, data)
}

func (m BmCv) String() string {
	return fmt.Sprintf("%T addr=%s dccaddr=%d cv=%d data=%d", m, m.Address, m.DccAddress, m.Cv, m.Data)
}

func decodeBmCv(addr bidib.Address, data []byte) (BmCv, error) {
	var result BmCv
	if err := validateDataLength(data, 5); err != nil {
		return result, err
	}
	result.Address = addr
	result.DccAddress = readUint16(data)
	result.Cv = readUint16(data[2:])
	result.Data = data[4]
	return result, nil
}

// Speed-message, followed by 4 bytes: ADDRL, ADDRH, SPEEDL, SPEEDH
type BmSpeed struct {
	BaseMessage
	DccAddress uint16
	Speed      uint16
}

func (m BmSpeed) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{0, 0, 0, 0}
	writeUint16(data, m.DccAddress)
	writeUint16(data[2:], m.Speed)
	bidib.EncodeMessage(write, bidib.MSG_BM_SPEED, m.Address, seqNum, data)
}

func (m BmSpeed) String() string {
	return fmt.Sprintf("%T addr=%s dccaddr=%d speed=%d", m, m.Address, m.DccAddress, m.Speed)
}

func decodeBmSpeed(addr bidib.Address, data []byte) (BmSpeed, error) {
	var result BmSpeed
	if err := validateDataLength(data, 4); err != nil {
		return result, err
	}
	result.Address = addr
	result.DccAddress = readUint16(data)
	result.Speed = readUint16(data[2:])
	return result, nil
}

// Status message of an decoder followed by 5 or more bytes, beginning with MNUM, ADDRL, ADDRH, DYN_NUM.
type BmDynState struct {
	BaseMessage
	MNum       uint8
	DccAddress uint16
	DynNum     uint8
	Value      uint8
}

func (m BmDynState) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{m.MNum, 0, 0, m.DynNum, m.Value}
	writeUint16(data[1:], m.DccAddress)
	bidib.EncodeMessage(write, bidib.MSG_BM_DYN_STATE, m.Address, seqNum, data)
}

func (m BmDynState) String() string {
	return fmt.Sprintf("%T addr=%s dccaddr=%d mnum=%d dynnum=%d value=%d", m, m.Address, m.DccAddress, m.MNum, m.DynNum, m.Value)
}

func decodeBmDynState(addr bidib.Address, data []byte) (BmDynState, error) {
	var result BmDynState
	if err := validateMinDataLength(data, 5); err != nil {
		return result, err
	}
	result.Address = addr
	result.MNum = data[0]
	result.DccAddress = readUint16(data[1:])
	result.DynNum = data[3]
	result.Value = data[4]
	return result, nil
}
