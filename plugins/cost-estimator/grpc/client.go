package grpc

import (
	"context"

	model "github.com/da-moon/coe865-final/model"
	stacktrace "github.com/palantir/stacktrace"
)

// Client is an implementation of shared.OverlayNetwork that talks over gRPC.
type Client struct {
	client model.OverlayNetworkClient
}

// EstimateCost ...
func (c *Client) EstimateCost(req *model.UpdateRequest) (*model.UpdateResponse, error) {

	_resp, err := c.client.EstimateCost(context.Background(), req)

	if err != nil {
		err = stacktrace.Propagate(err, "EstimateCost call failed with request %#v", req)
	}
	return _resp, err
}

// EstimateCost ...

func (c *Client) KeyExchange(req *model.KeyExchangeRequest) (*model.KeyExchangeResponse, error) {

	_resp, err := c.client.KeyExchange(context.Background(), req)
	if err != nil {
		err = stacktrace.Propagate(err, "KeyExchange call failed with request %#v", req)
	}
	return _resp, err
}
