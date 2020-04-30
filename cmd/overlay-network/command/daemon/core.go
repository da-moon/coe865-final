package daemon

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"

	"github.com/da-moon/coe865-final/internal/swarm"
	model "github.com/da-moon/coe865-final/model"
	config "github.com/da-moon/coe865-final/pkg/config"
	"github.com/da-moon/coe865-final/pkg/jsonutil"

	utils "github.com/da-moon/coe865-final/pkg/utils"
	stacktrace "github.com/palantir/stacktrace"
	cron "github.com/robfig/cron/v3"
)

type identityRecord struct {
	Identity             Identity
	AgentSequenceTracker sequenceTracker
}

// Core ...
type Core struct {
	lock          *sync.Mutex
	logger        *log.Logger
	shutdown      bool
	shutdownCh    chan struct{}
	shutdownLock  sync.Mutex
	cron          *cron.Cron
	conf          *config.Config
	key           *rsa.PrivateKey
	gossiper      swarm.Gossiper
	mu            sync.Mutex
	joinsByAddr   map[string]Message
	identities    map[fingerprint]*identityRecord
	agentSequence sequencer
	peerManager   swarm.PeerManager
	listener      net.Listener
}

// Create ...
func Create(conf *config.Config, logOutput io.Writer) (*Core, error) {

	if logOutput == nil {
		logOutput = os.Stderr
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", conf.Port))
	if err != nil {
		err = stacktrace.Propagate(err, "could not bind to port '%d'", conf.Port)
		return nil, err
	}
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		err = stacktrace.Propagate(err, "could not start core due to an issue with generating rsa key")
		return nil, err
	}

	logger := log.New(logOutput, "", log.LstdFlags)
	result := &Core{
		conf:        conf,
		logger:      logger,
		listener:    listener,
		shutdownCh:  make(chan struct{}),
		key:         key,
		joinsByAddr: make(map[string]Message),
		identities:  make(map[fingerprint]*identityRecord),
		cron: cron.New(
			cron.WithLogger(cron.PrintfLogger(logger)),
		),
	}

	result.gossiper = swarm.NewGossiper(result.handleGossip, result.logger)
	pmConf := swarm.PeerManagerConfig{
		NewPeer:   result.findPeer,
		OnConnect: result.onConnect,
		MinPeers:  1,
	}
	result.peerManager = swarm.NewPeerManager(result.gossiper, pmConf, result.logger)
	return result, nil
}

// Start makes the gossip agent
// listening to incomming connection and
// establish connection with bootstrap nodes in case it is not
// in dev mode
// it also starts background event handler and
// scheduler.
func (a *Core) Start() error {
	a.logger.Printf("[INFO] overlay network daemon core started!")
	a.logger.Printf("[INFO] creating graph based on config !")
	entryID, err := a.cron.AddFunc(a.conf.Cron, a.SendUpdateMessage())
	if err != nil {
		err = stacktrace.Propagate(err, "could not start cron job handler")
		a.logger.Printf(fmt.Sprintf(("[WARN] error : %#v"), err.Error()))
		return err
	}
	a.logger.Printf("[INFO] cron job entry ID %v", entryID)
	a.cron.Start()
	go a.listen()
	// TODO fix dev flag [CRITICAL]
	if len(a.conf.ConnectedRouteControllers) > 0 {
		a.logger.Printf("[INFO] '%s' : connecting to bootstrap nodes ", a.listener.Addr().String())
		a.bootstrap()
	}
	return nil
}

// listen spins a new goroutine
// per incomming connection as it is waiting for nodes
// to join . it then passes
// the listener connectiontion
// to swarm for handling comminucations to/from it
func (a *Core) listen() {
	a.logger.Printf("[INFO]  started listening for incomming connections on port '%d' ...", a.conf.Port)
	defer a.listener.Close()
	peerManager := a.peerManager

	for {
		select {
		case <-a.shutdownCh:
			return
		default:
			{
				if a.shutdown {
					a.logger.Printf("[WARN]  cannot accept any more incomming connection since it has shutdown")
				}
				conn, err := a.listener.Accept()
				if err != nil {
					a.logger.Printf("[WARN] listener failed to accept new connection : %v", err)
					continue
				}
				a.logger.Printf("[INFO] agent : recieved an incomming connection from peer with address %v", conn.LocalAddr().String())
				peer := NewPeer(conn)
				if err := peerManager.AddPeer(peer); err != nil {
					a.logger.Printf("[WARN] Error adding new peer %s: %s", peer, err)
				}
			}
		}
	}

}

// bootstrap makes the agent establish connections to a set of node address
// passed to it at the time it starts
// we assume the bootstrap node also are listenning
// to the same port agent is listening (as it is set in it's config)
func (a *Core) bootstrap() {
	a.logger.Printf("[INFO]  connecting to bootstrap nodes ...")
	peerManager := a.peerManager
	for _, rc := range a.conf.ConnectedRouteControllers {
		addr := rc.IP
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			a.logger.Printf("[WARN] could not connect to '%s' : %v", addr, err)
			continue
		}
		a.logger.Printf("[INFO] agent : established an outgoing connection to peer with address %v", conn.RemoteAddr().String())
		peer := NewPeer(conn)
		if err := peerManager.AddPeer(peer); err != nil {
			a.logger.Printf("[WARN] Error adding new peer %s: %s", peer, err)
		}
	}
}

// connect establishes gossip sessions with the given addresses,
// handing them off to peerManager once established.
func connect(rcs []config.RouteController, peerManager swarm.PeerManager) {

	for _, rc := range rcs {
		addr := rc.IP
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Fatalf("Error dialing initial peer %s: %s", addr, err)
		}
		peer := NewPeer(conn)
		if err := peerManager.AddPeer(peer); err != nil {
			log.Fatalf("Error adding initial peer %s: %s", peer, err)
		}
	}
}

// Shutdown ...
func (a *Core) Shutdown() error {

	a.shutdownLock.Lock()
	defer a.shutdownLock.Unlock()
	a.cron.Stop()
	a.gossiper.Close()
	a.shutdown = true
	close(a.shutdownCh)
	a.logger.Println("[INFO]", "overlay network daemon core: shutdown complete")
	return nil
}

// ShutdownCh ...
func (a *Core) ShutdownCh() <-chan struct{} {

	return a.shutdownCh
}

// SendUpdateMessage ...
func (a *Core) SendUpdateMessage() func() {

	return func() {
		req := &model.UpdateMessage{
			UUID: utils.UUID(),
		}
		req.SourceRouteController = &model.RouteController{
			ID:                     int32(a.conf.Self.ID),
			AutonomousSystemNumber: int32(a.conf.Self.AutonomousSystemNumber),
			IP:                     a.conf.Self.IP,
		}
		for _, v := range a.conf.ConnectedAutonomousSystems {
			req.DestinationAutonomousSystem = append(req.DestinationAutonomousSystem, &model.AutonomousSystem{
				Number:       int32(v.Number),
				LinkCapacity: int32(v.LinkCapacity),
				Cost:         int32(v.Cost),
			})
		}
		a.Broadcast(base64.StdEncoding.EncodeToString(jsonutil.EncodeJSONWithoutErr(req)))
	}
}
