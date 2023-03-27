package messages

import (
	"fmt"

	"github.com/binkynet/bidib"
)

// With this message, the current state of the output is reported, followed by one byte
// which encodes the current state. The state is encoded similar to MSG_CS_SET_STATE.
type CsState struct {
	BaseMessage
	State bidib.CsState
}

func (m CsState) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{uint8(m.State)}
	bidib.EncodeMessage(write, bidib.MSG_CS_STATE, m.Address, seqNum, data)
}

func (m CsState) String() string {
	return fmt.Sprintf("%T addr=%s state=%s", m, m.Address, m.State)
}

func decodeCsState(addr bidib.Address, data []byte) (CsState, error) {
	var result CsState
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.State = bidib.CsState(data[0])
	return result, nil
}

// Motion commands will be acknowledged with this command. Followed by further parameters:
type CsDriveAck struct {
	BaseMessage
	DccAddress uint16
	Ack        bidib.CsAck
}

func (m CsDriveAck) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{0, 0, byte(m.Ack)}
	writeUint16(data, m.DccAddress)
	bidib.EncodeMessage(write, bidib.MSG_CS_DRIVE_ACK, m.Address, seqNum, data)
}

func (m CsDriveAck) String() string {
	return fmt.Sprintf("%T addr=%s dccaddr=%d ack=%d", m, m.Address, m.DccAddress, m.Ack)
}

func decodeCsDriveAck(addr bidib.Address, data []byte) (CsDriveAck, error) {
	var result CsDriveAck
	if err := validateDataLength(data, 3); err != nil {
		return result, err
	}
	result.Address = addr
	result.DccAddress = readUint16(data)
	result.Ack = bidib.CsAck(data[2])
	return result, nil
}

// Switching commands will be acknowledged with this command. Followed by further parameters:
type CsAccessoryAck struct {
	BaseMessage
	DccAddress uint16
	Ack        bidib.CsAck
}

func (m CsAccessoryAck) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{0, 0, byte(m.Ack)}
	writeUint16(data, m.DccAddress)
	bidib.EncodeMessage(write, bidib.MSG_CS_ACCESSORY_ACK, m.Address, seqNum, data)
}

func (m CsAccessoryAck) String() string {
	return fmt.Sprintf("%T addr=%s dccaddr=%d ack=%d", m, m.Address, m.DccAddress, m.Ack)
}

func decodeCsAccessoryAck(addr bidib.Address, data []byte) (CsAccessoryAck, error) {
	var result CsAccessoryAck
	if err := validateDataLength(data, 3); err != nil {
		return result, err
	}
	result.Address = addr
	result.DccAddress = readUint16(data)
	result.Ack = bidib.CsAck(data[2])
	return result, nil
}

// POM commands will be acknowledged with this command. Followed by further parameters,
// address and acknowledge.
// Address is coded identically to the MSG_CS_POM command and will be just 'passed through' at the node.
type CsPomAck struct {
	BaseMessage
	DccAddress uint32
	Mid        uint8
	Ack        bidib.CsAck
}

func (m CsPomAck) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{0, 0, 0, 0, m.Mid, byte(m.Ack)}
	writeUint32(data, m.DccAddress)
	bidib.EncodeMessage(write, bidib.MSG_CS_POM_ACK, m.Address, seqNum, data)
}

func (m CsPomAck) String() string {
	return fmt.Sprintf("%T addr=%s dccaddr=%d mid=%d ack=%d", m, m.Address, m.DccAddress, m.Mid, m.Ack)
}

func decodeCsPomAck(addr bidib.Address, data []byte) (CsPomAck, error) {
	var result CsPomAck
	if err := validateDataLength(data, 6); err != nil {
		return result, err
	}
	result.Address = addr
	result.DccAddress = readUint32(data)
	result.Mid = data[4]
	result.Ack = bidib.CsAck(data[5])
	return result, nil
}

// Manual operation (via handheld) of an loco will be reported with this command.
// For this purpose, the feature FEATURE_GEN_NOTIFY_DRIVE_MANUAL has be set to 1.
// Followed by further parameters with the same structure like MSG_CS_DRIVE:
type CsDriveManual struct {
	BaseMessage
	DccAddress       uint16
	DccFormat        bidib.DccFormat
	OutputSpeed      bool
	OutputF1_F4      bool
	OutputF5_F8      bool
	OutputF9_F12     bool
	OutputF13_F20    bool
	OutputF21_F28    bool
	DirectionForward bool
	Speed            uint8
	Flags            bidib.DccFlags
}

