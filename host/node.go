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

	// connection used to communicate with the node
	conn transport.Connection
	// logger
	log zerolog.Logger
	// Last used sequence number
	lastSeqNum bidib.SequenceNumber
}

// newNode constructs a new node.
func newNode(addr bidib.Address, conn transport.Connection, log zerolog.Logger) *Node {
	return &Node{
		conn:    conn,
		Address: addr,
		log:     log.With().Str("addr", addr.String()).Logger(),
	}
}

// process the message that is targeted for this node.
func (n *Node) processMessage(m bidib.Message) error {
	switch m := m.(type) {
	case messages.SysMagic:
		n.Magic = m.Magic
	case messages.SysUniqueID:
		n.UniqueID = m.UniqueID
		n.FingerPrint = m.FingerPrint
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
	return n.sendMessages(messages.SysGetMagic{},
		messages.SysGetSwVersion{},
		messages.SysGetUniqueID{},
	)
}
