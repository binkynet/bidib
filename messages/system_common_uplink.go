package messages

import (
	"fmt"

	"github.com/binkynet/bidib"
)

// Transmission of the system identifier: This variable is used for identification and transmission control.
// Followed by 2 data bytes, MAGICL, MAGICH which indicates the system identifier.
// The system identifier is transmitted with a transmission sequence index 0, this
// restarts the synchronisation of message sequence at host side.
type SysMagic struct {
	BaseMessage
	Magic uint16
}

func (m SysMagic) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	var data [2]byte
	writeUint16(data[:], m.Magic)
	bidib.EncodeMessage(write, bidib.MSG_SYS_MAGIC, m.Address, seqNum, data[:])
}

func (m SysMagic) String() string {
	return fmt.Sprintf("%T addr=%s magic=0x%04x", m, m.Address, m.Magic)
}

func decodeSysMagic(addr bidib.Address, data []byte) (SysMagic, error) {
	var result SysMagic
	if err := validateDataLength(data, 2); err != nil {
		return result, err
	}
	result.Address = addr
	result.Magic = readUint16(data)
	return result, nil
}

// Followed by one byte. This message is a response to the MSG_SYS_PING request,
// while the transferred byte in PING will be sent back.
type SysPong struct {
	BaseMessage
	Value uint8
}

func (m SysPong) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{m.Value}
	bidib.EncodeMessage(write, bidib.MSG_SYS_PONG, m.Address, seqNum, data)
}

func (m SysPong) String() string {
	return fmt.Sprintf("%T addr=%s value=0x%02x", m, m.Address, m.Value)
}

func decodeSysPong(addr bidib.Address, data []byte) (SysPong, error) {
	var result SysPong
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.Value = data[0]
	return result, nil
}

// This message is used only on a local level to support the outage detection of the transport layer (e.g. serial link).
// Empty message, no other data will follow. This message is the response to MSG_LOCAL_PING.
type LocalPong struct {
	BaseMessage
}

func (m LocalPong) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_LOCAL_PONG, m.Address, seqNum, nil)
}

func (m LocalPong) String() string {
	return fmt.Sprintf("%T addr=%s", m, m.Address)
}

func decodeLocalPong(addr bidib.Address, data []byte) (LocalPong, error) {
	var result LocalPong
	if err := validateDataLength(data, 0); err != nil {
		return result, err
	}
	result.Address = addr
	return result, nil
}

// Transmission of the supported protocol version.
// Followed by 2 data bytes which encode the BiDiB protocol version.
type SysPVersion struct {
	BaseMessage
	Minor uint8
	Major uint8
}

func (m SysPVersion) String() string {
	return fmt.Sprintf("%T addr=%s verion=%d.%d", m, m.Address, m.Major, m.Minor)
}

func (m SysPVersion) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{m.Minor, m.Major}
	bidib.EncodeMessage(write, bidib.MSG_SYS_P_VERSION, m.Address, seqNum, data)
}

func decodeSysPVersion(addr bidib.Address, data []byte) (SysPVersion, error) {
	var result SysPVersion
	if err := validateDataLength(data, 2); err != nil {
		return result, err
	}
	result.Address = addr
	result.Minor = data[0]
	result.Major = data[1]
	return result, nil
}

// The node sends its unique identifier. Followed by 7 bytes with the Unique-ID and optionally 4 bytes with a configuration fingerprint.
// The fingerprint is a 32 bit checksum applied to all settings of the node. Those include:
// - Features (MSG_FEATURE)
// - User configuration (MSG_VENDOR and MSG_STRING)
// - Accessory configuration (MSG_ACCESSORY_PARA)
// - Port configuration (MSG_LC_CONFIGX)
// - Macro configuration (MSG_LC_MACRO and MSG_LC_MACRO_PARA)
// Explicitly exempt are all bus and operation states (even those that are persisted across power cycles),
// supported protocol versions and firmware revisions (assuming nothing else changes).
// The fingerprint is computed by the node using a good (uniformly distributed, chaotic, efficient) but
// not necessarily cryptographic hash function. When a configuration value changes, the fingerprint changes as well.
// Hints:
// Fingerprinting is optional for nodes. When it is not supported, only the Unique-ID is transmitted.
// The hash function needs to deliver the current value on every request, it is not sufficient to compute it
// only once during startup. For increasing the efficiency it is possible to choose an incremental algorithm
// which individually incorporates every change of a setting into the result.
// The fingerprint is destined for quickly loading node configurations at the start of a session.
// Once read or written, a host can locally store the configuration values for a node together with the
// corresponding fingerprint. When the node is met the next time, the values can be identified by Unique-ID and
// fingerprint in the cache, and may be loaded directly without needing to interrogate the node.
type SysUniqueID struct {
	BaseMessage
	UniqueID    bidib.UniqueID
	FingerPrint uint32
}

