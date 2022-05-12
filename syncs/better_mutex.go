package syncs

import "sync"

// Concurrency introduces a whole class of problems that often
// have to be solved with synchronization primitives like the mutex.
// Let's rewrite our little state wrapper with a couple of changes:
// * A private mutex
// * An RWMutex
//
// The RWMutex allows multiple readers to observe state with a semaphore
// for reading, but then an exclusive lock for writing.
//
// https://pkg.go.dev/sync#RWMutex explains how RWMutexes work.

type BetterLockableState struct {
	state string
	mtx   sync.RWMutex // Remember, no pointers, becase no copying sync package objects.
}

// MutateState changes state of a BetterLockableState. This operation
// blocks until the state can be updated. We've also removed the channel,
// because you should never use channels in your public API--as a general
// rule of thumb. If you find yourself using channels in your public API,
// you should probably take a very close look at what you've designed and
// think about what you've done.
func (b *BetterLockableState) MutateState(newState string) {
	b.mtx.Lock()
	defer b.mtx.Unlock() // always immediately defer
	b.state = newState
}

// State reeturns the current state of the BetterLockableState.
func (b *BetterLockableState) State() string {
	b.mtx.RLock()
	defer b.mtx.RUnlock()
	return b.state
}
