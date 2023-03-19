package messages

import (
	"fmt"

	"github.com/binkynet/bidib"
)

// Followed by a byte with content 0. (= local bus address of the host)
// The track-output node does not receive commands from any other local addresses.
// This lock is valid for 2 seconds and then expires by itself.
type CsAllocate struct {
	BaseMessage
}

func (m CsAllocate) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{0}
	bidib.EncodeMessage(write, bidib.MSG_CS_ALLOCATE, m.Address, seqNum, data)
}

func (m CsAllocate) String() string {
	return fmt.Sprintf("%T addr=%s", m, m.Address)
}

func decodeCsAllocate(addr bidib.Address, data []byte) (CsAllocate, error) {
	var result CsAllocate
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	// Data is ignored
	return result, nil
}

// With this command, the state of the track output is set or queried. Followed by one byte which encodes the new state.
// The node responds with a MSG_CS_STATE.
// Before first turn on, the speed settings of all locos should be checked resp. the loco stack should be cleared.
// This avoid unintentional driving of loco with speed settings from a previous session.
type CsSetState struct {
	BaseMessage
	State bidib.CsState
}

func (m CsSetState) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{uint8(m.State)}
	bidib.EncodeMessage(write, bidib.MSG_CS_SET_STATE, m.Address, seqNum, data)
}

func (m CsSetState) String() string {
	return fmt.Sprintf("%T addr=%s state=0x%02x", m, m.Address, m.State)
}

func decodeCsSetState(addr bidib.Address, data []byte) (CsSetState, error) {
	var result CsSetState
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.State = bidib.CsState(data[0])
	return result, nil
}

// Motion commands will be issued with this command. Following by further parammeters,
// which contains format and output functions. If an loco has no higher functions, the appropriate
// groups of the function commands should be marked as inactive. This conserves the limited bandwidth at the track.
// MSG_CS_DRIVE commands will be acknowledged by one or more MSG_CS_DRIVE_ACK messages.
// The various acknowledgement levels can be activated in the output unit via FEATURE_GEN_DRIVE_ACK.
// If multiple commands for the same DCC address are issued, the output device is allowed to combine them in case of low bandwidth.
// In this case, intermediate acknowledgements can be omitted.
// Motion commands will be passed always with the step number of 127 + direction bits.
// Only the track-output unit converts this motion command to the appropriate speed-step on the track depending from the selected format.
type CsDrive struct {
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

func (m CsDrive) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
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
	bidib.EncodeMessage(write, bidib.MSG_CS_DRIVE, m.Address, seqNum, data[:])
}

func (m CsDrive) String() string {
	return fmt.Sprintf("%T addr=%s dccAddr=0x%04x speed=%d forward=%t", m, m.Address, m.DccAddress, m.Speed, m.DirectionForward)
}

func decodeCsDrive(addr bidib.Address, data []byte) (CsDrive, error) {
	var result CsDrive
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
