package daemon

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"

	model "github.com/da-moon/coe865-final/model"
	config "github.com/da-moon/coe865-final/pkg/config"
	"github.com/da-moon/coe865-final/pkg/gossip/core"
	"github.com/da-moon/coe865-final/pkg/gossip/swarm"
	utils "github.com/da-moon/coe865-final/pkg/utils"
	prettyjson "github.com/hokaccha/go-prettyjson"
	stacktrace "github.com/palantir/stacktrace"
	cron "github.com/robfig/cron/v3"
)

// Core ...
type Core struct {
	lock         *sync.Mutex
	logger       *log.Logger
	shutdown     bool
	shutdownCh   chan struct{}
	shutdownLock sync.Mutex
	cron         *cron.Cron
	// used for handling events as they are broadcasted
	coreConf          *core.Config
	conf              *config.Config
	listener          net.Listener
	swarm             *swarm.Swarm
	eventCh           chan core.Event
	eventHandlersLock sync.Mutex
	eventHandlers     map[EventHandler]struct{}
	eventHandlerList  []EventHandler
}

// Create ...
func Create(conf *config.Config, coreConf *core.Config, logOutput io.Writer) (*Core, error) {

	if logOutput == nil {
		logOutput = os.Stderr
	}
	coreConf.LogOutput = logOutput
	coreConf.DevelopmentMode = conf.DevelopmentMode
	// fmt.Println("coreConf.DevelopmentMode", coreConf.DevelopmentMode)
	// fmt.Println("conf.DevelopmentMode", conf.DevelopmentMode)
	// core.NodeName = conf.
	// Create a channel to listen for events engines
	eventCh := make(chan core.Event, core.DefaultEventChannelSize)
	coreConf.ExternalEventCh = eventCh
	coreConf.Init()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", conf.Port))
	if err != nil {
		err = stacktrace.Propagate(err, "could not bind to port '%d'", conf.Port)
		return nil, err
	}
	swarmConf := swarm.DefaultConfig()
	swarmConf.MaxPeers = uint32(conf.MaxPeers)
	swarmConf.MinPeers = uint32(conf.MinPeers)
	logger := log.New(logOutput, "", log.LstdFlags)
	agentCore := &Core{
		coreConf:      coreConf,
		conf:          conf,
		eventCh:       eventCh,
		eventHandlers: make(map[EventHandler]struct{}),
		logger:        logger,
		shutdownCh:    make(chan struct{}),

		listener: listener,
		swarm:    swarm.New(swarmConf, coreConf),
		cron: cron.New(
			cron.WithLogger(cron.PrintfLogger(logger)),
		),
	}
	return agentCore, nil
}

// Start makes the gossip agent
// listening to incomming connection and
// establish connection with bootstrap nodes in case it is not
// in dev mode
// it also starts background event handler and
// scheduler.
func (a *Core) Start() error {
	a.logger.Printf("[INFO] overlay network daemon core started!")
	entryID, err := a.cron.AddFunc(a.conf.Cron, a.EstimateCost())
	if err != nil {
		err = stacktrace.Propagate(err, "could not start cron job handler")
		a.logger.Printf(fmt.Sprintf(("[WARN] error : %#v"), err.Error()))
		return err
	}
	a.logger.Printf("[INFO] cron job entry ID %v", entryID)
	a.cron.Start()

	go a.eventHandlerLoop()
	go a.listen()
	// TODO fix dev flag [CRITICAL]
	if len(a.conf.ConnectedRouteControllers) > 0 {
		a.logger.Printf("[INFO] '%s' : connecting to bootstrap nodes ", a.listener.Addr().String())
		go a.bootstrap()
	}
	a.swarm.Start()
	return nil
}

