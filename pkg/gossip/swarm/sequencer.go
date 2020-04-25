package swarm

import "fmt"

type sequencer uint
type sequenceTracker struct {
	horizon sequencer
	seen    map[sequencer]bool
}

// See ...
func (t *sequenceTracker) See(s sequencer) bool {
	fmt.Println("[CALL-STACK] [file] /workspace/coe865-final/internal/overlay/sequence.go [method] (t *sequenceTracker) See(s sequencer) bool ")
	if s < t.horizon || t.seen[s] {
		return false
	}
	if t.seen == nil {
		t.seen = make(map[sequencer]bool)
	}
	t.seen[s] = true
	for t.seen[t.horizon] {
		delete(t.seen, t.horizon)
		t.horizon++
	}
	return true
}
