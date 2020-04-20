package lamportclock

import (
	"sync/atomic"
)

// LamportClock ...
type LamportClock struct {
	counter uint64
}

// Time is used to return the current value of the lamport clock
func (l *LamportClock) Time() uint64 {
	return atomic.LoadUint64(&l.counter)
}

// Increment is used to increment and return the value of the lamport clock
func (l *LamportClock) Increment() uint64 {
	return atomic.AddUint64(&l.counter, 1)
}

// Witness is called to update our local clock if necessary after
// witnessing a clock value received from another process
func (l *LamportClock) Witness(v uint64) {
WITNESS:
	// If the other value is old, we do not need to do anything
	cur := atomic.LoadUint64(&l.counter)
	other := v
	if other < cur {
		return
	}

	// Ensure that our local clock is at least one ahead.
	if !atomic.CompareAndSwapUint64(&l.counter, cur, other+1) {
		// The CAS failed, so we just retry. Eventually our CAS should
		// succeed or a future witness will pass us by and our witness
		// will end.
		goto WITNESS
	}
}
