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

type DriveOptions struct {
	DccAddress       uint16
	DccFormat        bidib.DccFormat
	OutputSpeed      bool
	OutputF1_F4      bool
	OutputF5_F8      bool
	OutputF9_F12     bool
	OutputF13_F20    bool
	OutputF21_F28    bool
	DirectionForward bool
	Speed            uint8
	Flags            bidib.DccFlags
}

// Drive instructs the DCC generator to output given drive options
func (ncs *NodeCs) Drive(opts DriveOptions) {
	baseMsg := ncs.createBaseMessage()
	ncs.sendMessages(messages.CsDrive{
		BaseMessage:      baseMsg,
		DccAddress:       opts.DccAddress,
		DccFormat:        opts.DccFormat,
		OutputSpeed:      opts.OutputSpeed,
		OutputF1_F4:      opts.OutputF1_F4,
		OutputF5_F8:      opts.OutputF5_F8,
		OutputF9_F12:     opts.OutputF9_F12,
		OutputF13_F20:    opts.OutputF13_F20,
		OutputF21_F28:    opts.OutputF21_F28,
		DirectionForward: opts.DirectionForward,
		Speed:            opts.Speed,
		Flags:            opts.Flags.Clone(),
	})
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
