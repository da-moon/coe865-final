package core

import (
	"io"
	"log"
	"os"
	"sync"

	lamportclock "github.com/da-moon/coe865-final/pkg/gossip/lamport-clock"
	"github.com/da-moon/coe865-final/pkg/gossip/sentry"
	"github.com/palantir/stacktrace"
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
	sentry          *sentry.Sentry
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
		if c.sentry == nil {
			k, err := sentry.Default()
			if err != nil {
				err = stacktrace.Propagate(err, "could not create a new gossip agent core due to an issue with generating RSA key for the node")
				panic(err)

			}
			c.sentry = k
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

// Sentry ...
func (c *Config) Sentry() *sentry.Sentry {
	return c.sentry
}

// DefaultConfig ...
func DefaultConfig() *Config {

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	k, err := sentry.Default()
	if err != nil {
		err = stacktrace.Propagate(err, "could not create a new gossip agent core due to an issue with generating RSA key for the node")
		panic(err)
	}

	return &Config{
		LogOutput:       os.Stderr,
		NodeName:        hostname,
		ShutdownCh:      make(chan struct{}),
		ExternalEventCh: make(chan Event, DefaultEventChannelSize),
		InrernalEventCh: make(chan Event, DefaultEventChannelSize),
		sentry:          k,
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
