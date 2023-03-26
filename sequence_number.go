package bidib

import (
	"strconv"
)

// Message sequence number
type SequenceNumber uint8

// String returns a human readable representation of a MessageType.
func (sn SequenceNumber) String() string {
	return strconv.Itoa(int(sn))
}

// Next returns the next sequence number.
func (sn SequenceNumber) Next() SequenceNumber {
	if sn == 255 {
		return 1
	}
	return sn + 1
}

// Reset the sequence number to 0.
func (sn *SequenceNumber) Reset() {
	*sn = 0
}
