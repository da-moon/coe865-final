package flags

import (
	"flag"
	"os"
	"path/filepath"
)

// MasterKeyFlag ...
func MasterKeyFlag(f *flag.FlagSet) *string {
	result := os.Getenv("DARE_MASTER_KEY")
	if result == "" {
		result = "b6c4bba7a385aef779965cb0b7d66316ab091704042606797871"
	}
	return f.String("master-key", result,
		"Master Key used in encryption-decryption process.")
}

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