func (m CsDriveManual) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := [9]byte{}
	writeUint16(data[0:], m.DccAddress)
	data[2] = byte(m.DccFormat)
	if m.OutputSpeed {
		data[3] |= 0x01
	}
	if m.OutputF1_F4 {
		data[3] |= 0x02
	}
	if m.OutputF5_F8 {
		data[3] |= 0x04
	}
	if m.OutputF9_F12 {
		data[3] |= 0x08
	}
	if m.OutputF13_F20 {
		data[3] |= 0x10
	}
	if m.OutputF21_F28 {
		data[3] |= 0x20
	}
	data[4] = m.Speed
	if m.DirectionForward {
		data[4] |= 0x80
	}
	data[5] = m.Flags.GenerateBits(1, 4) | (m.Flags.GenerateBits(0, 0) << 4)
	data[6] = m.Flags.GenerateBits(5, 12)
	data[7] = m.Flags.GenerateBits(13, 20)
	data[8] = m.Flags.GenerateBits(21, 28)
	bidib.EncodeMessage(write, bidib.MSG_CS_DRIVE_MANUAL, m.Address, seqNum, data[:])
}

func (m CsDriveManual) String() string {
	return fmt.Sprintf("%T addr=%s dccAddr=0x%04x speed=%d forward=%t", m, m.Address, m.DccAddress, m.Speed, m.DirectionForward)
}

func decodeCsDriveManual(addr bidib.Address, data []byte) (CsDriveManual, error) {
	var result CsDriveManual
	if err := validateDataLength(data, 9); err != nil {
		return result, err
	}
	result.Address = addr
	result.DccAddress = readUint16(data)
	result.DccFormat = bidib.DccFormat(data[2])
	result.OutputSpeed = (data[3] & 0x01) != 0
	result.OutputF1_F4 = (data[3] & 0x02) != 0
	result.OutputF5_F8 = (data[3] & 0x04) != 0
	result.OutputF9_F12 = (data[3] & 0x08) != 0
	result.OutputF13_F20 = (data[3] & 0x10) != 0
	result.OutputF21_F28 = (data[3] & 0x20) != 0
	result.Speed = data[4] & 0b01111111
	result.DirectionForward = (data[4] & 0x80) != 0
	result.Flags.Set(0, (data[5]&0x10) != 0)
	result.Flags.SetBits(1, 4, data[5])
	result.Flags.SetBits(5, 12, data[6])
	result.Flags.SetBits(13, 20, data[7])
	result.Flags.SetBits(21, 28, data[8])
	return result, nil
}

// This command reports the events in a loco. Followed by further parameters:
type CsDriveEvent struct {
	BaseMessage
	DccAddress uint16
	Event      bidib.CsEvent
}

func (m CsDriveEvent) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{0, 0, byte(m.Event)}
	writeUint16(data, m.DccAddress)
	bidib.EncodeMessage(write, bidib.MSG_CS_DRIVE_EVENT, m.Address, seqNum, data)
}

func (m CsDriveEvent) String() string {
	return fmt.Sprintf("%T addr=%s dccaddr=%d event=%d", m, m.Address, m.DccAddress, m.Event)
}

func decodeCsDriveEvent(addr bidib.Address, data []byte) (CsDriveEvent, error) {
	var result CsDriveEvent
	if err := validateDataLength(data, 3); err != nil {
		return result, err
	}
	result.Address = addr
	result.DccAddress = readUint16(data)
	result.Event = bidib.CsEvent(data[2])
	return result, nil
}

// This command reports results from an programming operation in service mode. Followed by further parameters:
type CsProgState struct {
	BaseMessage
	State uint8
	Time  uint8
	Cv    uint16
	Data  uint8
}

func (m CsProgState) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{m.State, m.Time, 0, 0, m.Data}
	writeUint16(data[2:], m.Cv)
	bidib.EncodeMessage(write, bidib.MSG_CS_PROG_STATE, m.Address, seqNum, data)
}

func (m CsProgState) String() string {
	return fmt.Sprintf("%T addr=%s state=0x%02x time=%d cv=%d data=%d", m, m.Address, m.State, m.Time, m.Cv, m.Data)
}

func decodeCsProgState(addr bidib.Address, data []byte) (CsProgState, error) {
	var result CsProgState
	if err := validateMinDataLength(data, 4); err != nil {
		return result, err
	}
	result.Address = addr
	result.State = data[0]
	result.Time = data[1]
	result.Cv = readUint16(data[2:])
	if len(data) > 4 {
		result.Data = data[4]
	}
	return result, nil
}
