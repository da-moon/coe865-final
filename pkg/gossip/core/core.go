package core

import (
	"io"
	"log"
	"os"
	"sync"

	lamportclock "github.com/da-moon/coe865-final/pkg/gossip/lamport-clock"
)

const (
	// DefaultEventChannelSize ...
	DefaultEventChannelSize = 64
)

// Config ...
type Config struct {
	clock      lamportclock.LamportClock
	eventClock lamportclock.LamportClock
	sync.Once
	Initialized     bool
	LogOutput       io.Writer
	Logger          *log.Logger
	ShutdownCh      chan struct{}
	DevelopmentMode bool
	NodeName        string
	ExternalEventCh chan Event
	InrernalEventCh chan Event
}

// Init ...
func (c *Config) Init() {

	c.Do(func() {
		c.clock.Increment()
		c.eventClock.Increment()
		if c.ExternalEventCh == nil {
			c.ExternalEventCh = make(chan Event, DefaultEventChannelSize)
		}
		if c.InrernalEventCh == nil {
			c.InrernalEventCh = make(chan Event, DefaultEventChannelSize)
		}
		logOutput := c.LogOutput
		if logOutput == nil {
			logOutput = os.Stderr
		}
		c.Logger = log.New(logOutput, "", log.LstdFlags)
		c.Initialized = true
		// handle core events in background
		// go c.HandleCoreEvents()
	})
}

// DefaultConfig ...
func DefaultConfig() *Config {

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return &Config{
		LogOutput:       os.Stderr,
		NodeName:        hostname,
		ShutdownCh:      make(chan struct{}),
		ExternalEventCh: make(chan Event, 1024),
		InrernalEventCh: make(chan Event, 1024),
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
		return "core-initialized"
	case Running:
		return "core-running"
	case Shutdown:
		return "core-shutdown"
	default:
		return "core-unknown"
	}
}

// AgentHelloEvent ...
func (c *Config) AgentHelloEvent(name string, payload []byte) {
	// // fmt.Println("[INFO] AgentHelloEvent()")

	c.Logger.Printf("[DEBUG] core: Requesting agent hello send: %s. Payload: %#v",
		name, string(payload))
	if c.ExternalEventCh != nil {
		// // fmt.Println("[INFO] AgentHelloEvent() sending")

		c.ExternalEventCh <- HelloEvent{
			LTime:   c.eventClock.Time(),
			Name:    name,
			Payload: payload,
		}
	}
	return
}
