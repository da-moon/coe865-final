package netrpc

import (
	rpc "net/rpc"

	shared "github.com/da-moon/coe865-final/plugins/shared"
	plugin "github.com/hashicorp/go-plugin"
)

// OverlayNetwork - this is the interface that we're exposing as a plugin.
// Plugin - This is the implementation of plugin.Plugin so we can serve/consume this.
type Plugin struct {
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl shared.OverlayNetworkInterface
}

// GRPCClient - Required method to implement Plugin interface
func (p *Plugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {

	return &Client{client: c}, nil

}

// Server - Required method to implement Plugin interface

func (p *Plugin) Server(*plugin.MuxBroker) (interface{}, error) {

	return &Server{Impl: p.Impl}, nil
}
