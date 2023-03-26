package host

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/rs/zerolog"

	"github.com/binkynet/bidib"
	"github.com/binkynet/bidib/messages"
	"github.com/binkynet/bidib/transport"
	"github.com/binkynet/bidib/transport/serial"
)

// Host defines external interface of a Bidib host process.
type Host interface {
	// Returns the root of the node tree
	GetRootNode() *Node
	// Gets the node with the given address.
	// Returns nil, false if not found
	GetNode(addr bidib.Address) (*Node, bool)
	// Register a callback that gets invoked on every node change
	RegisterNodeChanged(func(*Node)) context.CancelFunc
	// Close the connections
	Close() error
}

// Host config
type Config struct {
	Serial *serial.Config
}

const (
	messageQueueBufLen = 64
)

// New constructs a new host process with given config
func New(cfg Config, log zerolog.Logger) (Host, error) {
	h := &host{
		Config:       cfg,
		log:          log,
		messageQueue: make(chan HostMessage, messageQueueBufLen),
	}
	if err := h.start(); err != nil {
		h.Close()
		return nil, err
	}
	return h, nil
}

// host implements the Bidib host process.
type host struct {
	Config
	log              zerolog.Logger
	conn             transport.Connection
	intfNode         *Node
	disabledState    int32
	nodeChangedEvent Event[*Node]
	messageQueue     chan HostMessage
	cancelQueue      context.CancelFunc
}

const (
	disabledStateUnknown  = 0
	disabledStateDisabled = 1
	disabledStateEnabled  = 2
)

// Open the transport connection and start the process
func (h *host) start() error {
	log := h.log

	// Prepare context & start message loop
	ctx, cancel := context.WithCancel(context.Background())
	h.cancelQueue = cancel
	go h.runMessageQueue(ctx)

	// Prepare transport connection
	if sCfg := h.Serial; sCfg != nil {
		// Connect using serial port
		conn, err := serial.New(*sCfg, h.log, h.parseAndQueue)
		if err != nil {
			return fmt.Errorf("host failed to initialize serial port: %w", err)
		}
		h.conn = conn
	} else {
		// No other transport protocol available
		cancel()
		return fmt.Errorf("no transport protocol configured")
	}

	// Build interface node
	h.intfNode = newNode(bidib.InterfaceAddress(), h, h.conn, log)

	// Disable all communication
	log.Debug().Msg("Disabling interface...")
	if err := h.conn.SendMessages([]bidib.Message{messages.SysReset{}}, 0); err != nil {
		cancel()
		return fmt.Errorf("failed to disable interface: %w", err)
	}

	// Get basic information of interface node
	log.Debug().Msg("Getting basic properties of interface...")
	if err := h.intfNode.readNodeProperties(); err != nil {
		cancel()
		return fmt.Errorf("failed to get basic node properties: %w", err)
	}

	return nil
}

// Close any connections
func (h *host) Close() error {
	var err error
	if conn := h.conn; conn != nil {
		h.conn = nil
		err = conn.Close()
	}
	if cancel := h.cancelQueue; cancel != nil {
		cancel()
	}
	return err
}

// Returns the root of the node tree
func (h *host) GetRootNode() *Node {
	return h.intfNode
}

// Gets the node with the given address.
// Returns nil, false if not found
func (h *host) GetNode(addr bidib.Address) (*Node, bool) {
	n := h.intfNode
	for idx := 0; idx < 4; idx++ {
		if addr[idx] == 0 {
			// We found our node
			return n, true
		}
		// Go to child nodes
		childFound := false
		for _, child := range n.table.children {
			if child != nil && child.Address.EqualsOrContains(addr) {
				n = child
				childFound = true
				break
			}
		}
		if !childFound {
			return nil, false
		}
	}
	return n, true
}

// Register a callback that gets invoked on every node change
func (h *host) RegisterNodeChanged(handler func(*Node)) context.CancelFunc {
	return h.nodeChangedEvent.Register(handler)
}

// Call all node changed handlers
func (h *host) invokeNodeChanged(n *Node) {
	h.log.Debug().Str("addr", n.Address.String()).Msg("invokeNodeChanged")
	h.nodeChangedEvent.Invoke(n)
}

// Send a DISABLE message to the interface, blocking spontaneous messages.
// Returns true if a DISABLE message was send, false is interface was already disabled.
func (h *host) disableSpontaneousMessages() bool {
	if atomic.SwapInt32(&h.disabledState, disabledStateDisabled) != disabledStateDisabled {
		h.intfNode.sendMessages(messages.SysDisable{})
		return true
	}
	return false
}

// Send a ENABLE message to the interface, blocking spontaneous messages.
// Returns true if a ENABLE message was send, false is interface was already enabled.
func (h *host) enableSpontaneousMessages() bool {
	if atomic.SwapInt32(&h.disabledState, disabledStateEnabled) != disabledStateEnabled {
		h.intfNode.sendMessages(messages.SysEnable{})
		return true
	}
	return false
}
