package host

import (
	"github.com/binkynet/bidib"
	"github.com/binkynet/bidib/messages"
)

// Process the given message
func (h *host) processMessage(mType bidib.MessageType, addr bidib.Address, seqNum bidib.SequenceNumber, data []byte) {
	h.log.Debug().
		Str("type", mType.String()).
		Str("addr", addr.String()).
		Str("num", seqNum.String()).
		Msg("ProcessMessage")
	pm, err := messages.Parse(mType, addr, seqNum, data)
	if err != nil {
		h.log.Warn().Err(err).Msg("failed to parse message")
	}
	h.log.Debug().Msgf("parse message %s", pm)
}
