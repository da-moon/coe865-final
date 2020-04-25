package daemon

import (
	"encoding/base64"
	"errors"
	"log"
	"math/rand"
	"net"

	"github.com/da-moon/coe865-final/internal/swarm"
	model "github.com/da-moon/coe865-final/model"
	"github.com/da-moon/coe865-final/pkg/jsonutil"
	prettyjson "github.com/hokaccha/go-prettyjson"
	"github.com/palantir/stacktrace"
)

// Broadcast ...
func (a *Core) Broadcast(text string) {

	a.mu.Lock()
	seq := a.agentSequence
	a.agentSequence++
	a.mu.Unlock()
	a.gossiper.Broadcast(NewMessage(a.key, AgentPayload{
		Sequence: seq,
		Text:     text,
	}))
}
func (a *Core) handleGossip(value interface{}) bool {

	msg, ok := value.(Message)
	if !ok {
		log.Printf("Discarding unexpected gossip value: %#v", value)
		return false
	}
	if err := msg.Verify(); err != nil {
		log.Printf("Discarding message with invalid signature. err: %s; msg: %v", err, msg)
		return false
	}
	payload, err := msg.Payload()
	if err != nil {
		log.Printf("Discarding message with invalid payload: %s", err)
	}
	id := msg.Origin
	fp := id.Fingerprint()
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.identities[fp]; !ok {
		a.identities[fp] = &identityRecord{
			Identity: id,
		}
	}
	record := a.identities[fp]
	switch payload := payload.(type) {
	case HelloPayload:
		return a.handleHello(record, payload, msg)
	case JoinPayload:
		return a.handleJoin(record, payload, msg)
	case AgentPayload:
		return a.handleAgent(record, payload, msg)
	default:
		log.Printf("Discarding message with unknown payload: %#v", payload)
		return false
	}
}
func (a *Core) handleHello(record *identityRecord, payload HelloPayload, msg Message) bool {

	tcpAddr, err := net.ResolveTCPAddr("tcp", payload.YourAddr)
	if err != nil {
		log.Printf("Received hello with invalid TCP address: %s", payload.YourAddr)
		return false
	}
	tcpAddr.Port = a.conf.Port
	go a.gossiper.Broadcast(NewMessage(a.key, JoinPayload{
		Addr: tcpAddr.String(),
	}))
	return false
}
func (a *Core) handleJoin(record *identityRecord, payload JoinPayload, msg Message) bool {

	_, ok := a.joinsByAddr[payload.Addr]
	if !ok {
		a.joinsByAddr[payload.Addr] = msg
		a.logger.Printf("[INFO] Agent %s joined from address '%s'", record.Identity.Fingerprint(), payload.Addr)
	}
	return !ok
}
func (a *Core) handleAgent(record *identityRecord, payload AgentPayload, msg Message) bool {

	isNew := record.AgentSequenceTracker.See(payload.Sequence)
	if isNew {
		decByte, err := base64.StdEncoding.DecodeString(payload.Text)
		if err != nil {
			err := stacktrace.Propagate(err, "agent %s could not decode recieved message from base64", record.Identity.Fingerprint())
			a.logger.Printf("[WARN] %v", err)
			return false
		}
		var update model.UpdateRequest
		err = jsonutil.DecodeJSON(decByte, &update)
		if err != nil {
			err := stacktrace.Propagate(err, "agent %s could not decode recieved message from json", record.Identity.Fingerprint())
			a.logger.Printf("[WARN] %v", err)
			return false
		}
		prettyreq, _ := prettyjson.Marshal(update)
		a.logger.Printf("[INFO] Agent %s : Message %s", record.Identity.Fingerprint(), string(prettyreq))
	}
	return isNew
}
func (a *Core) findPeer(connectedPeers map[swarm.PeerHandle]swarm.Peer) (swarm.Peer, error) {

	var addrs []string
	a.mu.Lock()
	for addr := range a.joinsByAddr {
		addrs = append(addrs, addr)
	}
	a.mu.Unlock()
	if len(addrs) == 0 {
		return nil, swarm.ErrNoPeers
	}
	addr := addrs[rand.Intn(len(addrs))]
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		a.mu.Lock()
		delete(a.joinsByAddr, addr)
		a.mu.Unlock()
		return nil, err
	}
	return NewPeer(conn), nil
}
func (a *Core) onConnect(peer swarm.Peer) error {

	agentPeer, ok := peer.(*AgentPeer)
	if !ok {
		return errors.New("not a AgentPeer")
	}
	var joins []Message
	a.mu.Lock()
	for _, join := range a.joinsByAddr {
		joins = append(joins, join)
	}
	a.mu.Unlock()
	hello := NewMessage(a.key, HelloPayload{
		YourAddr: agentPeer.conn.RemoteAddr().String(),
	})
	if err := peer.Write(hello); err != nil {
		return err
	}
	for _, join := range joins {
		if err := peer.Write(join); err != nil {
			return err
		}
	}
	return nil
}
