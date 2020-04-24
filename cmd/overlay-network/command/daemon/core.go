package daemon

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	model "github.com/da-moon/coe865-final/model"
	config "github.com/da-moon/coe865-final/pkg/config"
	utils "github.com/da-moon/coe865-final/pkg/utils"
	prettyjson "github.com/hokaccha/go-prettyjson"
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
			cron.WithLogger(cron.PrintfLogger(logger)),
		),
	}
	core.logger.SetPrefix("[Core]")
	return core, nil
}

// Start ...
func (a *Core) Start() error {
	a.logger.Printf("[INFO] overlay network daemon core started!")
	entryID, err := a.cron.AddFunc(a.conf.Cron, a.EstimateCost())
	if err != nil {
		err = stacktrace.Propagate(err, "could not start cron job handler")
		a.logger.Printf(fmt.Sprintf(("[WARN] error : %#v"), err.Error()))
		return err
	}
	a.logger.Printf("[INFO] cron job entry ID %v", entryID)
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
	a.logger.Println("[INFO]", "overlay network daemon core: shutdown complete")
	return nil
}

// ShutdownCh ...
func (a *Core) ShutdownCh() <-chan struct{} {
	return a.shutdownCh
}

// EstimateCost ...
func (a *Core) EstimateCost() func() {
	return func() {
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
		prettyreq, _ := prettyjson.Marshal(req)
		a.logger.Println("[INFO]", "req", string(prettyreq))
	}
}
