package loops

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMembers(t *testing.T) {
	r := &Role{
		Name: "admin",
	}

	people := []Person{{"John", "Doe"}, {"Jane", "Doe"}}
	r.AddMembers(people...)

	// We know that role.Members() is going to return an array of pointers
	// so let's make an array to test against
	members := r.Members()
	var actual []Person
	for _, m := range members {
		actual = append(actual, *m)
	}
	assert.Equal(t, people, actual)
}

// Oh no. Why do we just have two Jane Does?
//
// Let's head over to better_loop.go to find out.

func TestTestBetterAddMembers(t *testing.T) {
	r := &Role{
		Name: "admin",
	}

	people := []Person{{"John", "Doe"}, {"Jane", "Doe"}}
	r.BetterAddMembers(people...)

	members := r.Members()
	var actual []Person
	for _, m := range members {
		actual = append(actual, *m)
	}
	assert.Equal(t, people, actual)
}
