package host

import (
	"sync"
	"time"

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
	nextSeqNum bidib.SequenceNumber
	// Table of child nodes
	table struct {
		// Version of the table
		version uint8
		// Number of entries in the table
		count uint8
		// Child nodes
		children []*Node
		// Set when all child nodes are received
		ready bool
	}
	features struct {
		mutex sync.RWMutex
		all   map[bidib.FeatureID]uint8
	}
	extensions struct {
		cs *NodeCs
	}
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

// Gets the feature value with given id.
// Returns value, found
func (n *Node) GetFeature(feature bidib.FeatureID) (uint8, bool) {
	n.features.mutex.RLock()
	defer n.features.mutex.RUnlock()
	result, ok := n.features.all[feature]
	return result, ok
}

// Gets the commandstation extension.
// If this node does not have a DCC signal generator, the result is nil.
func (n *Node) Cs() *NodeCs {
	return n.extensions.cs
}

// Return a base message to include in all specific messages send to this node.
func (n *Node) createBaseMessage() messages.BaseMessage {
	return messages.BaseMessage{Address: n.Address}
}

// process the message that is targeted for this node.
func (n *Node) processMessage(m bidib.Message) error {
	baseMsg := n.createBaseMessage()
	switch m := m.(type) {
	case messages.SysMagic:
		n.Magic = m.Magic
		n.host.invokeNodeChanged(n)
	case messages.SysUniqueID:
		n.UniqueID = m.UniqueID
		n.FingerPrint = m.FingerPrint
		// Set extensions for this node
		n.setupExtensions()
		// If node class indicates subnodes, trigger table discovery
		if n.UniqueID.ClassID().HasSubNodes() {
			n.sendMessages(messages.NodeTabGetAll{BaseMessage: baseMsg})
		}
		n.host.invokeNodeChanged(n)
	case messages.NodeTabCount:
		// Reset node table
		n.table.count = m.TableLength
		n.table.children = nil
		n.table.ready = false
		if m.TableLength == 0 {
			// Table does not yet exist, try again in a bit
			go func() {
				time.Sleep(time.Millisecond * 20)
				n.sendMessages(messages.NodeTabGetAll{BaseMessage: baseMsg})
			}()
		} else if !n.table.ready {
			// Fetch next node table entry
			n.sendMessages(messages.NodeTabGetNext{BaseMessage: baseMsg})
		}
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
			if len(n.table.children) == int(n.table.count) {
				n.table.ready = true
			}
			// Fetch basic info for child node
			child.readNodeProperties()
		}
		// Fetch next node table entry (if any)
		if !n.hasCompleteNodeTable() {
			n.sendMessages(messages.NodeTabGetNext{BaseMessage: baseMsg})
		}
		n.host.invokeNodeChanged(n)
	case messages.NodeNew:
		// Reset node table
		n.table.count = 0
		n.table.children = nil
		n.table.ready = false
		// Refetch node table
		n.sendMessages(messages.NodeTabGetAll{BaseMessage: baseMsg})
		n.host.invokeNodeChanged(n)
	case messages.FeatureCount:
		n.features.mutex.Lock()
		n.features.all = nil
		n.features.mutex.Unlock()
		n.sendMessages(messages.FeatureGetNext{BaseMessage: baseMsg})
	case messages.Feature:
		n.features.mutex.Lock()
		if n.features.all == nil {
			n.features.all = make(map[bidib.FeatureID]uint8)
		}
		n.features.all[m.Feature] = m.Value
		n.features.mutex.Unlock()
		n.sendMessages(messages.FeatureGetNext{BaseMessage: baseMsg})
	}
	return nil
}

// sendMessages sends the given messages to the node, updating the sequence number.
func (n *Node) sendMessages(m ...bidib.Message) error {
	seqNum := n.nextSeqNum
	for i := 0; i < len(m); i++ {
		n.nextSeqNum = n.nextSeqNum.Next()
	}
	return n.conn.SendMessages(m, seqNum)
}

// readNodeProperties sends the commands needed to collect basic node information.
func (n *Node) readNodeProperties() error {
	baseMsg := messages.BaseMessage{Address: n.Address}
	return n.sendMessages(messages.SysGetMagic{BaseMessage: baseMsg},
		messages.SysGetSwVersion{BaseMessage: baseMsg},
		messages.SysGetUniqueID{BaseMessage: baseMsg},
		messages.FeatureGetAll{BaseMessage: baseMsg},
	)
}

// hasCompleteNodeTable returns true if the node has a complete list of child nodes.
func (n *Node) hasCompleteNodeTable() bool {
	if !n.UniqueID.ClassID().HasSubNodes() {
		// No subnodes, we're done
		return true
	}
	return n.table.ready
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

// Set all extensions depending on class ID.
func (n *Node) setupExtensions() {
	if n.UniqueID.ClassID().HasDCCSignalGenerator() {
		n.extensions.cs = &NodeCs{node: n}
	} else {
		n.extensions.cs = nil
	}
}
