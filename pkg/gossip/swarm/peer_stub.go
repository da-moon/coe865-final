package swarm

import (
	"github.com/da-moon/coe865-final/pkg/gossip/codec"
)

// Peer is used by
// swarm to encode/decode messages
// to peers
type Peer struct {
	codec codec.Codec
}
