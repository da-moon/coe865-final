package swarm

import (
	"log"
	"net"
	"os"
	"sync"
	"sync/atomic"

	"github.com/da-moon/coe865-final/pkg/gossip/codec"
	"github.com/da-moon/coe865-final/pkg/gossip/core"
	"github.com/palantir/stacktrace"
)

// Swarm ...
type Swarm struct {
	lock sync.Mutex
	// used for logging
	logger *log.Logger
	// used to store swarm config
	conf        *Config
	coreConf    *core.Config
	peerCounter uint32
	peers       map[uint32]*Peer
	shutdownCh  chan struct{}
	state       State
}

// New ...
func New(conf *Config, coreConf *core.Config) *Swarm {

	coreConf.Init()
	conf.Init()
	logger := coreConf.Logger
	if logger == nil {
		logOutput := coreConf.LogOutput
		if logOutput == nil {
			logOutput = os.Stderr
		}
		logger = log.New(logOutput, "", log.LstdFlags)
	}
	// // fmt.Println("[INFO] returning swarm")
	result := &Swarm{
		shutdownCh:  make(chan struct{}),
		logger:      logger,
		conf:        conf,
		coreConf:    coreConf,
		peers:       make(map[uint32]*Peer),
		state:       Initialized,
		peerCounter: 0,
	}
	return result
}

// Start ...
func (s *Swarm) Start() {

	s.logger.Printf("[INFO] swarm: starting ")
	s.state = Running
	// handle any background tasks here ...
}

// AddPeer is called on swarm by gossip Shutdown
// so that swarm can go through the process of adding
// it to its list
func (s *Swarm) AddPeer(conn net.Conn) error {

	s.lock.Lock()
	defer s.lock.Unlock()
	if s.state != Running {
		err := stacktrace.NewError("swarm cannot add a new peer since its state is '%s' . ", s.state.String())
		return err
	}
	// make sure it can accept more peers
	counter := atomic.AddUint32(&s.peerCounter, 1)
	if counter > s.conf.MaxPeers {
		err := stacktrace.NewError("swarm configuration does not allow for adding a new peer. Max number of peers '%d'", s.conf.MaxPeers)
		return err
	}
	s.logger.Printf("[INFO] swarm: creating peer stub for '%s'", conn.LocalAddr().String())
	peer := &Peer{
		codec: codec.NewJSONCodec(conn, conn),
	}
	s.peers[counter] = peer
	return nil
}

// Shutdown ...
func (s *Swarm) Shutdown() error {

	s.lock.Lock()
	defer s.lock.Unlock()
	s.logger.Printf("[INFO] swarm: gracefully shutting down ...")
	if s.state == Shutdown {
		return nil
	}
	s.state = Shutdown
	close(s.shutdownCh)
	return nil
}

// ShutdownCh ...
func (s *Swarm) ShutdownCh() <-chan struct{} {

	return s.shutdownCh
}
