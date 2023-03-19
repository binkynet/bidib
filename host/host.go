package host

import (
	"fmt"

	"github.com/rs/zerolog"

	"github.com/binkynet/bidib"
	"github.com/binkynet/bidib/messages"
	"github.com/binkynet/bidib/transport"
	"github.com/binkynet/bidib/transport/serial"
)

// Host defines external interface of a Bidib host process.
type Host interface {
	// Close the connections
	Close() error
}

// Host config
type Config struct {
	Serial *serial.Config
}

// New constructs a new host process with given config
func New(cfg Config, log zerolog.Logger) (Host, error) {
	h := &host{
		Config: cfg,
		log:    log,
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
	log      zerolog.Logger
	conn     transport.Connection
	intfNode *Node
}

// Open the transport connection and start the process
func (h *host) start() error {
	log := h.log

	// Prepare transport connection
	if sCfg := h.Serial; sCfg != nil {
		// Connect using serial port
		conn, err := serial.New(*sCfg, h.log, h.processMessage)
		if err != nil {
			return fmt.Errorf("host failed to initialize serial port: %w", err)
		}
		h.conn = conn
	} else {
		// No other transport protocol available
		return fmt.Errorf("no transport protocol configured")
	}

	// Build interface node
	h.intfNode = newNode(bidib.InterfaceAddress(), h.conn, log)

	// Disable all communication
	log.Debug().Msg("Disabling interface...")
	if err := h.conn.SendMessages([]bidib.Message{messages.SysDisable{}}, 0); err != nil {
		return fmt.Errorf("failed to disable interface: %w", err)
	}

	// Get basic information of interface node
	log.Debug().Msg("Getting basic properties of interface...")
	if err := h.intfNode.readNodeProperties(); err != nil {
		return fmt.Errorf("failed to get basic node properties: %w", err)
	}

	return nil
}

// Close any connections
func (h *host) Close() error {
	if conn := h.conn; conn != nil {
		h.conn = nil
		return conn.Close()
	}
	return nil
}

// Gets the node with the given address.
// Returns nil, false if not found
func (h *host) GetNode(addr bidib.Address) (*Node, bool) {
	if addr.GetLength() == 0 {
		return h.intfNode, true
	}
	// TODO recurse into node
	return nil, false
}
