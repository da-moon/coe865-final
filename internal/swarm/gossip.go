package swarm

import (
	"io"
	"log"
	"sync"
)

type incomingMessage struct {
	message    interface{}
	peerHandle PeerHandle
}
type outgoingMessage struct {
	message      interface{}
	excludePeers map[PeerHandle]bool
}
type gossiper struct {
	incomingMessages      chan incomingMessage
	outgoingMessages      chan outgoingMessage
	mu                    sync.Mutex
	nextPeerWatcherHandle PeerWatcherHandle
	peerWatchers          map[PeerWatcherHandle]PeerWatcher
	nextPeerHandle        PeerHandle
	peers                 map[PeerHandle]Peer
	closing               bool
	peerClosedChans       map[PeerHandle]chan<- bool
	closed                chan bool
}

// NewGossiper ...
func NewGossiper(updateFunc func(interface{}) bool) Gossiper {

	g := &gossiper{
		incomingMessages:      make(chan incomingMessage, 1000),
		outgoingMessages:      make(chan outgoingMessage, 1000),
		nextPeerWatcherHandle: 1,
		peerWatchers:          make(map[PeerWatcherHandle]PeerWatcher),
		nextPeerHandle:        peerHandleStart,
		peers:                 make(map[PeerHandle]Peer),
		peerClosedChans:       make(map[PeerHandle]chan<- bool),
		closed:                make(chan bool),
	}
	go g.pumpIncoming(updateFunc)
	go g.pumpOutgoing()
	return g
}

// AddPeer ...
func (g *gossiper) AddPeer(peer Peer) (handle PeerHandle, err error) {

	g.mu.Lock()
	defer g.mu.Unlock()
	if g.closing {
		err = ErrClosed
		return
	}
	handle = g.nextPeerHandle
	g.nextPeerHandle++
	g.peers[handle] = peer
	for _, peerWatcher := range g.peerWatchers {
		peerWatcher.PeerAdded(handle, peer)
	}
	go g.pumpPeerIncoming(handle, peer)
	return
}
func (g *gossiper) pumpOutgoing() {

	for outgoingMessage := range g.outgoingMessages {
		g.mu.Lock()
		peers := make(map[PeerHandle]Peer)
		for handle, peer := range g.peers {
			peers[handle] = peer
		}
		g.mu.Unlock()
		for handle, peer := range peers {
			if outgoingMessage.excludePeers[handle] {
				continue
			}
			err := peer.Write(outgoingMessage.message)
			if err != nil {
				if err != io.EOF {
					log.Printf("error writing message to peer %s: %s", peer, err)
				}
				g.RemovePeer(handle)
			}
		}
	}
	close(g.closed)
}
func (g *gossiper) pumpIncoming(updateFunc func(interface{}) bool) {

	for incomingMessage := range g.incomingMessages {
		if updateFunc(incomingMessage.message) || incomingMessage.peerHandle == selfHandle {
			g.outgoingMessages <- outgoingMessage{
				message:      incomingMessage.message,
				excludePeers: map[PeerHandle]bool{incomingMessage.peerHandle: true},
			}
		}
	}
	close(g.outgoingMessages)
}
func (g *gossiper) pumpPeerSingleIncoming(msg incomingMessage) bool {

	g.mu.Lock()
	defer g.mu.Unlock()
	if g.closing {
		return false
	}
	g.incomingMessages <- msg
	return true
}
func (g *gossiper) pumpPeerIncoming(handle PeerHandle, peer Peer) {

	for {
		message, err := peer.Read()
		if err != nil {
			if err != io.EOF {
				log.Printf("error reading from peer %s; disconnecting: %s", peer, err)
			}
			break
		}
		msg := incomingMessage{message, handle}
		if !g.pumpPeerSingleIncoming(msg) {
			break
		}
	}
	g.RemovePeer(handle)
}

// RemovePeer ...
func (g *gossiper) RemovePeer(handle PeerHandle) {

	g.mu.Lock()
	defer g.mu.Unlock()
	if peer, ok := g.peers[handle]; ok {
		if err := peer.Close(); err != nil {
			log.Printf("error closing peer %s: %s", peer, err)
		}
		delete(g.peers, handle)
		for _, peerWatcher := range g.peerWatchers {
			peerWatcher.PeerRemoved(handle, peer)
		}
		if c, ok := g.peerClosedChans[handle]; ok {
			c <- true
			delete(g.peerClosedChans, handle)
		}
	}
}

// Broadcast ...
func (g *gossiper) Broadcast(message interface{}) {

	g.incomingMessages <- incomingMessage{
		message:    message,
		peerHandle: selfHandle,
	}
}

// AddPeerWatcher ...
func (g *gossiper) AddPeerWatcher(peerWatcher PeerWatcher) PeerWatcherHandle {

	g.mu.Lock()
	defer g.mu.Unlock()
	handle := g.nextPeerWatcherHandle
	g.nextPeerWatcherHandle++
	g.peerWatchers[handle] = peerWatcher
	for handle, peer := range g.peers {
		peerWatcher.PeerAdded(handle, peer)
	}
	return handle
}

// RemovePeerWatcher ...
func (g *gossiper) RemovePeerWatcher(handle PeerWatcherHandle) {

	g.mu.Lock()
	defer g.mu.Unlock()
	delete(g.peerWatchers, handle)
}

// Close ...
func (g *gossiper) Close() {

	g.mu.Lock()
	g.closing = true
	nPeers := len(g.peers)
	c := make(chan bool)
	for handle, peer := range g.peers {
		g.peerClosedChans[handle] = c
		if err := peer.Close(); err != nil {
			log.Printf("error closing peer %s: %s", peer, err)
		}
	}
	g.mu.Unlock()
	for i := 0; i < nPeers; i++ {
		<-c
	}
	close(g.incomingMessages)
	<-g.closed
}
