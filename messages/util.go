package messages

import (
	"encoding/binary"
	"fmt"
)

// validateDataLength checks the length of the given data against the given expected length
func validateDataLength(data []byte, expectedLength int) error {
	l := len(data)
	if l != expectedLength {
		return fmt.Errorf("invalid data length; got %d, expected %d", l, expectedLength)
	}
	return nil
}

// validateMinDataLength checks the length of the given data against the given expected length
func validateMinDataLength(data []byte, expectedMinLength int) error {
	l := len(data)
	if l < expectedMinLength {
		return fmt.Errorf("invalid data length; got %d, expected >= %d", l, expectedMinLength)
	}
	return nil
}

// readUint16 read a 16-bit value from the given data slice
func readUint16(data []byte) uint16 {
	return binary.LittleEndian.Uint16(data)
}

// readUint32 read a 32-bit value from the given data slice
func readUint32(data []byte) uint32 {
	return binary.LittleEndian.Uint32(data)
}

// writeUint16 writes a given 16-bit value into the given data slice
func writeUint16(data []byte, value uint16) {
	binary.LittleEndian.PutUint16(data, value)
}

// writeUint32 writes a given 32-bit value into the given data slice
func writeUint32(data []byte, value uint32) {
	binary.LittleEndian.PutUint32(data, value)
}
