package agent

import (
	"fmt"
	"sync"
	"time"

	"github.com/da-moon/coe865-final/pkg/gossip/swarm"
)

const (
	// DefaultPort ...
	DefaultPort = 8080
)

// Config stores node configuration values
type Config struct {
	sync.Once
	Address        string
	BootstrapNodes []string
	MinPeers       uint32
	MaxPeers       uint32
	RetryDelay     time.Duration
}

// Init ...
func (c *Config) Init() {

	c.Do(func() {
		if len(c.Address) == 0 {
			c.Address = fmt.Sprintf("localhost:%d", DefaultPort)
		}
		if c.MinPeers == 0 {
			c.MinPeers = swarm.DefaultMinPeers
		}
		if c.MaxPeers == 0 {
			c.MaxPeers = swarm.DefaultMaxPeers
		}
		if c.RetryDelay == 0 {
			c.RetryDelay = swarm.DefaultRetryDelay
		}
	})
}

// DefaultConfig ...
func DefaultConfig() *Config {

	return &Config{
		Address:    fmt.Sprintf("localhost:%d", DefaultPort),
		MinPeers:   swarm.DefaultMinPeers,
		MaxPeers:   swarm.DefaultMaxPeers,
		RetryDelay: swarm.DefaultRetryDelay,
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
		return "Shutdown-initialized"
	case Running:
		return "Shutdown-running"
	case Shutdown:
		return "Shutdown-shutdown"
	default:
		return "Shutdown-unknown"
	}
}
