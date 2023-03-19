package messages

import (
	"fmt"

	"github.com/binkynet/bidib"
)

// This message is sent prior to the transmission of individual MSG_NODETAB, if the host has requested
// with MSG_NODETAB_GETALL.
// Followed by 1 byte with the node table length.
// This table is fetched with a corresponding number of MSG_NODETAB_GETNEXT queries.
type NodeTabCount struct {
	BaseMessage
	TableLength uint8
}

func (m NodeTabCount) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{m.TableLength}
	bidib.EncodeMessage(write, bidib.MSG_NODETAB_COUNT, m.Address, seqNum, data[:])
}

func (m NodeTabCount) String() string {
	return fmt.Sprintf("%T addr=%s tableLength=%d", m, m.Address, m.TableLength)
}

func decodeNodeTabCount(addr bidib.Address, data []byte) (NodeTabCount, error) {
	var result NodeTabCount
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.TableLength = data[0]
	return result, nil
}

// Followed by 9 bytes with an entry of the node mapping table:
// If a node has no subnodes (no class bit 'Hub' is set in the Unique-ID), the node table has
// only one entry length and contains the node itself.
// The transmission of the node table is done by one or more MSG_NODETAB messages.
// While transfer is in progress, no nodes should be added or removed from the table.
// If a change happens nonetheless, the interface must start over again with the transmission of a new MSG_NODETAB_COUNT.
type NodeTab struct {
	BaseMessage
	TableVersion uint8
	NodeAddress  uint8
	UniqueID     bidib.UniqueID
}

func (m NodeTab) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := [9]byte{m.TableVersion, m.NodeAddress}
	copy(data[2:], m.UniqueID[:])
	bidib.EncodeMessage(write, bidib.MSG_NODETAB, m.Address, seqNum, data[:])
}

func (m NodeTab) String() string {
	return fmt.Sprintf("%T addr=%s tableVersion=%d nodeAddr=%d uid=%s", m, m.Address, m.TableVersion, m.NodeAddress, m.UniqueID)
}

func decodeNodeTab(addr bidib.Address, data []byte) (NodeTab, error) {
	var result NodeTab
	if err := validateDataLength(data, 9); err != nil {
		return result, err
	}
	result.Address = addr
	result.TableVersion = data[0]
	result.NodeAddress = data[1]
	copy(result.UniqueID[:], data[2:])
	return result, nil
}

// With this message, a node reports the maximum message length that it can locally handle.
// This is generally restricted by the size of the receive buffer for packets of the respective transport
// layer (which is otherwise transparent towards the host). For packet based transmission of message sequences,
// the length corresponds to the maximum number of bytes for the packet content (as a sequence of only one message), e.g. 64 at BiDiBus.
// Followed by a length designation, consisting of 1 byte, value range 64…127.
// The minimum value is 64, smaller values are reserved and to be ignored. (MSB is reserved for length-extension)
type PktCapacity struct {
	BaseMessage
	Length uint8
}

func (m PktCapacity) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{m.Length}
	bidib.EncodeMessage(write, bidib.MSG_PKT_CAPACITY, m.Address, seqNum, data)
}

func (m PktCapacity) String() string {
	return fmt.Sprintf("%T addr=%s length=%d", m, m.Address, m.Length)
}

func decodePktCapacity(addr bidib.Address, data []byte) (PktCapacity, error) {
	var result PktCapacity
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.Length = data[0]
	return result, nil
}

// Followed by a byte with the (local) number of the addressed node.
// The message is rejected from the interface and will be returned if the host
// attempts to contact a node, which is not (or no longer) in the list.
// This message will be (with node 255) also sent, if all nodes has been already transferred by MSG_NODETAB_GETNEXT.
type NodeNa struct {
	BaseMessage
	NodeAddress uint8
}

func (m NodeNa) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{m.NodeAddress}
	bidib.EncodeMessage(write, bidib.MSG_NODE_NA, m.Address, seqNum, data)
}

func (m NodeNa) String() string {
	return fmt.Sprintf("%T addr=%s nodeAddr=%d", m, m.Address, m.NodeAddress)
}

func decodeNodeNa(addr bidib.Address, data []byte) (NodeNa, error) {
	var result NodeNa
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.NodeAddress = data[0]
	return result, nil
}

// Followed by the current version number of the node table and the table entry of the lost node (see MSG_NODETAB),
// consisting the local address (1 byte) and the Unique-ID (7-bytes).
// An already registered node does not respond any more. If (for example) the lost node is an detector,
// the host can (and should) take appropriate action (partial or general emergency-stop, traffic control).
// The MSG_NODE_LOST must be confirmed by the host. If this message will not be confirmed within 250 ms,
// the interface repeat it at a maximum of 16 times.
type NodeLost struct {
	BaseMessage
	NodeAddress uint8
	UniqueID    bidib.UniqueID
}

