package flags

import (
	"flag"
	"os"
	"path/filepath"
)

// CostEstimatorPathFlag ...
func CostEstimatorPathFlag(f *flag.FlagSet) *string {

	result := os.Getenv("OVERLAY_COST_ESTIMATOR_PLUGIN")
	if result == "" {
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		result = filepath.Join(path, "bin/cost-estimator-plugin")
	}
	return f.String("cost-estimator-path", result,
		"cost estimator plugin path.")
}

// ConfigFilePathFlag ...
func ConfigFilePathFlag(f *flag.FlagSet) *string {
	result := os.Getenv("OVERLAY_CONFIG_FILE")
	return f.String("config-file", result,
		"json config file path.")
}

// LogLevelFlag ...
func LogLevelFlag(f *flag.FlagSet) *string {
	result := os.Getenv("OVERLAY_LOG_LEVEL")
	if result == "" {
		result = "INFO"
	}
	return f.String("log-level", result,
		"flag used to indicate log level")

}

// DevFlag ...
func DevFlag(f *flag.FlagSet) *bool {
	// its false by default
	var result bool
	return f.Bool("dev", result,

		"Enable development mode.")
}

// RPCPortFlag ...
func RPCPortFlag(f *flag.FlagSet) *int {
	result := 1450

	return f.Int("rpc-port", result,
		"overlay daemon rpc port indicator.")
}

// CronFlag ...
func CronFlag(f *flag.FlagSet) *string {
	// cron job that sets
	// how often messages are
	// broadcastes
	result := "@every 10s"
	return f.String("cron", result,
		"controls how often messages are broadcasted.")
}
