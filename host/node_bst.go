package host

import (
	"github.com/binkynet/bidib"
	"github.com/binkynet/bidib/messages"
)

// NodeBst provides booster extension on the node.
type NodeBst struct {
	*Node
	actualBstState bidib.BstState
	actualBstDiag  struct {
		Current     string
		Voltage     string
		Temperature string
	}
}

// GetState returns the last reported BST state of the node.
func (ncs *NodeBst) GetState() bidib.BstState {
	return ncs.actualBstState
}

// Gets last reported current
func (ncs *NodeBst) GetCurrent() string {
	return ncs.actualBstDiag.Current
}

// Gets last reported voltage
func (ncs *NodeBst) GetVoltage() string {
	return ncs.actualBstDiag.Voltage
}

// Gets last reported temperature
func (ncs *NodeBst) GetTemperature() string {
	return ncs.actualBstDiag.Temperature
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
	case messages.BstDiag:
		changed := false
		if compareAndAssign(&ncs.actualBstDiag.Current, m.Current()) {
			changed = true
		}
		if compareAndAssign(&ncs.actualBstDiag.Voltage, m.Voltage()) {
			changed = true
		}
		if compareAndAssign(&ncs.actualBstDiag.Temperature, m.Temperature()) {
			changed = true
		}
		if changed {
			ncs.invokeNodeChanged(nil)
		}
	}
	return nil
}
