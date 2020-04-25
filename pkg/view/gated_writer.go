package view

import (
	"io"
	"sync"
)

// GatedWriter ...
type GatedWriter struct {
	Writer io.Writer
	buf    [][]byte
	flush  bool
	lock   sync.RWMutex
}

// Flush ...
func (w *GatedWriter) Flush() {
	w.lock.Lock()
	w.flush = true
	w.lock.Unlock()
	for _, p := range w.buf {
		w.Write(p)
	}
	w.buf = nil
}

// Write ...
func (w *GatedWriter) Write(p []byte) (n int, err error) {

	w.lock.RLock()
	defer w.lock.RUnlock()
	if w.flush {
		return w.Writer.Write(p)
	}
	p2 := make([]byte, len(p))
	copy(p2, p)
	w.buf = append(w.buf, p2)
	return len(p), nil
}
