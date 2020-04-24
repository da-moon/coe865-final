package agent

import (
	"os"

	"github.com/da-moon/coe865-final/pkg/gossip/core"
	"github.com/da-moon/coe865-final/pkg/jsonutil"

	"log"
	"net"
	"sync"

	"github.com/da-moon/coe865-final/pkg/gossip/sentry"
	"github.com/da-moon/coe865-final/pkg/gossip/swarm"
	"github.com/palantir/stacktrace"
)

// Agent wraps a key and
// it signs messages and marshalls and
// unmarshalls messages to/from the wire
type Agent struct {
	lock sync.Mutex
	// used for signing and generating ID
	sentry *sentry.Sentry
	// used by the node to accept incoming connections
	listener net.Listener
	// used to manage nodes it is directly connected to
	swarm *swarm.Swarm
	// used for logging
	logger *log.Logger
	// used to store node config
	conf *Config
	// used to store core config
	coreConf *core.Config
	// used to log warnings
	// warningCh  chan string
	shutdownCh chan struct{}
	state      State
}

// New returns a new instance of node
// it takes in a conn (net.Conn)
// so that it can create a codec to
// encode and write and read and decode
// to the said connection
func New(conf *Config, coreConf *core.Config) (*Agent, error) {
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
	// // fmt.Println("[INFO] preparing swarm conf")
	swarmConf := &swarm.Config{
		MinPeers:   conf.MinPeers,
		MaxPeers:   conf.MaxPeers,
		RetryDelay: conf.RetryDelay,
	}
	// // fmt.Println("[INFO] getting sentry")
	k, err := sentry.Default()
	if err != nil {
		err = stacktrace.Propagate(err, "could not create a new gossip agent due to an issue with generating RSA key for the node")
		return nil, err
	}
	// // fmt.Println("[INFO] getting listener")
	listener, err := net.Listen("tcp", conf.Address)
	if err != nil {
		err = stacktrace.Propagate(err, "could not create a new listener for node with address '%s'", conf.Address)
		return nil, err
	}
	// // fmt.Println("[INFO] returning result listener")
	result := &Agent{
		sentry:   k,
		conf:     conf,
		coreConf: coreConf,
		logger:   logger,
		listener: listener,
		swarm:    swarm.New(swarmConf, coreConf),
		state:    Initialized,
		// warningCh:  make(chan string),
		shutdownCh: make(chan struct{}),
	}
	return result, nil
}

// Start makes the agent start
// listening to incomming connection and
// establish connection with bootstrap nodes
func (a *Agent) Start() {

	// agent listens to incomming connections in the background
	a.logger.Printf("[INFO] '%s' : started listening for incomming connections", a.listener.Addr().String())
	// // fmt.Println("[INFO] agent. Start()")
	go a.listen()
	if len(a.conf.BootstrapNodes) > 0 {
		a.logger.Printf("[INFO] '%s' : connecting to bootstrap nodes", a.listener.Addr().String())
		a.connect(a.conf.BootstrapNodes)
	}
	// // fmt.Println("[INFO] agent.Start() About to send Hello MSG")
	a.coreConf.AgentHelloEvent("my-hello-event", jsonutil.EncodeJSONWithoutErr(map[string]string{
		"key": "value",
	}))
	// // checking for warnings
	// for {
	// 	select {
	// 	case warn := <-a.warningCh:
	// 		{
	// 			a.logger.Printf("[WARN] %s", warn)
	// 		}
	// 	case <-a.shutdownCh:
	// 		{
	// 			a.logger.Printf("[INFO] graceful shutdown completed. tearing down ...")
	// 			return
	// 		}
	// 	}
	// }
}

// listen spins a new goroutine
// per incomming connection as it is waiting for nodes
// to join . it then passes
// the listener connectiontion
// to swarm for handling comminucations to/from it
func (a *Agent) listen() {
	// // fmt.Println("[INFO] agent.listen()")

	defer a.listener.Close()
	for {
		conn, err := a.listener.Accept()
		if err != nil {
			// warn := fmt.Sprintf("listener failed to accept new connection : %v", err)
			// a.warningCh <- warn
			continue
		}
		a.logger.Printf("[INFO] node '%v' : recieved an incomming connection from peer with address %v", a.ID(), conn.LocalAddr().String())
		err = a.swarm.AddPeer(conn)
		if err != nil {
			// a.warningCh <- err.Error()
			// closing connection
			conn.Close()
			continue
		}
	}
}

// Connect makes the agent establish connections to a set of node address
func (a *Agent) connect(addrs []string) {

	for _, addr := range addrs {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			// warn := fmt.Sprintf("could not connect to '%s' : %v", addr, err)
			// a.warningCh <- warn
			continue
		}
		a.logger.Printf("[INFO] node '%v' : established an outgoing connection to peer with address %v", a.ID(), conn.LocalAddr().String())
		// peer := NewPeer(conn)
		// if err := peerManager.AddPeer(peer); err != nil {
		// 	log.Fatalf("Error adding initial peer %s: %s", peer, err)
		// }
	}
}

// ID returns node ID which
// is base64 encoded form of it's public
// key
func (a *Agent) ID() string {

	result, err := a.sentry.PublicKeyBase64()
	if err != nil {
		err = stacktrace.Propagate(err, "could not get node ID")
		panic(err)
	}
	return result
}

// Shutdown ...
func (a *Agent) Shutdown() error {

	a.lock.Lock()
	defer a.lock.Unlock()
	a.logger.Printf("[INFO] agent: gracefully shutting down ...")
	if a.state == Shutdown {
		return nil
	}
	err := a.swarm.Shutdown()
	if err != nil {
		err = stacktrace.Propagate(err, "cannot gracefully shutdown agent")
		return err
	}
	a.state = Shutdown
	close(a.shutdownCh)
	return nil
}

// ShutdownCh ...
func (a *Agent) ShutdownCh() <-chan struct{} {

	return a.shutdownCh
}