// listen spins a new goroutine
// per incomming connection as it is waiting for nodes
// to join . it then passes
// the listener connectiontion
// to swarm for handling comminucations to/from it
func (a *Core) listen() {
	a.logger.Printf("[INFO] agent '%v' : started listening for incomming connections on port '%d' ...", a.ID(), a.conf.Port)
	defer a.listener.Close()

	for {
		select {
		case <-a.shutdownCh:
			return
		default:
			{
				if a.shutdown {
					a.logger.Printf("[WARN] agent '%v' : cannot accept any more incomming connection since it has shutdown", a.ID())
				}
				conn, err := a.listener.Accept()
				if err != nil {
					a.logger.Printf("[WARN] listener failed to accept new connection : %v", err)
					continue
				}
				a.logger.Printf("[INFO] node '%v' : recieved an incomming connection from peer with address %v", a.ID(), conn.LocalAddr().String())
				err = a.swarm.Handshake(conn)
				if err != nil {
					a.logger.Printf("[WARN] %v", err)
					// closing connection
					conn.Close()
					continue
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
	a.logger.Printf("[INFO] agent '%v' : connecting to bootstrap nodes ...", a.ID())
	for _, rc := range a.conf.ConnectedRouteControllers {
		addr := rc.IP
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			a.logger.Printf("[WARN] could not connect to '%s' : %v", addr, err)
			continue
		}
		a.logger.Printf("[INFO] node '%v' : established an outgoing connection to peer with address %v", a.ID(), conn.LocalAddr().String())
		err = a.swarm.Handshake(conn)
		if err != nil {
			a.logger.Printf("[WARN] %v", err)
			// closing connection
			conn.Close()
			continue
		}
	}
}
func (a *Core) eventHandlerLoop() {
	a.logger.Printf("[INFO] agent: started event listener")
	for {
		select {
		case e := <-a.eventCh:
			a.logger.Printf("[INFO] agent: received event: %s", e.String())
		case <-a.shutdownCh:
			return
		}
	}
}

// RegisterEventHandler adds an event handler to receive event notifications
func (a *Core) RegisterEventHandler(eh EventHandler) {

	a.eventHandlersLock.Lock()
	defer a.eventHandlersLock.Unlock()
	a.eventHandlers[eh] = struct{}{}
	a.eventHandlerList = nil
	for eh := range a.eventHandlers {
		a.eventHandlerList = append(a.eventHandlerList, eh)
	}
}

// DeregisterEventHandler removes an EventHandler and prevents more invocations
func (a *Core) DeregisterEventHandler(eh EventHandler) {

	a.eventHandlersLock.Lock()
	defer a.eventHandlersLock.Unlock()
	delete(a.eventHandlers, eh)
	a.eventHandlerList = nil
	for eh := range a.eventHandlers {
		a.eventHandlerList = append(a.eventHandlerList, eh)
	}
}

// ID returns the first 6 characters of base64
// encoded public key to be used as agent ID
func (a *Core) ID() string {

	result, err := a.coreConf.Sentry().PublicKeyBase64()
	if err != nil {
		err = stacktrace.Propagate(err, "could not get node ID")
		panic(err)
	}
	return result[:8]
}

// Shutdown ...
func (a *Core) Shutdown() error {

	a.shutdownLock.Lock()
	defer a.shutdownLock.Unlock()
	a.cron.Stop()
	err := a.swarm.Shutdown()
	if err != nil {
		err = stacktrace.Propagate(err, "cannot gracefully shutdown agent")
		return err
	}
	a.shutdown = true
	close(a.shutdownCh)
	a.logger.Println("[INFO]", "overlay network daemon core: shutdown complete")
	return nil
}

// ShutdownCh ...
func (a *Core) ShutdownCh() <-chan struct{} {

	return a.shutdownCh
}

// EstimateCost ...
func (a *Core) EstimateCost() func() {

	return func() {
		req := &model.UpdateRequest{
			UUID: utils.UUID(),
		}
		req.SourceRouteController = &model.RouteController{
			ID:                     int32(a.conf.Self.ID),
			AutonomousSystemNumber: int32(a.conf.Self.AutonomousSystemNumber),
			IP:                     a.conf.Self.IP,
		}
		req.DestinationAutonomousSystem = &model.AutonomousSystem{
			Number:       int32(a.conf.ConnectedAutonomousSystems[0].Number),
			LinkCapacity: int32(a.conf.ConnectedAutonomousSystems[0].LinkCapacity),
			Cost:         int32(a.conf.ConnectedAutonomousSystems[0].Cost),
		}
		prettyreq, _ := prettyjson.Marshal(req)
		a.logger.Println("[INFO]", "req", string(prettyreq))
	}
}
