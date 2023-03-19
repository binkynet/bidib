package messages

import (
	"fmt"

	"github.com/binkynet/bidib"
)

// The addressed BiDiB Node should transmit the system identifier.
type SysGetMagic struct {
	BaseMessage
}

func (m SysGetMagic) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_SYS_GET_MAGIC, m.Address, seqNum, nil)
}

func (m SysGetMagic) String() string {
	return fmt.Sprintf("%T addr=%s", m, m.Address)
}

func decodeSysGetMagic(addr bidib.Address, data []byte) (SysGetMagic, error) {
	var result SysGetMagic
	if err := validateDataLength(data, 0); err != nil {
		return result, err
	}
	result.Address = addr
	return result, nil
}

// Query for the supported BiDiB protocol version.
type SysGetPVersion struct {
	BaseMessage
}

func (m SysGetPVersion) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_SYS_GET_P_VERSION, m.Address, seqNum, nil)
}

func (m SysGetPVersion) String() string {
	return fmt.Sprintf("%T addr=%s", m, m.Address)
}

func decodeSysGetPVersion(addr bidib.Address, data []byte) (SysGetPVersion, error) {
	var result SysGetPVersion
	if err := validateDataLength(data, 0); err != nil {
		return result, err
	}
	result.Address = addr
	return result, nil
}

// The node will be released, from now on, spontaneous messages are possible (e.g. change of occupancy states, new added hardware).
// The message is automatically passed to all other subnodes (inherited). No acknowledgement will follow.
type SysEnable struct {
	BaseMessage
}

func (m SysEnable) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_SYS_ENABLE, m.Address, seqNum, nil)
}

func (m SysEnable) String() string {
	return fmt.Sprintf("%T addr=%s", m, m.Address)
}

func decodeSysEnable(addr bidib.Address, data []byte) (SysEnable, error) {
	var result SysEnable
	if err := validateDataLength(data, 0); err != nil {
		return result, err
	}
	result.Address = addr
	return result, nil
}

// The BiDiB system will be blocked, spontaneous messages are disabled at this point.
// Events which occur in the SYS_DISABLE state will not be cached, yet node states can be queried targeted.
// The message is automatically passed to all other subnodes (inherited) and should therefore be
// addressed only to node 0. No acknowledgement will follow.
type SysDisable struct {
	BaseMessage
}

func (m SysDisable) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_SYS_DISABLE, m.Address, seqNum, nil)
}

func (m SysDisable) String() string {
	return fmt.Sprintf("%T addr=%s", m, m.Address)
}

func decodeSysDisable(addr bidib.Address, data []byte) (SysDisable, error) {
	var result SysDisable
	if err := validateDataLength(data, 0); err != nil {
		return result, err
	}
	result.Address = addr
	return result, nil
}

// Query for the Unique-ID and configuration fingerprint of a node. No other data will follow.
// The corresponding node responds with MSG_SYS_UNIQUE_ID.
type SysGetUniqueID struct {
	BaseMessage
}

func (m SysGetUniqueID) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_SYS_GET_UNIQUE_ID, m.Address, seqNum, nil)
}

func (m SysGetUniqueID) String() string {
	return fmt.Sprintf("%T addr=%s", m, m.Address)
}

func decodeSysGetUniqueID(addr bidib.Address, data []byte) (SysGetUniqueID, error) {
	var result SysGetUniqueID
	if err := validateDataLength(data, 0); err != nil {
		return result, err
	}
	result.Address = addr
	return result, nil
}

// Query of the node's installed software version(s). No other data will follow.
type SysGetSwVersion struct {
	BaseMessage
}

func (m SysGetSwVersion) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_SYS_GET_SW_VERSION, m.Address, seqNum, nil)
}

func (m SysGetSwVersion) String() string {
	return fmt.Sprintf("%T addr=%s", m, m.Address)
}

func decodeSysGetSwVersion(addr bidib.Address, data []byte) (SysGetSwVersion, error) {
	var result SysGetSwVersion
	if err := validateDataLength(data, 0); err != nil {
		return result, err
	}
	result.Address = addr
	return result, nil
}

