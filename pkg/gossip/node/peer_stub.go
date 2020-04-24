package node

import (
	"github.com/da-moon/coe865-final/pkg/gossip/codec"
)

// peer stub is used by
// node to encode/decode messages
// to peers 
type peerStub struc {
	codec codec.Codec
}