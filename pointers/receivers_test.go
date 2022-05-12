package pointers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSeenNow(t *testing.T) {
	// We don't need all these pesky fields, really.
	c := Certificate{}

	c.SeenNow()
	// Assert that LastSeen (which should be roughly right now) is greater
	// than 0 Unix Time (1969-12-31)
	assert.GreaterOrEqual(t, c.LastSeen(), time.Unix(0, 0))
}

// Our test fails!
//
// Error:      	"0001-01-01 00:00:00 +0000 UTC" is not greater than or equal to "1969-12-31 19:00:00 -0500 EST"
//
// LastSeen() has a 0 value, because c.SeenNow() has a non-pointer method receiver
//
// func (c Certificate) SeenNow() time.Time {}
//
// Which works on a *copy* of the Certificate and not the original object
// itself. This is useful sometimes, but if you know your method has to
// modify the receiver, then you should use a pointer:
//
// func (c *Certificate) SeenNow() time.Time {}
//
// But be careful when you do this. Go does implicit referencing and dereferencing.
// If you use the above signature, and you don't also explicitly use pointers in your
// code, it can be confusing and unclear whether or not you mean to be using
// pointers or values.
//
// The clearest way to write this would be:
//
// c := &Certificate{}
// c.SeenNow()

func TestBetterSeenNow(t *testing.T) {
	// We don't need all these pesky fields, really.
	c := Certificate{}

	c.BetterSeenNow()
	// Assert that LastSeen (which should be roughly right now) is greater
	// than 0 Unix Time (1969-12-31)
	assert.GreaterOrEqual(t, c.LastSeen(), time.Unix(0, 0))
}
