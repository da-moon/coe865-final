package daemon

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"

	model "github.com/da-moon/coe865-final/model"
	"github.com/da-moon/coe865-final/pkg/config"
	"github.com/da-moon/coe865-final/pkg/utils"
	costEstimatorGrpc "github.com/da-moon/coe865-final/plugins/cost-estimator/grpc"
	costEstimatorRPC "github.com/da-moon/coe865-final/plugins/cost-estimator/net-rpc"
	shared "github.com/da-moon/coe865-final/plugins/shared"
	plugin "github.com/hashicorp/go-plugin"
	stacktrace "github.com/palantir/stacktrace"
	cron "github.com/robfig/cron/v3"
)

// Core ...
type Core struct {
	conf         *config.Config
	lock         *sync.Mutex
	logger       *log.Logger
	shutdown     bool
	shutdownCh   chan struct{}
	shutdownLock sync.Mutex
	cron         *cron.Cron
}

// Create ...
func Create(conf *config.Config, logOutput io.Writer) (*Core, error) {

	if logOutput == nil {
		logOutput = os.Stderr
	}
	logger := log.New(logOutput, "", log.LstdFlags)
	core := &Core{
		conf:       conf,
		logger:     logger,
		shutdownCh: make(chan struct{}),
		cron: cron.New(
			cron.WithLogger(cron.VerbosePrintfLogger(logger)),
		),
	}
	return core, nil
}

// Start ...
func (a *Core) Start() error {
	// @TODO check errors
	a.logger.Printf("overlay network daemon core started!")
	a.cron.AddFunc(a.conf.Cron, a.EstimateCost())
	a.cron.Start()
	return nil
}

// Shutdown ...
func (a *Core) Shutdown() error {
	a.shutdownLock.Lock()
	defer a.shutdownLock.Unlock()
	a.cron.Stop()
	a.shutdown = true
	close(a.shutdownCh)
	a.logger.Println("[INFO] overlay network daemon core: shutdown complete")
	return nil
}

// ShutdownCh ...
func (a *Core) ShutdownCh() <-chan struct{} {

	return a.shutdownCh
}

// EstimateCost ...
func (a *Core) EstimateCost() func() {
	return func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		path := a.conf.CostEstimatorPath
		if len(path) == 0 {
			err := stacktrace.NewError("cost estimator plugin engine binary path is empty")
			a.logger.Println(fmt.Sprintf(("error : %#v"), err.Error()))
			return
		}
		a.logger.Printf("[DEBUG] cost estimator path is %s", path)
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: shared.HandshakeConfig,
			Plugins: map[string]plugin.Plugin{
				"cost_estimator_grpc": &costEstimatorGrpc.Plugin{},
				"cost_estimator":      &costEstimatorRPC.Plugin{},
			},
			Cmd: exec.Command("sh", "-c", path),
			AllowedProtocols: []plugin.Protocol{
				plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
		})
		defer client.Kill()
		_, err := client.Start()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to start the protocol client for EstimateCost engine connection")
			a.logger.Printf(fmt.Sprintf(("error : %#v"), err.Error()))
			return
		}
		// Connect via RPC
		rpcClient, err := client.Client()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to return the protocol client for EstimateCost engine connection")
			a.logger.Printf(fmt.Sprintf(("error : %#v"), err.Error()))
			return
		}
		// Request the plugin
		raw, err := rpcClient.Dispense("cost_estimator_grpc")
		if err != nil {
			err = stacktrace.Propagate(err, "RPC Client could not dispense a new instance of cost_estimator_grpc")
			a.logger.Printf(fmt.Sprintf(("error : %#v"), err.Error()))
			return
		}
		// We should have a overlay network store now! This feels like a normal interface
		// implementation but is in fact over an RPC connection.
		overlay, ok := raw.(shared.OverlayNetworkInterface)
		if !ok {
			err = stacktrace.NewError("failed to convert remote raw client to OverlayNetworkInterface")
			a.logger.Printf(fmt.Sprintf(("error : %#v"), err.Error()))
			return
		}
		// for dst :=range a.conf.ConnectedRouteControllers {
		req := &model.UpdateRequest{
			UUID: utils.UUID(),
		}
		req.SourceRouteController = &model.RouteController{
			ID:                     int32(a.conf.Self.ID),
			AutonomousSystemNumber: int32(a.conf.Self.AutonomousSystemNumber),
			IP:                     a.conf.Self.IP,
		}
		req.DestinationAutonomousSystem = &model.AutonomousSystem{
			Number:       int32(a.conf.ConnectedAutonomousSystems[0].Number),
			LinkCapacity: int32(a.conf.ConnectedAutonomousSystems[0].LinkCapacity),
			Cost:         int32(a.conf.ConnectedAutonomousSystems[0].Cost),
		}
		res, err := overlay.EstimateCost(req)
		if err != nil {
			err = stacktrace.Propagate(err, "cost estimator failed to EstimateCost given input")
			a.logger.Printf(fmt.Sprintf(("error : %#v"), err.Error()))
			return
		}
		a.logger.Printf(fmt.Sprintf(("state : %#v"), res))
		// }
	}
}
