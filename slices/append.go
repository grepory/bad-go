package slices

import "strings"

/*
Let's say you're looping through a list of strings, and you want to normalize
them to all lower-case. You might do something like this:
*/

func AppendName(names []string, name string) {
	// let's make sure we don't overflow our array!
	length := len(names)
	if length == cap(names) {
		newNames := make([]string, length, 2*length+1)
		// get all the stuff out of slice into newslice
		copy(newNames, names)
		names = newNames
	}
	// in case we didn't have to increase capacity.
	names = names[0 : length+1]
	// add our thing!
	names[length] = strings.ToLower(name)
}

// Let's go write a test!
