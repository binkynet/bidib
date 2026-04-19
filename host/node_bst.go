package host

import (
	"github.com/binkynet/bidib"
	"github.com/binkynet/bidib/messages"
)

// NodeBst provides booster extension on the node.
type NodeBst struct {
	*Node
	actualBstState bidib.BstState
}

// GetState returns the last reported BST state of the node.
func (ncs *NodeBst) GetState() bidib.BstState {
	return ncs.actualBstState
}

// Set the Booster in On state.
func (ncs *NodeBst) On() {
	ncs.host.postOnQueue(func() {
		baseMsg := ncs.createBaseMessage()
		ncs.sendMessages(messages.BoostOn{
			BaseMessage:     baseMsg,
			CurrentNodeOnly: false,
		})
	})
}

// Set the Booster in Off state.
func (ncs *NodeBst) Off() {
	ncs.host.postOnQueue(func() {
		baseMsg := ncs.createBaseMessage()
		ncs.sendMessages(messages.BoostOff{
			BaseMessage:     baseMsg,
			CurrentNodeOnly: false,
		})
	})
}

// process the message that is targeted for this node.
func (ncs *NodeBst) processMessage(m bidib.Message) error {
	switch m := m.(type) {
	case messages.BstState:
		if ncs.actualBstState != m.State {
			ncs.actualBstState = m.State
			ncs.invokeNodeChanged(nil)
		}
		ncs.host.invokeBstStateChanged(m)
	}
	return nil
}
