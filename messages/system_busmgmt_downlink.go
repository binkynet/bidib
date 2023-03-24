package messages

import (
	"fmt"

	"github.com/binkynet/bidib"
)

// The BiDiB system will be reset with regard to the host interface and the allocation of all nodes is carried out again.
// The previous assignment table is void. Interfaces inherit this message automatically to all following sub nodes.
// All message sequence numbers in the upstream will be set back to zero but the function of the node remains.
// If this message is addressed to a node, it shall log off (ie. shutdown for 1s, the interface will drop the node) an
// try to reconnect again. Internal states of the node may be lost.
// Attention: All undelivered messages in the uplink and downlink are lost by a MSG_SYS_RESET and the BiDiB system goes into
// the status SYS_DISABLE. After a reset, the data consistency must be checked and all information must be read again from the BiDiB system.
// Attention: A bus structure will be reset through a MSG_SYS_RESET, this means a 1 second break on the bus in case of
// BiDiBus (RS485) in order to allow all node to switch into disconnected state. Only after this time, the bus is back
// up and the node table is valid again. During this time, MSG_NODETAB_GETALL is answered with 0.
// Attention: The access to nodes and forwarding of messages is only possible if the node table becomes valid again,
// especially for broadcast messages, such as MSG_SYS_ENABLE or MSG_SYS_CLOCK.
type SysReset struct {
	BaseMessage
}

func (m SysReset) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_SYS_RESET, m.Address, seqNum, nil)
}

func (m SysReset) String() string {
	return fmt.Sprintf("%T addr=%s", m, m.Address)
}

func decodeSysReset(addr bidib.Address, data []byte) (SysReset, error) {
	var result SysReset
	if err := validateDataLength(data, 0); err != nil {
		return result, err
	}
	result.Address = addr
	return result, nil
}

// With this command, the interface is caused to transfer the current assignment table of Unique-ID and local address.
// This transfer is a series of messages, it is started with a MSG_NODETAB_COUNT and will be followed by MSG_NODETAB,
// each of which are triggered by MSG_NODETAB_GETNEXT.
// While the transfer is in progress, new inquiries with MSG_NODETAB_GETALL lead to abortion and restart of the transfer.
// If the table does not yet exist, the interface responds with a MSG_NODETAB_COUNT = 0 message.
// In this case, the host must ask for it again after a few ms.
// If the table is existing, the interface responds with MSG_NODETAB_COUNT = 'table length'.
// Hint:
// This message applies only to nodes which have a registered hub in their Unique-ID. Nodes, which contain no
// substructure must answer this request anyway. In this case the node table has only one entry, local address is 0.
type NodeTabGetAll struct {
	BaseMessage
}

func (m NodeTabGetAll) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_NODETAB_GETALL, m.Address, seqNum, nil)
}

func (m NodeTabGetAll) String() string {
	return fmt.Sprintf("%T addr=%s", m, m.Address)
}

func decodeNodeTabGetAll(addr bidib.Address, data []byte) (NodeTabGetAll, error) {
	var result NodeTabGetAll
	if err := validateDataLength(data, 0); err != nil {
		return result, err
	}
	result.Address = addr
	return result, nil
}

// This command causes the interface to send the next line of the node table. No parameters follow.
// The node responds with a MSG_NODETAB message. In case there is no (more) line on hand,
// it responds with MSG_NODE_NA 255 instead. If there was a change in the node table since the last
// transmission of MSG_NODETAB_COUNT, the node responds with MSG_NODETAB_COUNT and starts over with
// sending MSG_NODETAB messages (with the incremented version number).
type NodeTabGetNext struct {
	BaseMessage
}

func (m NodeTabGetNext) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_NODETAB_GETNEXT, m.Address, seqNum, nil)
}

func (m NodeTabGetNext) String() string {
	return fmt.Sprintf("%T addr=%s", m, m.Address)
}

func decodeNodeTabGetNext(addr bidib.Address, data []byte) (NodeTabGetNext, error) {
	var result NodeTabGetNext
	if err := validateDataLength(data, 0); err != nil {
		return result, err
	}
	result.Address = addr
	return result, nil
}

