package daemon

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/gob"
	"errors"
)

func init() {

	gob.Register(Message{})
	gob.Register(HelloPayload{})
	gob.Register(JoinPayload{})
	gob.Register(AgentPayload{})
}

// Message ...
type Message struct {
	Origin       Identity
	PayloadBytes []byte
	Signature    []byte
}

// NewMessage ...
func NewMessage(key *rsa.PrivateKey, payload interface{}) Message {

	var payloadBuf bytes.Buffer
	gob.NewEncoder(&payloadBuf).Encode(&payload)
	payloadBytes := payloadBuf.Bytes()
	hashed := sha256.Sum256(payloadBytes)
	sig, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}
	return Message{
		Origin:       Identity(key.PublicKey),
		PayloadBytes: payloadBytes,
		Signature:    sig,
	}
}

// Verify ...
func (m Message) Verify() error {

	if len(m.Signature) == 0 {
		return errors.New("empty signature")
	}
	hashed := sha256.Sum256(m.PayloadBytes)
	pubKey := rsa.PublicKey(m.Origin)
	return rsa.VerifyPKCS1v15(&pubKey, crypto.SHA256, hashed[:], m.Signature)
}

// Payload ...
func (m Message) Payload() (interface{}, error) {

	var value interface{}
	err := gob.NewDecoder(bytes.NewBuffer(m.PayloadBytes)).Decode(&value)
	return value, err
}

// HelloPayload ...
type HelloPayload struct {
	YourAddr string
}

// JoinPayload ...
type JoinPayload struct {
	Addr string
}

// AgentPayload ...
type AgentPayload struct {
	Sequence sequencer
	Text     string
}
