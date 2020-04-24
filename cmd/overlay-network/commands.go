package main
import (
	"os"
	"os/signal"
	command "github.com/da-moon/coe865-final/cmd/overlay-network/command"
	daemon "github.com/da-moon/coe865-final/cmd/overlay-network/command/daemon"
	cli "github.com/mitchellh/cli"
)
// Commands is the mapping of all the available Serf commands.
var Commands map[string]cli.CommandFactory
func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}
	Commands = map[string]cli.CommandFactory{
		"daemon": func() (cli.Command, error) {
			return &daemon.Command{
				Ui:         ui,
				ShutdownCh: make(chan struct{}),
			}, nil
		},
		"keygen": func() (cli.Command, error) {
			return &command.KeygenCommand{
				Ui: ui,
			}, nil
		},
		"parse-config": func() (cli.Command, error) {
			return &command.ParseConfigCommand{
				Ui: ui,
			}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Name:              EntryPointName,
				Revision:          GitCommit,
				Version:           Version,
				VersionPrerelease: VersionPrerelease,
				Ui:                ui,
			}, nil
		},
	}
}
// makeShutdownCh returns a channel that can be used for shutdown
// notifications for commands. This channel will send a message for every
// interrupt received.
func makeShutdownCh() <-chan struct{} {
	resultCh := make(chan struct{})
	signalCh := make(chan os.Signal, 4)
	signal.Notify(signalCh, os.Interrupt)
	go func() {
		for {
			<-signalCh
			resultCh <- struct{}{}
		}
	}()
	return resultCh
}
