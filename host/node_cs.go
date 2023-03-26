package host

import (
	"github.com/binkynet/bidib"
	"github.com/binkynet/bidib/messages"
)

// NodeCs provides commandstation extension on the node.
type NodeCs struct {
	node    *Node
	csState bidib.CsState
}

// GetState returns the last reported CS state of the node.
func (ncs *NodeCs) GetState() bidib.CsState {
	return ncs.csState
}

// Set the CS in Off state.
func (ncs *NodeCs) Off() {
	baseMsg := ncs.node.createBaseMessage()
	ncs.node.sendMessages(messages.CsSetState{
		BaseMessage: baseMsg,
		State:       bidib.BIDIB_CS_STATE_OFF,
	})
}

// Set the CS in Go state.
func (ncs *NodeCs) Go() {
	baseMsg := ncs.node.createBaseMessage()
	ncs.node.sendMessages(messages.CsSetState{
		BaseMessage: baseMsg,
		State:       bidib.BIDIB_CS_STATE_GO,
	})
}

// Set the CS in Stop state.
func (ncs *NodeCs) Stop() {
	baseMsg := ncs.node.createBaseMessage()
	ncs.node.sendMessages(messages.CsSetState{
		BaseMessage: baseMsg,
		State:       bidib.BIDIB_CS_STATE_STOP,
	})
}

// process the message that is targeted for this node.
func (ncs *NodeCs) processMessage(m bidib.Message) error {
	switch m := m.(type) {
	case messages.CsState:
		ncs.csState = m.State
		ncs.node.invokeNodeChanged()
	}
	return nil
}
