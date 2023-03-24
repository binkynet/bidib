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

// Accessory items will be controlled with this command. Followed by 4 bytes: ADDRL, ADDRH, DATA, TIME
// The accessory decoder with address [ADDRH * 256 + ADDRL] is driven by the term in DATA.
// DATA is a bit structure, constisting of CONFIG (Bit 7,6) ACTIVATE (Bit 5) and ASPECT (Bit 4 – Bit 0).
type CsAccessory struct {
	BaseMessage
	DccAddress           uint16
	Extended             bool // Extended=true, Normal=false
	OutputUnitDoesTiming bool
	Activate             bool
	Aspect               uint8 // Accessory decoder with 2 aspects will be controlled with (ASPECT 0 = red) and (ASPECT 1 = green).
	TimeUnitSec          bool  // Base unit: 1sec=true, 100ms=false
	TimeValue            uint8 // 0..127
}

func (m CsAccessory) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := [4]byte{}
	writeUint16(data[0:], m.DccAddress)
	if m.Extended {
		data[2] |= 0x80
	}
	if m.OutputUnitDoesTiming {
		data[2] |= 0x40
	}
	if m.Activate {
		data[2] |= 0x20
	}
	data[2] |= (m.Aspect & 0b00011111)
	if m.TimeUnitSec {
		data[3] |= 0x80
	}
	data[3] |= (m.TimeValue & 0b01111111)
	bidib.EncodeMessage(write, bidib.MSG_CS_ACCESSORY, m.Address, seqNum, data[:])
}

func (m CsAccessory) String() string {
	return fmt.Sprintf("%T addr=%s dccAddr=0x%04x extended=%t outputunitdoestiming=%t activate=%t aspect=0x%02x timeunitsec=%t timevalue=%d", m, m.Address, m.DccAddress, m.Extended, m.OutputUnitDoesTiming, m.Activate, m.Aspect, m.TimeUnitSec, m.TimeValue)
}

func decodeCsAccessory(addr bidib.Address, data []byte) (CsAccessory, error) {
	var result CsAccessory
	if err := validateDataLength(data, 4); err != nil {
		return result, err
	}
	result.Address = addr
	result.DccAddress = readUint16(data)
	result.Extended = (data[2] & 0x80) != 0
	result.OutputUnitDoesTiming = (data[2] & 0x40) != 0
	result.Activate = (data[2] & 0x20) != 0
	result.Aspect = data[2] & 0b00011111
	result.TimeUnitSec = (data[3] & 0x80) != 0
	result.TimeValue = data[3] & 0b01111111
	return result, nil
}

// Programming commands for the main track (Program On Main) will be issued with this command.
// Followed by other parameters that describe address, selected CV, data and operation to be performed.
// Programming commands will be acknowledged by one or multiple MSG_CS_POM_ACK messages.
// The various acknowledgement levels can be activated in the output unit via FEATURE_GEN_DRIVE_ACK.
// If a track output unit can not issue a command (e.g. because the operation is not implemented),
// it sends a MSG_CS_POM_ACK with ACK=0 nonetheless.
// If a Railcom-capable decoder answered the POM command, the appropriate bidi detector generates a MSG_BM_CV or MSG_BM_XPOM message.
// POM commands are available in several variations, with the following key differences:
// Addressing of the target decoder can be done either via the DCC address (normal procedure)
// or via the decoder identifier. The decoder identifier is a 40-bit number consisting an manufacturer ID and a unique ID.
// In the command the MID (manufacturer ID) field is used to decide whether addressing is done via loco address or DID.
// The target decoder can be a loco or an accessory decoder. At accessory decoders there is a further
// distinction between accessory and extended accessory. This is encoded in the two MSBs of the address, similar to addresses from bidi detectors.
// Addressing of the target CV within the decoder: There is a classic version (POM) where 1024 CV's will
// be addressed and the content (read/write) is one or four bytes. With the introduction of Railcom,
// another version came up which extends the CV address to 24 bits and transmits 32 bits (XPOM).
// These are distinguished by the OPCODE of the BiDiB command.
// The parameters of the command MSG_CS_POM will always be encoded with the maximum field size,
// even if only short CV's will be addressed. Therefore, the address field is always 8+32 bit,
// the CV address field is 24 bit and the data field is 32 bit wide. As customary in BiDiB,
// the LSB will be transmitted first (little-endian).
type CsPom struct {
	BaseMessage
	DccAddress uint32
	Mid        uint8 // 0: Addressing via loco address, 1…255: Addressing via decoder ID, then this field is the manufacturer ID (=DID4)
	OpCode     bidib.CsPomOpCode
	Cv         uint32
	Data       [4]byte
}

func (m CsPom) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := [13]byte{}
	writeUint32(data[0:], m.DccAddress)
	data[4] = m.Mid
	data[5] = byte(m.OpCode)
	data[6] = byte(m.Cv & 0xff)
	data[7] = byte((m.Cv >> 8) & 0xff)
	data[8] = byte((m.Cv >> 16) & 0xff)
	copy(data[9:], m.Data[:])
	bidib.EncodeMessage(write, bidib.MSG_CS_POM, m.Address, seqNum, data[:])
}

func (m CsPom) String() string {
	return fmt.Sprintf("%T addr=%s dccAddr=%d mid=%d opcode=0x%02x cv=%d data=%v", m, m.Address, m.DccAddress, m.Mid, m.OpCode, m.Cv, m.Data)
}