// With this command, it is possible to read in the maximum message length that a node can handle.
// This corresponds to the maximum length of a message sequence when it consists of only one message,
// and thereby the maximum number of bytes in a packet (between two frame markers).
// The node responds with a MSG_PKT_CAPACITY message. Until a node responds with a value above 64,
// the default restriction to 64 is in effect.
type GetPktCapacity struct {
	BaseMessage
}

func (m GetPktCapacity) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_GET_PKT_CAPACITY, m.Address, seqNum, nil)
}

func (m GetPktCapacity) String() string {
	return fmt.Sprintf("%T addr=%s", m, m.Address)
}

func decodeGetPktCapacity(addr bidib.Address, data []byte) (GetPktCapacity, error) {
	var result GetPktCapacity
	if err := validateDataLength(data, 0); err != nil {
		return result, err
	}
	result.Address = addr
	return result, nil
}

// Followed by one byte with the confirmed sequence number (version number of the node table)
// of NODE_NEW or NODE_LOST message. The host sends this message within 250ms to an interface
// when he received a notification for a lost or newly added node.
// If the interface gets the same version of node table that it has sent in the last change notification,
// this and all previous changes are considered as acknowledged.
type NodeChangedAck struct {
	BaseMessage
	VersionNumber uint8
}

func (m NodeChangedAck) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{m.VersionNumber}
	bidib.EncodeMessage(write, bidib.MSG_NODE_CHANGED_ACK, m.Address, seqNum, data)
}

func (m NodeChangedAck) String() string {
	return fmt.Sprintf("%T addr=%s versionNum=0x%02x", m, m.Address, m.VersionNumber)
}

func decodeNodeChangedAck(addr bidib.Address, data []byte) (NodeChangedAck, error) {
	var result NodeChangedAck
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.VersionNumber = data[0]
	return result, nil
}

// This message is used only on a local level to manage the address assignment in the transport layer (e.g. BiDiBus).
// Followed by one byte with the local address (NODE_ADDR) and 7 bytes with the Unique-ID.
// Only if the node has verified that the received Unique-ID and the internal Unique-ID is identical,
// he may set his local address to received NODE_ADDR.
// This message will be sent as broadcast and with MNUM = 0. It is always interpreted,
// even if no login attempt was made. It is therefore possible to assign a local address to a node before the general logon.
type LocalLogonAck struct {
	BaseMessage
	NodeAddress uint8
	UniqueID    bidib.UniqueID
}

func (m LocalLogonAck) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := [8]byte{}
	data[0] = m.NodeAddress
	copy(data[1:], m.UniqueID[:])
	bidib.EncodeMessage(write, bidib.MSG_LOGON_ACK, m.Address, seqNum, data[:])
}

func (m LocalLogonAck) String() string {
	return fmt.Sprintf("%T addr=%s nodeAddr=%d uid=%s", m, m.Address, m.NodeAddress, m.UniqueID)
}

func decodeLocalLogonAck(addr bidib.Address, data []byte) (LocalLogonAck, error) {
	var result LocalLogonAck
	if err := validateDataLength(data, 8); err != nil {
		return result, err
	}
	result.Address = addr
	result.NodeAddress = data[0]
	copy(result.UniqueID[:], data[1:])
	return result, nil
}

// This message is used only on a local level to manage the address assignment in the transport layer (e.g. BiDiBus).
// Followed by 7 bytes with the Unique-ID. The logon attempts of the addressed node are refused.
// Possible causes for the rejection of the LOGON can be:
// - The interface table is full, the max. number of participants is reached.
// - A double Unique-ID was recognized on the bus.
// Simultaneously with the MSG_LOCAL_LOGON_REJECTED, the interface sends an error message with BIDIB_ERR_BUS to the host.
type LocalLogonRejected struct {
	BaseMessage
	UniqueID bidib.UniqueID
}

func (m LocalLogonRejected) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_LOGON_REJECTED, m.Address, seqNum, m.UniqueID[:])
}

func (m LocalLogonRejected) String() string {
	return fmt.Sprintf("%T addr=%s uid=%s", m, m.Address, m.UniqueID)
}

func decodeLocalLogonRejected(addr bidib.Address, data []byte) (LocalLogonRejected, error) {
	var result LocalLogonRejected
	if err := validateDataLength(data, 7); err != nil {
		return result, err
	}
	result.Address = addr
	copy(result.UniqueID[:], data[:])
	return result, nil
}
