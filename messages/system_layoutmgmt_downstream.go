package messages

import (
	"fmt"

	"github.com/binkynet/bidib"
)

// This command transmits a model time for layout appliances.
// This clock typically runs accelerated compared to the real time.
// Followed by 4 bytes (TCODE0, TCODE1, TCODE2, TCODE3) with the time value.
// The coding of these bytes is the same as the coding of the corresponding DCC command.
type SysClock struct {
	BaseMessage
	Minutes      uint8
	Hours        uint8
	Weekday      uint8
	Acceleration uint8
}

func (m SysClock) Encode(write func(uint8), seqNum bidib.SequenceNumber) {
	data := []byte{
		m.Minutes, /*| 0b00000000*/
		m.Hours | 0b10000000,
		m.Weekday | 0b01000000,
		m.Acceleration | 0b11000000,
	}
	bidib.EncodeMessage(write, bidib.MSG_SYS_CLOCK, m.Address, seqNum, data)
}

func (m SysClock) String() string {
	return fmt.Sprintf("%T addr=%s time=%d:%d weekday=%d accel=%d", m, m.Address, m.Hours, m.Minutes, m.Weekday, m.Acceleration)
}

func decodeSysClock(addr bidib.Address, data []byte) (SysClock, error) {
	var result SysClock
	if err := validateDataLength(data, 4); err != nil {
		return result, err
	}
	result.Address = addr
	result.Minutes = data[0] & 0b00111111
	result.Hours = data[1] & 0b00011111
	result.Weekday = data[2] & 0b00000111
	result.Acceleration = data[3] & 0b00111111
	return result, nil
}