func decodeCsPom(addr bidib.Address, data []byte) (CsPom, error) {
	var result CsPom
	if err := validateDataLength(data, 13); err != nil {
		return result, err
	}
	result.Address = addr
	result.DccAddress = readUint32(data)
	result.Mid = data[4]
	result.OpCode = bidib.CsPomOpCode(data[5])
	result.Cv = uint32(data[6])
	result.Cv |= (uint32(data[7]) << 8)
	result.Cv |= (uint32(data[8]) << 16)
	copy(result.Data[:], data[9:])
	return result, nil
}

// With this command, individual actions can be triggered at the decoder, e.g. activate coupling or analogue functions.
// Followed by 5 bytes: ADDRL, ADDRH, STATEL, STATEH, DATA. STATE denotes the type of DCC message.
type CsBinState struct {
	BaseMessage
	DccAddress uint16
	State      uint16
	Data       uint8
}

func (m CsBinState) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := [5]byte{}
	writeUint16(data[0:], m.DccAddress)
	writeUint16(data[2:], m.State)
	data[4] = m.Data
	bidib.EncodeMessage(write, bidib.MSG_CS_BIN_STATE, m.Address, seqNum, data[:])
}

func (m CsBinState) String() string {
	return fmt.Sprintf("%T addr=%s dccAddr=%d state=%d data=%d", m, m.Address, m.DccAddress, m.State, m.Data)
}

func decodeCsBinState(addr bidib.Address, data []byte) (CsBinState, error) {
	var result CsBinState
	if err := validateDataLength(data, 5); err != nil {
		return result, err
	}
	result.Address = addr
	result.DccAddress = readUint16(data)
	result.State = readUint16(data[2:])
	result.Data = data[4]
	return result, nil
}

// A query of active vehicles is performed with this command. Followed by 1 or more bytes encoding the query: QUERY[, ADDRL, ADDRH]
type CsQuery struct {
	BaseMessage
	// false: Single query of object defined by ADDRL and ADDRH
	//   The node responds with a MSG_CS_DRIVE_STATE.
	// true: Query all objects of the specified type
	//   The node autonomously sends one MSG_CS_DRIVE_STATE for each address known in repeat memory,
	//   adjusting to the available transport capacity itself.
	//   The node must be able to receive and respond to other messages while delivering the answer sequence.
	//   If no object is known, the node responds with a MSG_CS_DRIVE_STATE on address 0.
	QueryAll   bool
	DccAddress uint16
}

func (m CsQuery) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	if m.QueryAll {
		data := []byte{0b10000001}
		bidib.EncodeMessage(write, bidib.MSG_CS_QUERY, m.Address, seqNum, data[:])
	} else {
		data := []byte{0b00000001, 0, 0}
		writeUint16(data[1:], m.DccAddress)
		bidib.EncodeMessage(write, bidib.MSG_CS_QUERY, m.Address, seqNum, data[:])
	}
}

func (m CsQuery) String() string {
	return fmt.Sprintf("%T addr=%s dccAddr=%d all=%t", m, m.Address, m.DccAddress, m.QueryAll)
}

func decodeCsQuery(addr bidib.Address, data []byte) (CsQuery, error) {
	var result CsQuery
	result.Address = addr
	if len(data) == 1 {
		result.QueryAll = true
	} else {
		if err := validateMinDataLength(data, 3); err != nil {
			return result, err
		}
		result.DccAddress = readUint16(data[1:])
	}
	return result, nil
}

// Service mode commands (at programming track) will be issued with this command.
// These track commands are only supported from output units with enabled class bit 3.
// Followed by further parameters, which describes the addressed CV, data and operations to be performed.
// MSG_CS_PROG command will be acknowledged by one or more MSG_CS_PROG_STATE message(s).
// Programming commands are available in several variations (for historical reasons) and they have the following key differences:
// Address programming: simple, direct programming from the target address, no CV select possible.
// Will not be supported from BiDiB.
// Register programming (paged): Selection from a CV within the range 1…8, further CVs can be addressed
// by setting of an page address. This is not any more supported by BiDiB.
// CV-Programming, byte-mode: CV's can be queried byte-wise within an range of 1-1024.
// CV-Programming, bit-mode: A single bit can be queried or set within a CV.
// The target decoder can be an loco decoder (standard) or an accessory decoder.
// There is no difference between an loco or accessory decoder if the programming is made at the programming track.
// Addressing the target CV within the decoder: 1024 CVs can be addressed with the programming track commands.
// Any higher CVs can be addressed by setting of an index register (CV 31 and CV 32) at host side.
type CsProg struct {
	BaseMessage
	OpCode bidib.CsProgOpCode
	Cv     uint16
	Data   uint8
}

func (m CsProg) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{byte(m.OpCode), 0, 0, m.Data}
	writeUint16(data[1:], m.Cv)
	bidib.EncodeMessage(write, bidib.MSG_CS_PROG, m.Address, seqNum, data[:])
}

func (m CsProg) String() string {
	return fmt.Sprintf("%T addr=%s opcode=%d cv=%d data=%d", m, m.Address, m.OpCode, m.Cv, m.Data)
}

func decodeCsProg(addr bidib.Address, data []byte) (CsProg, error) {
	var result CsProg
	if err := validateMinDataLength(data, 3); err != nil {
		return result, err
	}
	result.Address = addr
	result.OpCode = bidib.CsProgOpCode(data[0])
	result.Cv = readUint16(data[1:])
	if len(data) > 3 {
		result.Data = data[3]
	}
	return result, nil
}
