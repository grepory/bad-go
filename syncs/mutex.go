package syncs

import (
	"sync"
)

// Warning: This code has multiple problems in it, for extra
// fun!

// This is a synchronized lockable state for use in an application.
// Global state is supposedly bad, but we know what we're doing.
// It's not exported. What could go wrong? This is convenient for
// our users that need a synchronized state machine. All they have to
// do is UpdateState().
var globalState *LockableState

func init() {
	globalState = &LockableState{}
}

// "I just want to lock this object."
type LockableState struct {
	state string
	sync.Mutex
}

// Sometimes we think we want to lock an object so that we can
// share it between goroutines. It Just Makes Sense, after all.
// And look how easy Go makes this. You can just throw a mutex
// in there and use it directly.
// Remember: this needs to be a pointer, because a non-pointer
// method receiver would never actually update its state.
func (l *LockableState) MutateState(newState string) {
	l.state = newState
}

// Not everything requires a lock. At any point we can just
// inspect state, right?
func (l *LockableState) State() string {
	return l.state
}

// Now, we have a state machine that asyncrhonously takes in a value
// from a channel and updates the state once the value is sent to
// the channel.
func UpdateState(inputs chan string) {
	go func() {
		// We've left it up to the user to make sure that they lock and unlock
		// the shared state.
		globalState.Lock()
		defer globalState.Unlock()

		globalState.MutateState(<-inputs)
	}()
}

func State() string {
	return globalState.State()
}
