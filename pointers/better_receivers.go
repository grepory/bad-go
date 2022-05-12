package pointers

import "time"

// If we want to modify the Certificate in place, we need to operate
// on a pointer and not a value.
func (c *Certificate) BetterSeenNow() time.Time {
	c.lastSeen = time.Now()
	return c.lastSeen
}

func (c *Certificate) BetterLastSeen() time.Time {
	return c.lastSeen
}
