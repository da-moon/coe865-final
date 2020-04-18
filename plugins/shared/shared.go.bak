package shared

import (
	model "github.com/da-moon/coe865-final/model"
	plugin "github.com/hashicorp/go-plugin"
)

// OverlayNetworkInterface - this is the interface that we're exposing as a plugin.
type OverlayNetworkInterface interface {
	EstimateCost(req *model.UpdateRequest) (*model.UpdateResponse, error)
	KeyExchange(req *model.KeyExchangeRequest) (*model.KeyExchangeResponse, error)
}

// HandshakeConfig - engine-interface handshake configuration
var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  2,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}
