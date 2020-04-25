package daemon

import (
	"encoding/gob"
	"net"
)

// AgentPeer ...
type AgentPeer struct {
	conn net.Conn
	enc  *gob.Encoder
	dec  *gob.Decoder
}

// NewPeer ...
func NewPeer(conn net.Conn) *AgentPeer {

	return &AgentPeer{
		conn: conn,
		enc:  gob.NewEncoder(conn),
		dec:  gob.NewDecoder(conn),
	}
}

// String ...
func (p *AgentPeer) String() string {

	return p.conn.RemoteAddr().String()
}

// Write ...
func (p *AgentPeer) Write(msg interface{}) error {

	return p.enc.Encode(&msg)
}

// Read ...
func (p *AgentPeer) Read() (interface{}, error) {

	var val interface{}
	err := p.dec.Decode(&val)
	return val, err
}

// Close ...
func (p *AgentPeer) Close() error {

	return p.conn.Close()
}
