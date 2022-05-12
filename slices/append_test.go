package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Let's make sure we're increasing capacity and getting what we want!
func TestAppendNameIncreasesCapacity(t *testing.T) {
	rawNames := []string{"Greg", "Miguel", "Robbie"}
	expected := []string{"greg", "miguel", "robbie"}
	actual := make([]string, 0, 2)

	for _, v := range rawNames {
		AppendName(actual, v)
	}

	assert.Equal(t, expected, actual)
}

/*

Oh no! Our test fails. Why? Aren't slices reference types?

expected: []string{"greg", "miguel", "robbie"}
actual  : []string{"", ""}

Well, no. They aren't!

what even is a slice?
   a slice is defined like so:

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

The "array" variable is a pointer to the beginning of the array and points to
a region in memory of size(T)*cap where T is the type. So if you are using an int,
then each element is an 8-byte chunk of memory. To access an element in an array, Go looks
at the array pointer and does some basic math:

elementAddress = array + (sizeOf(T) * index)

What's passed to our AppendName function isn't the slice, but this struct.
When we increased the capacity of the array, that array pointer was invalidated.

The region pointed to by the array pointer is a contiguous region on the heap
of sizeOf(T) * capacity. We can't overrun that region, so when we increase capacity,
we allocate a new region of memory and assign the start of the region to the
array pointer.

So if we don't return the new array and capture its value, we will lose this array
forever.
*/

// This works! Take a look at better_append.go to see why.
func TestBetterAppendNameIncreasesCapacity(t *testing.T) {
	rawNames := []string{"Greg", "Miguel", "Robbie"}
	expected := []string{"greg", "miguel", "robbie"}
	actual := make([]string, 0, 2)

	for _, v := range rawNames {
		actual = BetterAppendName(actual, v)
	}

	assert.Equal(t, expected, actual)
}
