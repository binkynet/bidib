package bidib

import "fmt"

// DCC Flags
// Index 0 = FL (lights)
// Index 1 = F1
// ...
type DccFlags []bool

// Clone creates an identical copy
func (f DccFlags) Clone() DccFlags {
	result := make(DccFlags, len(f))
	copy(result, f)
	return result
}

// Gets the flag at given index
func (f DccFlags) Get(index int) bool {
	if index >= 0 && index < len(f) {
		return f[index]
	}
	return false
}

// Sets the flag at given index
func (f DccFlags) Set(index int, value bool) error {
	if index >= 0 && index < len(f) {
		f[index] = value
		return nil
	}
	return fmt.Errorf("dcc flags index %d out of range [0-%d]", index, len(f)-1)
}

// GenerateBits turns boolean flags into bit fields.
// The flag with the start index will be the lsb, increasing 1 bit at a time.
func (f DccFlags) GenerateBits(startIndex, endIndex int) uint8 {
	result := uint8(0)
	bit := uint8(1)
	for i := startIndex; i <= endIndex; i++ {
		if f.Get(i) {
			result |= bit
		}
		bit <<= 1
	}
	return result
}

// SetBits turns bit field into boolean flags
// The flag with the start index will be the lsb, increasing 1 bit at a time.
func (f DccFlags) SetBits(startIndex, endIndex int, value uint8) {
	bit := uint8(1)
	for i := startIndex; i <= endIndex; i++ {
		if value&bit != 0 {
			f.Set(i, true)
		} else {
			f.Set(i, false)
		}
		bit <<= 1
	}
}
