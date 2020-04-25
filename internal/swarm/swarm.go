package swarm

import (
	"errors"
	"time"
)

var (
	// ErrNoPeers ...
	ErrNoPeers = errors.New("no peers known")
	// ErrClosed ...
	ErrClosed = errors.New("shutting down")
)

// DefaultRetryDelay ...
const (
	// DefaultRetryDelay ...
	DefaultRetryDelay = 5 * time.Second
)

// PeerManager ...
type PeerManager interface {
	AddPeer(peer Peer) error
	Close()
}

// Gossiper ...
type Gossiper interface {
	Broadcast(message interface{})
	AddPeer(peer Peer) (PeerHandle, error)
	RemovePeer(handle PeerHandle)
	AddPeerWatcher(PeerWatcher) PeerWatcherHandle
	RemovePeerWatcher(PeerWatcherHandle)
	Close()
}

// Peer ...
type Peer interface {
	Read() (interface{}, error)
	Write(interface{}) error
	Close() error
}

// PeerWatcher ...
type PeerWatcher interface {
	PeerAdded(handle PeerHandle, peer Peer)
	PeerRemoved(handle PeerHandle, peer Peer)
}

// PeerHandle ...
type PeerHandle uint

const (
	selfHandle      PeerHandle = 0
	peerHandleStart PeerHandle = 1
)

// PeerWatcherHandle ...
type PeerWatcherHandle uint
