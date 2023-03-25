package host

import (
	"fmt"

	"github.com/binkynet/bidib"
	"github.com/binkynet/bidib/messages"
)

// Process the given message
func (h *host) processMessage(mType bidib.MessageType, addr bidib.Address, seqNum bidib.SequenceNumber, data []byte) {
	log := h.log.With().
		Str("type", mType.String()).
		Str("addr", addr.String()).
		Str("num", seqNum.String()).
		Logger()
	pm, err := messages.Parse(mType, addr, seqNum, data)
	if err != nil {
		log.Warn().Err(err).Msg("failed to parse message")
		return
	}

	// Find node that for the address
	node, found := h.GetNode(addr)
	if !found {
		log.Warn().
			Str("addr", addr.String()).
			Msg("received message for unknown node")
		return
	}

	// Let node process message
	if err := node.processMessage(pm); err != nil {
		log.Warn().
			Str("addr", addr.String()).
			Interface("msg", pm).
			Msg("failed to process message for node")
	}
	log.Trace().
		Str("msg", fmt.Sprintf("%s", pm)).
		Msg("processed message for node")

	// Post process specific messages
	switch pm.(type) {
	case messages.NodeTabCount:
		// If we get a new node table count, disable the interface.
		h.intfNode.sendMessages(messages.SysDisable{})
	case messages.NodeTab:
		// If we have the complete (recursive) node tables,
		// we will enable the interface.
		if h.intfNode.hasCompleteNodeTableRecursive() {
			h.log.Info().Msg("Enabling Bidib")
			h.intfNode.sendMessages(messages.SysEnable{})
			h.invokeNodeChanged(h.intfNode)
		}
	}
}
