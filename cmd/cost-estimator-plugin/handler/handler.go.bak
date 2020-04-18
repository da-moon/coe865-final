package handler

import (
	"net"
	"strconv"
	"time"

	model "github.com/da-moon/coe865-final/model"
	"github.com/palantir/stacktrace"
)

// OverlayNetwork - this is the struct that implements plugin's actual operations
type OverlayNetwork struct{}

// EstimateCost - Implementation of EstimateCost method for plugin
func (OverlayNetwork) EstimateCost(req *model.UpdateRequest) (*model.UpdateResponse, error) {
	result := &model.UpdateResponse{}
	result.DestinationAutonomousSystem = new(model.AutonomousSystem)
	result.DestinationAutonomousSystem.Cost = 123
	return result, nil
}

// KeyExchange - Implementation of KeyExchange method for plugin
func (OverlayNetwork) KeyExchange(req *model.KeyExchangeRequest) (*model.KeyExchangeResponse, error) {
	result := &model.KeyExchangeResponse{}
	result.IsOk = true
	return result, nil
}
func getAPIListener(addr string) (net.Listener, error) {
	x, _ := tcpAddressFromString(addr)
	l, err := net.Listen("tcp", x.String())
	if err != nil {
		l.Close()
		err = stacktrace.Propagate(err, "failed to start the daemon api listener for address: %s", addr)
		return nil, err
	}
	apiListener := &tcpKeepAliveListener{l.(*net.TCPListener)}
	return apiListener, nil
}

// tcpAddress -
func tcpAddress(ip string, port int) *net.TCPAddr {
	result := &net.TCPAddr{IP: net.ParseIP(ip), Port: port}

	return result
}

// tcpAddressFromString -
func tcpAddressFromString(addr string) (*net.TCPAddr, error) {
	h, p, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Err] could not generate tcp address from string")
	}

	port, err := strconv.Atoi(p)
	if err != nil {
		return nil, stacktrace.Propagate(err, "[Err] could not generate tcp address from string due to invalid port")

	}

	return tcpAddress(h, port), nil
}

// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used so dead TCP connections eventually go away.
type tcpKeepAliveListener struct {
	*net.TCPListener
}

// Accept ...
func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(30 * time.Second)
	return tc, nil
}
