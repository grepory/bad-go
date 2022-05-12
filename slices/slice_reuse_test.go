package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastFour(t *testing.T) {
	// Our array.
	arr := []string{"a", "b", "c", "d", "e", "f"}

	// Just want the last four
	newArr := LastFour(arr)

	// Now let's make sure that worked.
	expected := []string{"c", "d", "e", "f"}
	assert.Equal(t, expected, newArr)

	// But what if we do something like this?
	arr[len(arr)-1] = "a"
	assert.Equal(t, expected, newArr)
}

// Whaaaaat?!
// The last element of our *NEW* slice was changed as well.
// If you have not read the explanation of slices in append_test.go,
// go do that now.
//
// Recall that a slice is a header that contains a pointer to the start
// of a region of memory. Because our assignment didn't require an
// allocation of a new region of memory, we reused the memory from the
// original array. So the unsafe.Pointer from arr is the same unsafe.Pointer
// in the header of the new array.
//
// So changing data in the original memory location arr[n] changes it for
// both arr and newArr.

// So let's test the better version.
func TestBetterLastFour(t *testing.T) {
	// Our array.
	arr := []string{"a", "b", "c", "d", "e", "f"}

	// Just want the last four
	newArr := BetterLastFour(arr)

	// Now let's make sure that worked.
	expected := []string{"c", "d", "e", "f"}
	assert.Equal(t, expected, newArr)

	// But what if we do something like this?
	arr[len(arr)-1] = "a"
	assert.Equal(t, expected, newArr)
}
