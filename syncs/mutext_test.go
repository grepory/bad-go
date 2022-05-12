package syncs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// We have written the most useful lockable state machine ever.
// Let's test it out.
func TestSynchronizedStateUpdate(t *testing.T) {
	var inputs chan string

	states := []string{
		"new",
		"doing",
		"done",
	}

	// Let's asynchronously update some state and then
	// test that it's done. Testing asynchronous code is
	// really hard.
	go func() {
		for _, v := range states {
			inputs <- v
		}
	}()

	// Now we'll throw the channel over to update state
	for range states {
		UpdateState(inputs)
	}

	// This is the worst way to test asynchronous code, and I would argue
	// that, in general, this was a terrible way to *write* asynchronous code.
	// The parts that make code asynchronous shouldn't need to be tested.
	// The parts that happen inside the asynchronous code are what we want
	// to really test. Not the asynchronous nature of the code.
	// But here we are, because we didn't write testable code.
	time.Sleep(3 * time.Second) // should take too long

	// We updated the state synchronously on the channel so we
	// should end up at done.
	assert.Equal(t, "done", State())
}

// OH NO! Our test failed.
// State was the empty string... so state never got updated! What happened?
//
// This is where I tell you there were two problems, but one was to demonstrate
// the other. First, sending to or receiving from a nil channel blocks forever.
// Super useful, right? What good are nil channels for then? Why can they be
// nil at all? I don't know. Ask Rob Pike.
//
// But what I want to point out is that embedded mutexes seem like a neat idea,
// but you should avoid requiring the caller to lock/unlock a shared resource.
// Embedding the mutex also *exposes the Lock and Unlock methods*. This is even
// more fun, because if you are locking and unlocking internally, the caller might
// think that they have to lock and unlock it.
//
// Embedded mutexes are bad code. Always use a mutex internal to the object,
// and lock and unlock the mutex yourself. It's an implementation detail.
// You should always do:
//
// type Lockable struct {
//   mutex sync.Mutex
// }
//
// Two additional notes about this:
//
// 1. Never use a pointer mutex type. Pointers can be copied and none of the objects
// in the sync package are ever to be copied.
//
// 2. Once you add a mutex to an object, you can never copy the object, because copying
// the object will copy the Mutex. And nothing in the sync package can ever be copied.
