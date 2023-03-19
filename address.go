package bidib

import (
	"fmt"
	"strconv"
)

// Address is a stack up to 4 bytes.
type Address [4]byte

// InterfaceAddress returns the address of the connected interface (0)
func InterfaceAddress() Address {
	return Address{}
}

// NewAddress constructs a new address
func NewAddress(addr ...uint8) (Address, error) {
	if len(addr) > 4 {
		return Address{}, fmt.Errorf("address cannot be longer than 4 elements")
	}
	for _, x := range addr {
		if x == 0 {
			return Address{}, fmt.Errorf("address cannot contain 0")
		}
	}
	var result Address
	copy(result[:], addr)

	return result, nil
}

// GetLength returns the length of the address stack (the amount of leading non-zero elements).
func (a Address) GetLength() uint8 {
	result := uint8(0)
	for _, x := range a {
		if x == 0 {
			return result
		}
		result++
	}
	return result
}

// String converts an address into a readable string
func (a Address) String() string {
	if a[0] == 0 {
		return ""
	}
	if a[1] == 0 {
		return strconv.Itoa(int(a[0]))
	}
	if a[2] == 0 {
		return strconv.Itoa(int(a[0])) + "," + strconv.Itoa(int(a[1]))
	}
	if a[3] == 0 {
		return strconv.Itoa(int(a[0])) + "," + strconv.Itoa(int(a[1])) +
			"," + strconv.Itoa(int(a[2]))
	}
	return strconv.Itoa(int(a[0])) + "," + strconv.Itoa(int(a[1])) +
		"," + strconv.Itoa(int(a[2])) + "," + strconv.Itoa(int(a[3]))
}
