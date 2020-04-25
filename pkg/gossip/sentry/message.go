package sentry

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"errors"

	"github.com/da-moon/coe865-final/pkg/jsonutil"
	"github.com/palantir/stacktrace"
)

// SignedMessage ...
type SignedMessage struct {
	Origin       rsa.PublicKey `json:"-" mapstructure:"origin"`
	PayloadBytes []byte        `json:"payload_bytes" mapstructure:"payload_bytes"`
	Signature    []byte        `json:"signature" mapstructure:"signature"`
}

// NewMessage ...
func (s *Sentry) NewMessage(payload interface{}) (*SignedMessage, error) {
	enc, err := jsonutil.EncodeJSON(payload)
	if err != nil {
		err = stacktrace.Propagate(err, "sentry could not encode message")
		return nil, err
	}
	hashed := sha256.Sum256(enc)
	sig, err := rsa.SignPKCS1v15(rand.Reader, s.private, crypto.SHA256, hashed[:])
	if err != nil {
		err = stacktrace.Propagate(err, "sentry could not sign the message")
		return nil, err
	}
	result := &SignedMessage{
		Origin:       s.private.PublicKey,
		PayloadBytes: enc,
		Signature:    sig,
	}
	return result, nil
}

// Verify ...
func (s *SignedMessage) Verify() error {
	if len(s.Signature) == 0 {
		return errors.New("empty signature")
	}
	hashed := sha256.Sum256(s.PayloadBytes)
	pubKey := rsa.PublicKey(s.Origin)
	return rsa.VerifyPKCS1v15(&pubKey, crypto.SHA256, hashed[:], s.Signature)
}

// Payload ...
func (s *SignedMessage) Payload() (interface{}, error) {
	var out interface{}
	err := jsonutil.DecodeJSON(s.PayloadBytes, out)
	if err != nil {
		err = stacktrace.Propagate(err, "signed message payload could not be decoded")
		return nil, err
	}
	return out, nil
}
