package host

import (
	"context"
	"fmt"
	"time"

	"github.com/binkynet/bidib"
	"github.com/binkynet/bidib/messages"
)

// Any message that is put on the host message queue
type HostMessage interface {
}

// uplinkMessage is send into the host message queue when a messages was received from the interface.
type uplinkMessage struct {
	Addr    bidib.Address
	Message bidib.Message
	Num     bidib.SequenceNumber
}

const (
	uplinkMessageTimeout      = time.Millisecond * 100
	defaultPostOnQueueTimeout = time.Millisecond * 50
)

// callbackMessage is a function that is placed on the host message queue to run
// code in the context of the message queue.
type callbackMessage struct {
	Callback func()
}

// Parse the given message and put into the message queue.
func (h *host) parseAndQueue(mType bidib.MessageType, addr bidib.Address, seqNum bidib.SequenceNumber, data []byte) {
	log := h.log.With().
		Str("type", mType.String()).
		Str("num", seqNum.String()).
		Logger()
	pm, err := messages.Parse(mType, addr, seqNum, data)
	if err != nil {
		log.Warn().Err(err).Msg("failed to parse message")
		return
	}
	if err := h.enqueueMessage(uplinkMessage{
		Addr:    addr,
		Message: pm,
		Num:     seqNum,
	}, uplinkMessageTimeout); err != nil {
		log.Warn().Err(err).Msg("failed to enqueue message")
		return
	}
}

// postOnQueue posts the given function to be called in the context of the message queue.
func (h *host) postOnQueue(cb func(), timeout ...time.Duration) error {
	t := defaultPostOnQueueTimeout
	if len(timeout) > 0 {
		t = timeout[0]
	}
	return h.enqueueMessage(callbackMessage{Callback: cb}, t)
}

// postDelayedOnQueue waits (async) for a given delay, before posting the given message
// on the message queue.
func (h *host) postDelayedOnQueue(cb func(), delay time.Duration, timeout ...time.Duration) {
	go func() {
		time.Sleep(delay)
		h.postOnQueue(cb, timeout...)
	}()
}

// Post the given message onto the message queue
// This is a low level function. Prefer using postOnQueue.
func (h *host) enqueueMessage(msg HostMessage, timeout time.Duration) error {
	select {
	case h.messageQueue <- msg:
		return nil
	case <-time.After(timeout):
		return fmt.Errorf("timeout enqueing %#v", msg)
	}
}

// Process messages from the message queue
func (h *host) runMessageQueue(ctx context.Context) {
	defer close(h.messageQueue)
	for {
		select {
		case <-ctx.Done():
			// Context canceled
			return
		case msg := <-h.messageQueue:
			switch msg := msg.(type) {
			case uplinkMessage:
				h.processUplinkMessage(msg)
			case callbackMessage:
				msg.Callback()
			}
		}
	}
}

// Process a message send to the host by a node in the bidib network.
// This function is to be called by the message loop.
func (h *host) processUplinkMessage(msg uplinkMessage) {
	log := h.log.With().
		Str("num", msg.Num.String()).
		Logger()
	// Find node that for the address
	addr := msg.Addr
	node, found := h.GetNode(addr)
	if !found {
		log.Warn().
			Str("addr", addr.String()).
			Msg("received message for unknown node")
		return
	}

	// Let node process message
	pm := msg.Message
	if err := node.processMessage(pm); err != nil {
		log.Warn().
			Str("addr", addr.String()).
			Interface("msg", pm).
			Msg("failed to process message for node")
	}
	if _, ok := pm.(messages.SysError); ok {
		log.Warn().
			Str("msg", pm.String()).
			Msg("Received error from node")
	} else {
		log.Trace().
			Str("msg", pm.String()).
			Msg("processed message for node")
	}

	// Post process specific messages
	switch pm.(type) {
	case messages.NodeTabCount:
		// If we get a new node table count, disable the interface.
		h.disableSpontaneousMessages()
	case messages.NodeTab:
		// If we have the complete (recursive) node tables,
		// we will enable the interface.
		if h.intfNode.hasCompleteNodeTableRecursive() {
			h.log.Info().Msg("Enabling Bidib")
			h.enableSpontaneousMessages()
			h.invokeNodeChanged(h.intfNode)
		}
	}
}
