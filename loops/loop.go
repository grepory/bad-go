package loops

// We often need to loop over an array and do something with
// the values. Loops in Go have an interesting implementation
// detail that can cause a lot of confusion until you understand
// what's happening.

// Let's say we have some code that associates People with a
// a Role.
type Person struct {
	First string
	Last  string
}

type Role struct {
	Name    string
	members []*Person
}

// AddMembers adds one or more members to the role.
func (r *Role) AddMembers(p ...Person) {
	for _, v := range p {
		r.members = append(r.members, &v)
	}
}

func (r *Role) Members() []*Person {
	members := make([]*Person, len(r.members))
	copy(members, r.members)
	return members
}

// Let's go test it!
