package loops

// BetterAddMembers really adds one or more members to the role.
func (r *Role) BetterAddMembers(p ...Person) {
	// In Go, the iterator values, in our case "v", are reused each
	// iteration. If you reference the iterator value itself, that
	// pointer will always point to the same location in memory.
	//
	// So when you do:
	// for _, v := range p {
	//     r.members = append(r.members, &v)
	// }
	//
	// You're actually referencing "v" itself and NOT the value.
	// At the time that the iteration is executed, the result is what
	// you expect, but it's not DOING what you probably expect.
	// By *capturing* the value of v during the iteration and referencing
	// that (in a value that IS NOT reused during each iteration), you can
	// reference the value itself and not the variable v.
	for _, v := range p {
		// Capture the value of v in a variable that isn't reused between iterations
		member := v
		// Reference THAT.
		r.members = append(r.members, &member)
	}
}
