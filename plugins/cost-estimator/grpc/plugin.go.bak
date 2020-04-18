package grpc

import (
	"context"

	model "github.com/da-moon/coe865-final/model"
	shared "github.com/da-moon/coe865-final/plugins/shared"
	plugin "github.com/hashicorp/go-plugin"
	grpcx "google.golang.org/grpc"
)

// GRPCClient is an implementation of shared that talks over gRPC.
type Plugin struct {
	// Plugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl shared.OverlayNetworkInterface
}

// GRPCClient - Required method to implement Plugin interface
func (p *Plugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpcx.ClientConn) (interface{}, error) {
	return &Client{client: model.NewOverlayNetworkClient(c)}, nil
}

// GRPCServer - Required method to implement Plugin interface
func (p *Plugin) GRPCServer(broker *plugin.GRPCBroker, s *grpcx.Server) error {
	model.RegisterOverlayNetworkServer(s, &Server{Impl: p.Impl})
	return nil
}
