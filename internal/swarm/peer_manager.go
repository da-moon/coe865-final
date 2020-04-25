package swarm

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// PeerManagerConfig ...
type PeerManagerConfig struct {
	NewPeer            func(currentPeers map[PeerHandle]Peer) (Peer, error)
	OnConnect          func(peer Peer) error
	MinPeers, MaxPeers int
	RetryDelay         time.Duration
}

// NewPeerManager ...
func NewPeerManager(gossiper Gossiper, config PeerManagerConfig) PeerManager {

	pm := &peerManager{
		gossiper:     gossiper,
		newPeer:      config.NewPeer,
		onConnect:    config.OnConnect,
		minPeers:     config.MinPeers,
		maxPeers:     config.MaxPeers,
		retryDelay:   config.RetryDelay,
		currentPeers: make(map[PeerHandle]Peer),
		notifyChan:   make(chan bool, 1),
		closed:       make(chan bool, 1),
	}
	if pm.retryDelay == 0 {
		pm.retryDelay = DefaultRetryDelay
	}
	pm.gossiperWatchHandle = gossiper.AddPeerWatcher(pm)
	go pm.watchPeers()
	pm.notify()
	return pm
}

type peerManager struct {
	gossiper            Gossiper
	gossiperWatchHandle PeerWatcherHandle
	newPeer             func(currentPeers map[PeerHandle]Peer) (Peer, error)
	onConnect           func(peer Peer) error
	minPeers, maxPeers  int
	retryDelay          time.Duration
	mu                  sync.Mutex
	currentPeers        map[PeerHandle]Peer
	notifyChan          chan bool
	closing             bool
	closed              chan bool
}

// AddPeer ...
func (pm *peerManager) AddPeer(peer Peer) error {

	if _, err := pm.gossiper.AddPeer(peer); err != nil {
		peer.Close()
		return err
	}
	if pm.onConnect != nil {
		if err := pm.onConnect(peer); err != nil {
			peer.Close()
			return err
		}
	}
	return nil
}

// Close ...
func (pm *peerManager) Close() {

	pm.mu.Lock()
	if !pm.closing {
		close(pm.notifyChan)
		pm.gossiper.RemovePeerWatcher(pm.gossiperWatchHandle)
		pm.closing = true
	}
	pm.mu.Unlock()
	<-pm.closed
}
func (pm *peerManager) notify() {

	pm.mu.Lock()
	defer pm.mu.Unlock()
	if pm.closing {
		return
	}
	select {
	case pm.notifyChan <- true:
	default:
	}
}
func (pm *peerManager) notifyAfter(d time.Duration) {

	go func() {
		time.Sleep(d)
		pm.notify()
	}()
}
func (pm *peerManager) watchPeers() {

	for _ = range pm.notifyChan {
		currentPeers := pm.currentPeersCopy()
		n := len(currentPeers)
		if n < pm.minPeers {
			pm.findPeerToAdd(currentPeers)
		} else if pm.maxPeers != 0 && n > pm.maxPeers {
			pm.findPeerToRemove(currentPeers)
		}
	}
	pm.closed <- true
}
func (pm *peerManager) findPeerToAdd(currentPeers map[PeerHandle]Peer) {

	if pm.newPeer == nil {
		return
	}
	peer, err := pm.newPeer(currentPeers)
	if err != nil {
		if err != ErrNoPeers {
			log.Printf("Error finding a new peer to add. Retrying in %s. err: %s", pm.retryDelay, err)
		}
		pm.notifyAfter(pm.retryDelay)
		return
	}
	if err := pm.AddPeer(peer); err != nil {
		log.Printf("Error adding new peer %s. Retrying in %s. err: %s", peer, pm.retryDelay, err)
		pm.notifyAfter(pm.retryDelay)
		return
	}
}
func (pm *peerManager) findPeerToRemove(currentPeers map[PeerHandle]Peer) {

	var victimHandle PeerHandle
	var victim Peer
	var ok bool
	var i int
	for handle, peer := range currentPeers {
		i++
		if rand.Intn(i) == 0 {
			victimHandle = handle
			victim = peer
			ok = true
		}
	}
	if ok {
		log.Printf("Too many connected peers; removing %s.", victim)
		pm.gossiper.RemovePeer(victimHandle)
	}
}
func (pm *peerManager) currentPeersCopy() map[PeerHandle]Peer {

	pm.mu.Lock()
	defer pm.mu.Unlock()
	currentPeers := make(map[PeerHandle]Peer, len(pm.currentPeers))
	for handle, peer := range pm.currentPeers {
		currentPeers[handle] = peer
	}
	return currentPeers
}

// PeerAdded ...
func (pm *peerManager) PeerAdded(handle PeerHandle, peer Peer) {

	pm.mu.Lock()
	pm.currentPeers[handle] = peer
	pm.mu.Unlock()
	pm.notify()
}

// PeerRemoved ...
func (pm *peerManager) PeerRemoved(handle PeerHandle, peer Peer) {

	pm.mu.Lock()
	delete(pm.currentPeers, handle)
	pm.mu.Unlock()
	pm.notify()
}
