package swarm

import (
	"log"
	"sync"
	"time"

	"github.com/palantir/stacktrace"
)

// Config ...
type Config struct {
	MinPeers   int
	MaxPeers   int
	RetryDelay time.Duration
}

// Swarm ...
type Swarm struct {
	lock sync.Mutex
	// used for logging
	logger *log.Logger
	// used to store swarm config
	conf       *Config
	shutdownCh <-chan struct{}
}

// New ...
func New(logger *log.Logger, conf *Config) (*Swarm, error) {
	if logger == nil {
		err := stacktrace.NewError("cannot create a swarm manager since passed logger was nil")
		return nil, err
	}
	if conf == nil {
		err := stacktrace.NewError("cannot create a swarm manager since passed config struct was nil")
		return nil, err
	}
	result := &Swarm{
		shutdownCh: make(chan struct{}),
		logger:     logger,
		conf:       conf,
	}
	return result
}

// Shutdown ...
func (s *Swarm) Shutdown() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.logger.Printf("[INFO] gracefully shutting down swarm manager")
	close(s.shutdownCh)
}

// ShutdownCh ...
func (s *Swarm) ShutdownCh() chan struct{} {
	return s.shutdownCh
}
