package daemon

import (
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
)

// Identity ...
type Identity rsa.PublicKey

// Fingerprint ...
func (id Identity) Fingerprint() fingerprint {

	pubKey := rsa.PublicKey(id)
	derEncoded, err := x509.MarshalPKIXPublicKey(&pubKey)
	if err != nil {
		panic(err)
	}
	return fingerprint(sha256.Sum256(derEncoded))
}

type fingerprint [sha256.Size]byte

// String ...
func (fp fingerprint) String() string {

	return hex.EncodeToString(fp[:8])
}