// The corresponding node in the BiDiB system is initiated to send an empty message (MSG_SYS_PONG) back.
// This response must be received within 250 ms, otherwise the host has to consider the corresponding node as failed.
// The passed parameter (byte) is returned by MSG_SYS_PONG.
// Followed by one byte.
type SysPing struct {
	BaseMessage
	Value uint8
}

func (m SysPing) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{m.Value}
	bidib.EncodeMessage(write, bidib.MSG_SYS_PING, m.Address, seqNum, data)
}

func (m SysPing) String() string {
	return fmt.Sprintf("%T addr=%s value=0x%02x", m, m.Address, m.Value)
}

func decodeSysPing(addr bidib.Address, data []byte) (SysPing, error) {
	var result SysPing
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.Value = data[0]
	return result, nil
}

// This message is used only on a local level to support the outage detection of the transport layer (e.g. serial link).
// No other data will follow.
// The corresponding node in the BiDiB system is initiated to send a message of type MSG_LOCAL_PONG back.
// This response must be received within 250 ms, otherwise the host has to consider the corresponding node
// as failed and transmit this failure with MSG_NODE_LOST to the host.
// (In case of BiDiBus, the token and the answer to that token takes over this function.)
type LocalPing struct {
	BaseMessage
}

func (m LocalPing) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_LOCAL_PING, m.Address, seqNum, nil)
}

func (m LocalPing) String() string {
	return fmt.Sprintf("%T addr=%s", m, m.Address)
}

func decodeLocalPing(addr bidib.Address, data []byte) (LocalPing, error) {
	var result LocalPing
	if err := validateDataLength(data, 0); err != nil {
		return result, err
	}
	result.Address = addr
	return result, nil
}

// Followed by one byte: Identify is switched off, 1: Identify is switched on.
// The corresponding node in the BiDiB system is instructed to display a local identify
// indicator (e.g. a flashing LED). The node responds with a MSG_SYS_IDENTIFY_STATE message.
type SysIdentify struct {
	BaseMessage
	Value bool
}

func (m SysIdentify) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{0}
	if m.Value {
		data[0] = 1
	}
	bidib.EncodeMessage(write, bidib.MSG_SYS_IDENTIFY, m.Address, seqNum, data)
}

func (m SysIdentify) String() string {
	return fmt.Sprintf("%T addr=%s value=%v", m, m.Address, m.Value)
}

func decodeSysIdentify(addr bidib.Address, data []byte) (SysIdentify, error) {
	var result SysIdentify
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.Value = data[0] != 0
	return result, nil
}

// The last occurred (but not any spontaneous) error message is read.
// The error memory is cleared through reading. If there is no error,
// an empty error message (i.e. error number 0) is returned.
type SysGetError struct {
	BaseMessage
}

func (m SysGetError) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_SYS_GET_ERROR, m.Address, seqNum, nil)
}

func (m SysGetError) String() string {
	return fmt.Sprintf("%T addr=%s", m, m.Address)
}

func decodeSysGetError(addr bidib.Address, data []byte) (SysGetError, error) {
	var result SysGetError
	if err := validateDataLength(data, 0); err != nil {
		return result, err
	}
	result.Address = addr
	return result, nil
}

// This message is used only on a local level to support the system time synchronisation on the transport layer.
// Followed by two bytes (TIMEL, TIMEH) with the BiDiB system time, TIME indicates the point in time of the
// last frame marker prior to the message. The node sets its local clock to the received timestamp,
// possible data transit times must to be compensated by corresponding offsets. The message is not replied to.
// Example: A node receives MSG_LOCAL_SYNC 3. The duration from the frame signal to the processing of the
// message amounts to 2 ms. The internal clock will be set to 5.
type LocalSync struct {
	BaseMessage
	Time uint16
}

func (m LocalSync) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	var data [2]byte
	writeUint16(data[:], m.Time)
	bidib.EncodeMessage(write, bidib.MSG_LOCAL_SYNC, m.Address, seqNum, data[:])
}

func (m LocalSync) String() string {
	return fmt.Sprintf("%T addr=%s time=0x%04d", m, m.Address, m.Time)
}

func decodeLocalSync(addr bidib.Address, data []byte) (LocalSync, error) {
	var result LocalSync
	if err := validateDataLength(data, 2); err != nil {
		return result, err
	}
	result.Address = addr
	result.Time = readUint16(data)
	return result, nil
}
