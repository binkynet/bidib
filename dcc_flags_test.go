package bidib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDccFormatGetSet(t *testing.T) {
	f := make(DccFlags, 4)
	assert.False(t, f.Get(0))
	assert.False(t, f.Get(1))
	assert.False(t, f.Get(2))
	assert.False(t, f.Get(3))
	assert.False(t, f.Get(4))

	assert.NoError(t, f.Set(0, true))
	assert.NoError(t, f.Set(1, false))
	assert.NoError(t, f.Set(2, true))
	assert.NoError(t, f.Set(3, false))
	assert.Error(t, f.Set(4, false))

	assert.True(t, f.Get(0))
	assert.False(t, f.Get(1))
	assert.True(t, f.Get(2))
	assert.False(t, f.Get(3))
	assert.False(t, f.Get(4))
}

func TestDccFormatGenerateBits(t *testing.T) {
	f := DccFlags{true, false, false, true, true, true, true, false, true, true} // FL .. F9
	assert.Equal(t, uint8(0b00001001), f.GenerateBits(0, 3))
	assert.Equal(t, uint8(0b00001101), f.GenerateBits(6, 9))
	assert.Equal(t, uint8(0b10111100), f.GenerateBits(1, 9))
}

func TestDccFormatSetBits(t *testing.T) {
	f := make(DccFlags, 10)
	f.SetBits(1, 4, 0b1010)
	assert.False(t, f.Get(0)) // Not set
	assert.False(t, f.Get(1))
	assert.True(t, f.Get(2))
	assert.False(t, f.Get(3))
	assert.True(t, f.Get(4))
}
