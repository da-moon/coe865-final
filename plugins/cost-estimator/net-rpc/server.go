package netrpc

import (
	model "github.com/da-moon/coe865-final/model"
	shared "github.com/da-moon/coe865-final/plugins/shared"
	stacktrace "github.com/palantir/stacktrace"
)

// Server - This is the RPC server that Client talks to, conforming to the requirements of net/rpc
type Server struct {
	Impl shared.OverlayNetworkInterface
}

// EstimateCost ...
func (s *Server) EstimateCost(_req *model.UpdateRequest, _resp *model.UpdateResponse) error {
	_resp, err := s.Impl.EstimateCost(_req)
	if err != nil {
		err = stacktrace.Propagate(err, "EstimateCost call failed with request %#v", _req)
		return err
	}
	return nil
}

// KeyExchange ...
func (s *Server) KeyExchange(_req *model.KeyExchangeRequest, _resp *model.KeyExchangeResponse) error {
	_resp, err := s.Impl.KeyExchange(_req)
	if err != nil {
		err = stacktrace.Propagate(err, "KeyExchange call failed with request %#v", _req)
		return err
	}
	return nil
}
