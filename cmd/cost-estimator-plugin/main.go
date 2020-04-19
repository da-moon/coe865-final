package main

import (
	handler "github.com/da-moon/coe865-final/cmd/cost-estimator-plugin/handler"
	grpc "github.com/da-moon/coe865-final/plugins/cost-estimator/grpc"
	shared "github.com/da-moon/coe865-final/plugins/shared"
	plugin "github.com/hashicorp/go-plugin"
)

// ServeConfig - This is the plugin config thet is used in main function of engine
func main() {

	plugin.Serve(&plugin.ServeConfig{

		HandshakeConfig: shared.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"cost_estimator_grpc": &grpc.Plugin{Impl: &handler.OverlayNetwork{}},
		},
		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
