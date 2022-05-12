package slices

// It's not uncommon to want to reuse data in one slice within
// another slice. "I only need the first/last n elements".

// Return the last four elements of a slice or the slice if the
// length is less than four.
func LastFour(slice []string) []string {
	if len(slice) <= 4 {
		return slice
	}

	return slice[len(slice)-4:]
}
