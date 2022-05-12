package pointers

import "time"

// We use structs to represent the data in our domain.
type Certificate struct {
	Serial     int
	CommonName string
	NotAfter   time.Time

	// lastSeen is the last time our monitoring script successfully contacted
	// an endpoint using this certificate. We don't want people setting this
	// once it's set, so we'll use a private field combined with a getter.
	lastSeen time.Time
}

// Sometimes, values in struct fields may change due to application
// behavior.
func (c Certificate) SeenNow() time.Time {
	c.lastSeen = time.Now()
	return c.lastSeen
}

func (c Certificate) LastSeen() time.Time {
	return c.lastSeen
}
