package host

import (
	"time"

	"github.com/binkynet/bidib"
	"github.com/binkynet/bidib/messages"
)

// NodeCs provides commandstation extension on the node.
type NodeCs struct {
	*Node
	actualCsState, desiredCsState bidib.CsState
}

// GetState returns the last reported CS state of the node.
func (ncs *NodeCs) GetState() bidib.CsState {
	return ncs.actualCsState
}

// Set the CS in Off state.
func (ncs *NodeCs) Off() {
	ncs.setState(bidib.BIDIB_CS_STATE_OFF)
}

// Set the CS in Go state.
func (ncs *NodeCs) Go() {
	ncs.setState(bidib.BIDIB_CS_STATE_GO)
}

// Repeat the GO state until the desired state is not longer GO.
func (ncs *NodeCs) repeatGo() {
	wdTimeout, _ := ncs.GetFeature(bidib.FEATURE_GEN_WATCHDOG)
	if wdTimeout == 0 {
		// No watchdog timeout
		return
	}
	delay := time.Millisecond * 100 * time.Duration(wdTimeout/2)
	ncs.host.postDelayedOnQueue(func() {
		if ncs.desiredCsState == bidib.BIDIB_CS_STATE_GO {
			ncs.setState(ncs.desiredCsState)
		}
	}, delay)
}

// Set the CS in Stop state.
func (ncs *NodeCs) Stop() {
	ncs.setState(bidib.BIDIB_CS_STATE_STOP)
}

// setState request a CS state change
func (ncs *NodeCs) setState(state bidib.CsState) {
	ncs.host.postOnQueue(func() {
		ncs.desiredCsState = state
		baseMsg := ncs.createBaseMessage()
		ncs.sendMessages(messages.CsSetState{
			BaseMessage: baseMsg,
			State:       state,
		})
		if state == bidib.BIDIB_CS_STATE_GO {
			ncs.repeatGo()
		}
	})
}

// process the message that is targeted for this node.
func (ncs *NodeCs) processMessage(m bidib.Message) error {
	switch m := m.(type) {
	case messages.CsState:
		if ncs.actualCsState != m.State {
			ncs.actualCsState = m.State
			ncs.invokeNodeChanged()
		}
	}
	return nil
}
