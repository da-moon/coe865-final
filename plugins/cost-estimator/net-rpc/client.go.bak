package netrpc

import (
	rpc "net/rpc"

	model "github.com/da-moon/coe865-final/model"
	stacktrace "github.com/palantir/stacktrace"
)

// Client is an implementation of shared that talks over RPC.
type Client struct{ client *rpc.Client }

// EstimateCost ...
func (c *Client) EstimateCost(req *model.UpdateRequest) (*model.UpdateResponse, error) {
	var _resp model.UpdateResponse
	err := c.client.Call("Plugin.EstimateCost", req, &_resp)
	if err != nil {
		err = stacktrace.Propagate(err, "EstimateCost call failed with request %#v", req)
	}
	return &_resp, err
}

// KeyExchange ...
func (c *Client) KeyExchange(req *model.KeyExchangeRequest) (*model.KeyExchangeResponse, error) {
	var _resp model.KeyExchangeResponse
	err := c.client.Call("Plugin.KeyExchange", req, &_resp)
	if err != nil {
		err = stacktrace.Propagate(err, "KeyExchange call failed with request %#v", req)
	}
	return &_resp, err
}