func (m NodeLost) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := [8]byte{m.NodeAddress}
	copy(data[1:], m.UniqueID[:])
	bidib.EncodeMessage(write, bidib.MSG_NODE_LOST, m.Address, seqNum, data[:])
}

func (m NodeLost) String() string {
	return fmt.Sprintf("%T addr=%s nodeAddr=%d uid=%s", m, m.Address, m.NodeAddress, m.UniqueID)
}

func decodeNodeLost(addr bidib.Address, data []byte) (NodeLost, error) {
	var result NodeLost
	if err := validateDataLength(data, 8); err != nil {
		return result, err
	}
	result.Address = addr
	result.NodeAddress = data[0]
	copy(result.UniqueID[:], data[1:])
	return result, nil
}

// A new, not yet existing node is detected and added to the node list.
// Followed by the current version number of the node table, and the table entry of this
// new node (see MSG_NODETAB) consisting of local address (1 byte) and Unique-ID (7-bytes).
// The messages for MSG_NODE_LOST and MSG_NODE_NEW will be sent only after the first reading of the node
// table and only if the (spontaneous)-enable at the interface is activated. MSG_NODE_NEW must be confirmed to the host,
// as well as MSG_NODE_LOST. This is done with MSG_NODE_CHANGED_ACK or a complete query beginning with MSG_NODETAB_GETALL,
// otherwise up to 16 repeats take place.
// If multiple changes occur in succession, the version number is incremented each time and a message is generated,
// but only the last change will be repeated.
type NodeNew struct {
	BaseMessage
	TableVersion uint8
	NodeAddress  uint8
	UniqueID     bidib.UniqueID
}

func (m NodeNew) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := [9]byte{m.TableVersion, m.NodeAddress}
	copy(data[2:], m.UniqueID[:])
	bidib.EncodeMessage(write, bidib.MSG_NODE_NEW, m.Address, seqNum, data[:])
}

func (m NodeNew) String() string {
	return fmt.Sprintf("%T addr=%s tableVersion=%d nodeAddr=%d uid=%s", m, m.Address, m.TableVersion, m.NodeAddress, m.UniqueID)
}

func decodeNodeNew(addr bidib.Address, data []byte) (NodeNew, error) {
	var result NodeNew
	if err := validateDataLength(data, 9); err != nil {
		return result, err
	}
	result.Address = addr
	result.TableVersion = data[0]
	result.NodeAddress = data[1]
	copy(result.UniqueID[:], data[2:])
	return result, nil
}

// Followed by a byte which identifies the status.
// 0:	The node operates normal
// 1:	A node sends this message if he detects that his output data buffer is going to be full and
// therefore the current downstream message can not be handled. Such a situation may occur if the host
// "deluge" the node with requests. STALL can also occur if e.g. a sublevel of an interface
// has a lower bandwidth: the interface node isn't able to forward all messages to its subnodes.
// In this case, the host shall not continue to send messages to subnodes this interface.
// A MSG_STALL=1 will be terminated from the node with a MSG_STALL=0.
// Hint:
// In BiDiB, it is desired and allowed that the host summarizes messages and transfers to the node.
// It should, however, not lead to an overloading of the node. For example, this occurs at a BiDiBus structure
// if the answer to an transmitted data package generates more then 48 bytes in total.
// This means, it is possible without any problems, to send a block of
// MSG_GET_SW_VERSION, MSG_GET_P_VERSION, MSG_FEATURE_GETALL, MSG_FEATURE_GETNEXT, MSG_FEATURE_GETNEXT, MSG_FEATURE_GETNEXT
// to read data from the node.
type Stall struct {
	BaseMessage
	Status uint8
}

func (m Stall) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{m.Status}
	bidib.EncodeMessage(write, bidib.MSG_STALL, m.Address, seqNum, data)
}

func (m Stall) String() string {
	return fmt.Sprintf("%T addr=%s status=%d", m, m.Address, m.Status)
}

func decodeStall(addr bidib.Address, data []byte) (Stall, error) {
	var result Stall
	if err := validateDataLength(data, 1); err != nil {
		return result, err
	}
	result.Address = addr
	result.Status = data[0]
	return result, nil
}

// This message is used only on a local level to manage the address assignment in the transport layer (e.g. BiDiBus).
// Followed by 7 bytes with the Unique-ID. The node is trying to logon.
// This message is used at system start up in the process of assigning the local bus addresses.
type LocalLogon struct {
	BaseMessage
	UniqueID bidib.UniqueID
}

func (m LocalLogon) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	bidib.EncodeMessage(write, bidib.MSG_LOGON, m.Address, seqNum, m.UniqueID[:])
}

func (m LocalLogon) String() string {
	return fmt.Sprintf("%T addr=%s uid=%s", m, m.Address, m.UniqueID)
}

func decodeLocalLogon(addr bidib.Address, data []byte) (LocalLogon, error) {
	var result LocalLogon
	if err := validateDataLength(data, 7); err != nil {
		return result, err
	}
	result.Address = addr
	copy(result.UniqueID[:], data)
	return result, nil
}
