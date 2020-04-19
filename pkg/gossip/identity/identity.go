package identity

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"

	"github.com/palantir/stacktrace"
)

const defaultSize = 4096

// Identity is used as identitier of a peer.
type Identity struct {
	private *rsa.PrivateKey
}

// Default ...
func Default() (*Identity, error) {
	return New(defaultSize)
}

// New ...
func New(size int) (*Identity, error) {
	private, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		err = stacktrace.Propagate(err, "could not generate a new RSA key with size '%d'", size)
		return nil, err
	}

	result := &Identity{
		private: private,
	}
	return result, nil
}

// Sha256 ...
func (i *Identity) Sha256() ([]byte, error) {
	pubKey := i.private.PublicKey
	derEncoded, err := x509.MarshalPKIXPublicKey(&pubKey)
	if err != nil {
		err = stacktrace.Propagate(err, "could not extract public key")
		return nil, err
	}
	if derEncoded == nil {
		err = stacktrace.NewError("encoded sha256 hash was a nil slice")
		return nil, nil
	}
	if len(derEncoded) == 0 {
		err = stacktrace.NewError("encoded sha256 hash was an empty slice")
		return nil, err
	}
	result := sha256.Sum256(derEncoded)
	return result[:], nil
}

// Sha256String ...
func (i *Identity) Sha256String() (string, error) {
	hash, err := i.Sha256()
	if err != nil {
		err = stacktrace.Propagate(err, "could not calculate sha256 hash of public key")
		return "", err
	}
	if hash == nil {
		err = stacktrace.NewError("calculate sha256 hash was a nil slice")
		return "", err
	}
	if len(hash) == 0 {
		err = stacktrace.NewError("calculate sha256 hash was an empty slice")
		return "", err
	}
	return hex.EncodeToString(hash), nil
}
