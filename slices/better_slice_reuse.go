package slices

// We have to make sure that we're using a different
// slice if we want to make sure that the data is
// actually separated. In general, this is what you should
// always do--otherwise you may inadvertently introduce
// a bug later if the kind of behavior observed in LastFour
// makes its way into the code. Don't worry about the original
// array needing to be GC'd. If the array is so large that you're
// worried about allocating a second version, then it's probably time
// to rethink data structure or the architecture of the program.
func BetterLastFour(slice []string) []string {
	// Honor the original contract.
	if len(slice) <= 4 {
		return slice
	}

	lastFour := make([]string, 4)
	copy(lastFour, slice[len(slice)-4:])
	return lastFour
}
