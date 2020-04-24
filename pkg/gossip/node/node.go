package node

import (
	// "github.com/da-moon/coe865-final/pkg/gossip/codec"
	"net"

	// "github.com/da-moon/coe865-final/pkg/gossip/codec"
	"github.com/da-moon/coe865-final/pkg/gossip/key"
	"github.com/palantir/stacktrace"
)

// NodeConfig stores node configuration values
type NodeConfig struct {
	address string
}

// Node wraps a key and codec
// it signs messages and marshalls and
// unmarshalls messages to/from the wire
type Node struct {
	// used for signing and generating ID
	key *key.Key
	// used by the node to accept incoming connections
	listener net.Listener
	// used for logging
	logger *log.Logger
	// used to store node config
	conf *NodeConfig
}

// New returns a new instance of node
// it takes in a conn (net.Conn)
// so that it can create a codec to
// encode and write and read and decode
// to the said connection
func New(logger *log.Logger, conf *NodeConfig) (*Node, error) {
	if logger == nil {
		err := stacktrace.NewError("cannot create a new node since passed logger was nil")
		return nil, err
	}
	if conf == nil {
		err := stacktrace.NewError("cannot create a new node since passed config struct was nil")
		return nil, err
	}
	k, err := key.Default()
	if err != nil {
		err = stacktrace.Propagate(err, "could not create a new node due to an issue with generating RSA key for the node")
		return nil, err
	}
	listener, err := net.Listen("tcp", conf.address)
	if err != nil {
		err = stacktrace.Propagate(err, "could not create a new listener for node with address '%s'", conf.address)
		return nil, err
	}

	result := &Node{
		key:      k,
		conf:     conf,
		logger:   logger,
		listener: listener,
	}
	go result.listen()
	return result, nil
}

// listen spins a new goroutine
// per incomming connection as it is waiting for nodes
// to join .

func (n *Node) listen() (string, error) {
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			n.logger.Printf("[WARN] listener failed to accept new connection : %v", err)
			continue
		}
		n.logger.Printf("[INFO] node '%v' : recieved an incomming connection from peer with address %v", conn.LocalAddr().String())
		// incommingPeer := peerStub{
		// 	codec: codec.NewJSONCodec(conn, conn),
		// }
		// peer := NewPeer(conn)
		// if err := peerManager.AddPeer(peer); err != nil {
		// 	log.Printf("Error adding new peer %s: %s", peer, err)
		// }
	}
}

// ID returns node ID which
// is base64 encoded form of it's public
// key
func (n *Node) ID() (string, error) {
	result, err := n.key.PublicKeyBase64()
	if err != nil {
		err = stacktrace.Propagate(err, "could not get node ID")
		return "", err
	}
	return result, nil
}
