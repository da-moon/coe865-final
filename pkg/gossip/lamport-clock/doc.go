// Package lamportclock is a thread-safe implementation of a lamport clock. It
// uses efficient atomic operations for all of its functions, falling back
// to a heavy lock only if there are enough CAS failures.
// taken from
package lamportclock
