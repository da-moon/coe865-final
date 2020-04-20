package node

import (
	"encoding/json"
	"net"
)

// Node ...
type Node interface {
	Read() (interface{}, error)
	Write(interface{}) error
	Close() error
}

// Node implements Peer, sending each message gob-encoded over the wire.
type node struct {
	conn net.Conn
	enc  *json.Encoder
	dec  *json.Decoder
}

// New returns a struct compliant with Node insterface
// in which encapsulates a net.listener (eg. tcp socket)
// and encodes and decodes messages based on given encoding
// type
func New(conn net.Conn) Node {
	enc := json.NewEncoder(conn)
	dec := json.NewDecoder(conn)
	return &node{
		conn: conn,
		enc:  enc,
		dec:  dec,
	}
}

// String returns a string representation of this peer.
func (p *node) String() string {
	return p.conn.RemoteAddr().String()
}

// Write sends a message to the remote peer.
func (p *node) Write(msg interface{}) error {
	return p.enc.Encode(&msg)
}

// Read receives a message from the remote peer.
func (p *node) Read() (interface{}, error) {
	var val interface{}
	err := p.dec.Decode(&val)
	return val, err
}

// Close closes the connection to the remote peer.
func (p *node) Close() error {
	return p.conn.Close()
}
