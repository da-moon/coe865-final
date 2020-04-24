package daemon
import (
	"github.com/da-moon/coe865-final/pkg/config"
	"github.com/palantir/stacktrace"
	// "crypto/rand"
	"flag"
	"fmt"
	flags "github.com/da-moon/coe865-final/cmd/overlay-network/flags"
)
func (c *Command) readConfig() *config.Config {
	const entrypoint = "daemon"
	cmdFlags := flag.NewFlagSet(entrypoint, flag.ContinueOnError)
	cmdConfigFactory := config.DefaultConfigFactory()
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	configFile := flags.ConfigFilePathFlag(cmdFlags)
	dev := flags.DevFlag(cmdFlags)
	logLevel := flags.LogLevelFlag(cmdFlags)
	port := flags.RPCPortFlag(cmdFlags)
	costEstimatorPath := flags.CostEstimatorPathFlag(cmdFlags)
	cron := flags.CronFlag(cmdFlags)
	err := cmdFlags.Parse(c.args)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("[ERROR] could not parse arguments : %v", err))
		return nil
	}
	if configFile == nil {
		err = stacktrace.NewError("config file was not provided")
		c.Ui.Error(fmt.Sprintf("[ERROR]: %s", err.Error()))
		return nil
	}
	cmdConfigFactory.DevelopmentMode = *dev
	cmdConfigFactory.LogLevel = *logLevel
	cmdConfigFactory.Port = *port
	cmdConfigFactory.CostEstimatorPath = *costEstimatorPath
	cmdConfigFactory.Cron = *cron
	factory := config.DefaultConfigFactory()
	factory = config.MergeFactory(factory, cmdConfigFactory)
	mapping, err := factory.ReadConfigPaths([]string{*configFile}, config.JSON)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("[ERROR]: %s", err.Error()))
		return nil
	}
	result, ok := mapping[*configFile]
	if !ok {
		err := stacktrace.NewError("could not extract config file from map with key '%s'", *configFile)
		c.Ui.Error(fmt.Sprintf("[ERROR]: %s", err.Error()))
		return nil
	}
	// fmt.Println("readConfig()  result", result)
	return &result
}
