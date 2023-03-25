package bidib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddressAppend(t *testing.T) {
	a := MustNewAddress()
	assert.Equal(t, "[]", a.String())

	a = a.Append(0)
	assert.Equal(t, "[]", a.String())

	a = a.Append(1)
	assert.Equal(t, "[1]", a.String())

	a = a.Append(2)
	assert.Equal(t, "[1,2]", a.String())

	a = a.Append(3)
	assert.Equal(t, "[1,2,3]", a.String())

	a = a.Append(4)
	assert.Equal(t, "[1,2,3,4]", a.String())
}

func TestAddressEquals(t *testing.T) {
	a := MustNewAddress(1, 2)
	b := MustNewAddress(1, 2)
	c := MustNewAddress(1, 2, 3)
	assert.True(t, a.Equals(b))
	assert.True(t, b.Equals(a))
	assert.False(t, b.Equals(c))
}

func TestAddressEqualsOr(t *testing.T) {
	empty := InterfaceAddress()
	a := MustNewAddress(1, 2)
	b := MustNewAddress(1, 2)
	c := MustNewAddress(1, 2, 3)
	d := MustNewAddress(11, 12, 13)
	assert.True(t, a.EqualsOrContains(b))
	assert.True(t, b.EqualsOrContains(a))
	assert.True(t, a.EqualsOrContains(c))
	assert.False(t, a.EqualsOrContains(d))
	assert.False(t, a.EqualsOrContains(empty))
	assert.True(t, empty.EqualsOrContains(a))
}
