package slices

import "strings"

// This fix is a lot more obvious than the previous version.
// There wasn't really anything wrong with our original code!
// All we need to do is return the resulting array.

func BetterAppendName(names []string, name string) []string {
	// let's make sure we don't overflow our array!
	length := len(names)
	if length+1 > cap(names) {
		newNames := make([]string, length, 2*(length+1))
		// get all the stuff out of slice into new slice
		copy(newNames, names)
		// So that the rest of the code works, we now reassign
		// names to our new slice.
		names = newNames
	}
	// in case we didn't have to increase capacity.
	names = names[0 : length+1]
	// add our thing! reminder: arrays are 0-indexed, so length will be
	// 1 past the last element in the array.
	names[length] = strings.ToLower(name)

	return names
}