func (m SysUniqueID) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	if m.FingerPrint == 0 {
		bidib.EncodeMessage(write, bidib.MSG_SYS_UNIQUE_ID, m.Address, seqNum, m.UniqueID[:])
	} else {
		data := [7 + 4]byte{}
		copy(data[0:], m.UniqueID[:])
		writeUint32(data[7:], m.FingerPrint)
		bidib.EncodeMessage(write, bidib.MSG_SYS_UNIQUE_ID, m.Address, seqNum, data[:])
	}
}

func (m SysUniqueID) String() string {
	return fmt.Sprintf("%T addr=%s uid=%s fingerprint=0x%08x", m, m.Address, m.UniqueID, m.FingerPrint)
}

func decodeSysUniqueID(addr bidib.Address, data []byte) (SysUniqueID, error) {
	var result SysUniqueID
	result.Address = addr
	if len(data) == 7 {
		copy(result.UniqueID[:], data)
	} else if err := validateDataLength(data, 7+4); err != nil {
		return result, err
	} else {
		result.FingerPrint = readUint32(data[7:])
	}
	return result, nil
}

// SubRevision, something, Main Revision
type VersionTriple [3]byte

func (vt VersionTriple) String() string {
	return fmt.Sprintf("%d.%d.%d", vt[2], vt[1], vt[0])
}

// Transmission of the software version: Followed by 1 to 16 triples (3 bytes each), vendor specific.
// Inside each triple the sub revision index is transmitted first, the main revision index is transferred last.
// Newer versions have a numerically larger version index.
// The first triple denotes the software version of the node, additional triples may code the
// version of subsystems (like coprocessors, hardware).
// Note: up to BiDiB specification revision 1.21 only one triple was defined as answer.
type SysSwVersion struct {
	BaseMessage
	Versions []VersionTriple
}

func (m SysSwVersion) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	var data []byte
	if len(m.Versions) > 0 {
		data = make([]byte, 3*len(m.Versions))
		for i, v := range m.Versions {
			copy(data[i*3:], v[:])
		}
	}
	bidib.EncodeMessage(write, bidib.MSG_SYS_SW_VERSION, m.Address, seqNum, data)
}

func (m SysSwVersion) String() string {
	return fmt.Sprintf("%T addr=%s verions=%v", m, m.Address, m.Versions)
}

func decodeSysSwVersion(addr bidib.Address, data []byte) (SysSwVersion, error) {
	var result SysSwVersion
	result.Address = addr
	dataLen := len(data)
	if dataLen%3 != 0 {
		return result, fmt.Errorf("data length must be multiple of 3")
	}
	versions := dataLen / 3
	if versions < 1 || versions > 16 {
		return result, fmt.Errorf("data must contain between 1 and 16 versions, got %d", versions)
	}
	result.Versions = make([]VersionTriple, versions)
	for i := 0; i < versions; i++ {
		copy(result.Versions[i][:], data[i*3:])
	}
	return result, nil
}

// Followed by 1 byte with the identify status: 0: off, 1: on.
// This message is sent when identification of the node was triggered, either by host
// command (MSG_SYS_IDENTIFY) or locally by the identification button.
// Recommendation: If the identify button is assigned to more than one function (e.g. if a decoder can
// also be programmed via DCC address-learning), a short press should execute identify, a long press
// should execute the DCC learning mode.
// This recommendation ensures the same behaviour across different modules.
type SysIdentityState struct {
	BaseMessage
	Value bool
}

func (m SysIdentityState) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{0}
	if m.Value {
		data[0] = 1
	}
	bidib.EncodeMessage(write, bidib.MSG_SYS_IDENTIFY_STATE, m.Address, seqNum, data)
}

func (m SysIdentityState) String() string {
	return fmt.Sprintf("%T addr=%s value=%v", m, m.Address, m.Value)
}

func decodeSysIdentityState(addr bidib.Address, data []byte) (SysIdentityState, error) {
	var result SysIdentityState
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.Value = data[0] != 0
	return result, nil
}

// Error message of a node. The errors take place either by a query (by MSG_SYS_GET_ERROR)
// or spontaneous (if the node is enabled). Followed by one byte with the error type and
// occasionally other parameters.
// Depending on the error, the processing of the data will not be possible any more.
type SysError struct {
	BaseMessage
	Error bidib.ErrorCode
}

func (m SysError) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{byte(m.Error)}
	bidib.EncodeMessage(write, bidib.MSG_SYS_ERROR, m.Address, seqNum, data)
}

func (m SysError) String() string {
	return fmt.Sprintf("%T addr=%s error=0x%02x", m, m.Address, m.Error)
}

func decodeSysError(addr bidib.Address, data []byte) (SysError, error) {
	var result SysError
	if err := validateMinDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.Error = bidib.ErrorCode(data[0])
	return result, nil
}
