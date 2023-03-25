package host

import (
	"github.com/binkynet/bidib"
	"github.com/binkynet/bidib/messages"
	"github.com/binkynet/bidib/transport"
	"github.com/rs/zerolog"
)

// Node represents a node in the bidib network.
type Node struct {
	// Address of the node
	Address bidib.Address
	// Unique ID of the node
	UniqueID bidib.UniqueID
	// Fingerprint of the node
	FingerPrint uint32
	// Magic of the node
	Magic uint16

	// Host containing this node
	host *host
	// connection used to communicate with the node
	conn transport.Connection
	// logger
	log zerolog.Logger
	// Last used sequence number
	lastSeqNum bidib.SequenceNumber
	// Table of child nodes
	table struct {
		// Version of the table
		version uint8
		// Number of entries in the table
		count uint8
		// Child nodes
		children []*Node
	}
	features map[bidib.FeatureID]uint8
}

// newNode constructs a new node.
func newNode(addr bidib.Address, host *host, conn transport.Connection, log zerolog.Logger) *Node {
	return &Node{
		host:    host,
		conn:    conn,
		Address: addr,
		log:     log.With().Str("addr", addr.String()).Logger(),
	}
}

// ForEachChild calls the given function for each (direct) child of this node.
func (n *Node) ForEachChild(cb func(*Node)) {
	for _, child := range n.table.children {
		if child != nil {
			cb(child)
		}
	}
}

// process the message that is targeted for this node.
func (n *Node) processMessage(m bidib.Message) error {
	baseMsg := messages.BaseMessage{Address: n.Address}
	switch m := m.(type) {
	case messages.SysMagic:
		n.Magic = m.Magic
		n.host.invokeNodeChanged(n)
	case messages.SysUniqueID:
		n.UniqueID = m.UniqueID
		n.FingerPrint = m.FingerPrint
		// If node class indicates subnodes, trigger table discovery
		if n.UniqueID.ClassID().HasSubNodes() {
			n.sendMessages(messages.NodeTabGetAll{BaseMessage: baseMsg})
		}
		n.host.invokeNodeChanged(n)
	case messages.NodeTabCount:
		// Reset node table
		n.table.count = m.TableLength
		n.table.children = nil
		// Fetch next node table entry
		n.sendMessages(messages.NodeTabGetNext{BaseMessage: baseMsg})
		n.host.invokeNodeChanged(n)
	case messages.NodeTab:
		n.table.version = m.TableVersion
		if m.NodeAddress == 0 {
			// Got my own node
			n.table.children = append(n.table.children, nil)
		} else {
			// Found new child node
			childAddr := n.Address.Append(m.NodeAddress)
			child := newNode(childAddr, n.host, n.conn, n.log)
			n.table.children = append(n.table.children, child)
			// Fetch basic info for child node
			child.readNodeProperties()
			child.readNodeFeatures()
		}
		// Fetch next node table entry (if any)
		if !n.hasCompleteNodeTable() {
			n.sendMessages(messages.NodeTabGetNext{BaseMessage: baseMsg})
		}
		n.host.invokeNodeChanged(n)
	case messages.FeatureCount:
		n.features = nil
		n.sendMessages(messages.FeatureGetNext{BaseMessage: baseMsg})
	case messages.Feature:
		if n.features == nil {
			n.features = make(map[bidib.FeatureID]uint8)
		}
		n.features[m.Feature] = m.Value
		n.sendMessages(messages.FeatureGetNext{BaseMessage: baseMsg})
	}
	return nil
}

// sendMessages sends the given messages to the node, updating the sequence number.
func (n *Node) sendMessages(m ...bidib.Message) error {
	seqNum := n.lastSeqNum + 1
	n.lastSeqNum += bidib.SequenceNumber(len(m))
	return n.conn.SendMessages(m, seqNum)
}

// readNodeProperties sends the commands needed to collect basic node information.
func (n *Node) readNodeProperties() error {
	baseMsg := messages.BaseMessage{Address: n.Address}
	return n.sendMessages(messages.SysGetMagic{BaseMessage: baseMsg},
		messages.SysGetSwVersion{BaseMessage: baseMsg},
		messages.SysGetUniqueID{BaseMessage: baseMsg},
	)
}

// readNodeFeatures sends the commands deeded to collect feature information of this node
func (n *Node) readNodeFeatures() error {
	baseMsg := messages.BaseMessage{Address: n.Address}
	return n.sendMessages(messages.FeatureGetAll{BaseMessage: baseMsg})
}

// hasCompleteNodeTable returns true if the node has a complete list of child nodes.
func (n *Node) hasCompleteNodeTable() bool {
	if !n.UniqueID.ClassID().HasSubNodes() {
		// No subnodes, we're done
		return true
	}
	return n.table.count > 0 && len(n.table.children) == int(n.table.count)
}

// hasCompleteNodeTableRecursive returns true if the node has a complete list of child nodes
// AND all child nodes also have a complete list of tables.
func (n *Node) hasCompleteNodeTableRecursive() bool {
	if !n.hasCompleteNodeTable() {
		return false
	}
	for _, child := range n.table.children {
		if child == nil {
			continue
		}
		if !child.hasCompleteNodeTableRecursive() {
			return false
		}
	}
	return true
}
