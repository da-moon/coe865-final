package grpc

import (
	"context"

	model "github.com/da-moon/coe865-final/model"
	shared "github.com/da-moon/coe865-final/plugins/shared"
	stacktrace "github.com/palantir/stacktrace"
)

// Server - Here is the gRPC server that Client talks to.
type Server struct {
	Impl shared.OverlayNetworkInterface
}

// EstimateCost ...
func (s *Server) EstimateCost(ctx context.Context, _req *model.UpdateRequest) (*model.UpdateResponse, error) {
	resp, err := s.Impl.EstimateCost(_req)
	if err != nil {
		err = stacktrace.Propagate(err, "EstimateCost call failed with request %#v", _req)
	}
	return resp, nil
}

// EstimateCost ...
func (s *Server) KeyExchange(ctx context.Context, _req *model.KeyExchangeRequest) (*model.KeyExchangeResponse, error) {
	resp, err := s.Impl.KeyExchange(_req)
	if err != nil {
		err = stacktrace.Propagate(err, "KeyExchange call failed with request %#v", _req)
	}
	return resp, nil
}
