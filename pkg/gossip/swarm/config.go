package swarm

import (
	"sync"
	"time"
)

const (
	// DefaultRetryDelay ...
	DefaultRetryDelay = 10 * time.Second
	// DefaultMinPeers ...
	DefaultMinPeers = 1
	// DefaultMaxPeers ...
	DefaultMaxPeers = 10
)

// Config ...
type Config struct {
	sync.Once
	MinPeers   uint32
	MaxPeers   uint32
	RetryDelay time.Duration
}

// Init ...
func (c *Config) Init() {

	c.Do(func() {
		if c.MinPeers == 0 {
			c.MinPeers = DefaultMinPeers
		}
		if c.MaxPeers == 0 {
			c.MaxPeers = DefaultMaxPeers
		}
		if c.RetryDelay == 0 {
			c.RetryDelay = DefaultRetryDelay
		}
	})
}

// DefaultConfig ...
func DefaultConfig() *Config {

	return &Config{
		MinPeers:   DefaultMinPeers,
		MaxPeers:   DefaultMaxPeers,
		RetryDelay: DefaultRetryDelay,
	}
}

// State ...
type State uint32

const (
	// Initialized ...
	Initialized State = iota
	// Running ...
	Running
	// Shutdown ...
	Shutdown
)

// String ...
func (s State) String() string {

	switch s {
	case Initialized:
		return "swarm-initialized"
	case Running:
		return "swarm-running"
	case Shutdown:
		return "swarm-shutdown"
	default:
		return "swarm-unknown"
	}
}
