package core

type Core interface {
}
type core struct {
	// mutex , used for safeguarding struct
	lock sync.Mutex
	// is set to true when node is
	// about to shut down
	gracefulShutdown bool
	// is set to true if node is shutdown
	shutdownCh chan bool
}

// New creates a new gossip core
func New() Core {
	c := &core{
		gracefulShutdown: false,
		shutdownCh:       make(chan bool),
	}
	return c
}
