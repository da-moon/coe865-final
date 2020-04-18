package daemon

import (
	// "crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	flags "github.com/da-moon/coe865-final/cmd/overlay-network/flags"
	mapstructure "github.com/mitchellh/mapstructure"
)

type dirEnts []os.FileInfo

// Config ...
type Config struct {
	CostEstimatorPath          string             `json:"cost_estimator_path" mapstructure:"cost_estimator_path"`
	LogLevel                   string             `json:"log_level" mapstructure:"log_level"`
	Protocol                   int                `json:"protocol" mapstructure:"protocol"`
	DevelopmentMode            bool               `json:"development_mode" mapstructure:"development_mode"`
	Cron                       string             `json:"cron" mapstructure:"cron"`
	Port                       int                `json:"port" mapstructure:"port"`
	Self                       *RouteController   `json:"self" mapstructure:"self"`
	ConnectedRouteControllers  []RouteController  `json:"connected_route_controllers" mapstructure:"connected_route_controllers"`
	ConnectedAutonomousSystems []AutonomousSystem `json:"connected_autonomous_systems" mapstructure:"connected_autonomous_systems"`
}

// RouteController ...
type RouteController struct {
	ID                     int    `json:"id" mapstructure:"id"`
	AutonomousSystemNumber int    `json:"autonomous_system_number" mapstructure:"autonomous_system_number"`
	IP                     string `json:"ip" mapstructure:"ip"`
}

// RouteController ...
type AutonomousSystem struct {
	Number       int `json:"number" mapstructure:"number"`
	LinkCapacity int `json:"link_capacity" mapstructure:"link_capacity"`
	Cost         int `json:"cost" mapstructure:"cost"`
}

// DefaultConfig ...
func DefaultConfig() *Config {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// bytes := make([]byte, jwtSecretLen)
	// if _, err := rand.Read(bytes); err != nil {
	// 	panic(err)

	// }
	result := &Config{
		LogLevel:        "INFO",
		DevelopmentMode: false,
		// Protocol:          daemon.CoreVersionMax,
		CostEstimatorPath: filepath.Join(path, "cost-estimator"),
		Cron:              "@every 10s",
		Port:              1450,
		Self: &RouteController{
			ID:                     1,
			AutonomousSystemNumber: 100,
			IP:                     "10.2.2.1",
		},
		ConnectedRouteControllers:  make([]RouteController, 0),
		ConnectedAutonomousSystems: make([]AutonomousSystem, 0),
	}

	result.ConnectedRouteControllers = append(
		result.ConnectedRouteControllers,
		RouteController{
			ID:                     2,
			AutonomousSystemNumber: 200,
			IP:                     "10.1.1.2",
		})
	result.ConnectedRouteControllers = append(
		result.ConnectedRouteControllers,
		RouteController{
			ID:                     3,
			AutonomousSystemNumber: 300,
			IP:                     "11.1.1.2",
		},
	)
	result.ConnectedAutonomousSystems = append(
		result.ConnectedAutonomousSystems,
		AutonomousSystem{
			Number:       10,
			LinkCapacity: 2,
			Cost:         5,
		},
	)
	result.ConnectedAutonomousSystems = append(
		result.ConnectedAutonomousSystems,
		AutonomousSystem{
			Number:       200,
			LinkCapacity: 10,
			Cost:         5,
		})
	result.ConnectedAutonomousSystems = append(
		result.ConnectedAutonomousSystems,
		AutonomousSystem{
			Number:       300,
			LinkCapacity: 10,
			Cost:         5,
		})
	return result
}

func (c *Command) readConfig() *Config {
	var cmdConfig Config
	var configFiles []string
	const entrypoint = "daemon"
	cmdFlags := flag.NewFlagSet(entrypoint, flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	cmdFlags.Var((*flags.AppendSliceValue)(&configFiles), "config-file",
		"json file to read config from")
	logLevel := flags.LogLevelFlag(cmdFlags)
	costEstimatorPath := flags.CostEstimatorPathFlag(cmdFlags)
	dev := flags.DevFlag(cmdFlags)
	if err := cmdFlags.Parse(c.args); err != nil {
		return nil
	}
	cmdConfig.DevelopmentMode = *dev
	cmdConfig.LogLevel = *logLevel
	cmdConfig.CostEstimatorPath = *costEstimatorPath
	config := DefaultConfig()
	if len(configFiles) > 0 {
		fileConfig, err := ReadConfigPaths(configFiles)
		if err != nil {
			c.Ui.Error(fmt.Sprintf("[ERROR]: %s", err.Error()))
			return nil
		}
		config = MergeConfig(config, fileConfig)
	}
	config = MergeConfig(config, &cmdConfig)
	return config
}

// DecodeConfig ...
func DecodeConfig(r io.Reader) (*Config, error) {
	var raw interface{}
	dec := json.NewDecoder(r)
	if err := dec.Decode(&raw); err != nil {
		return nil, err
	}

	var md mapstructure.Metadata
	var result Config
	msdec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:    &md,
		Result:      &result,
		ErrorUnused: true,
	})
	if err != nil {
		return nil, err
	}

	if err := msdec.Decode(raw); err != nil {
		return nil, err
	}

	return &result, nil
}

func containsKey(keys []string, key string) bool {
	for _, k := range keys {
		if k == key {
			return true
		}
	}
	return false
}

// MergeConfig ...
func MergeConfig(a, b *Config) *Config {
	result := *a

	if b.CostEstimatorPath != "" {
		result.CostEstimatorPath = b.CostEstimatorPath
	}

	if b.LogLevel != "" {
		result.LogLevel = b.LogLevel
	}
	if b.Protocol > 0 {
		result.Protocol = b.Protocol
	}
	result.DevelopmentMode = b.DevelopmentMode

	return &result
}

// ReadConfigPaths reads the paths in the given order to load configurations.
func ReadConfigPaths(paths []string) (*Config, error) {
	result := new(Config)
	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("Error reading '%s': %s", path, err)
		}

		fi, err := f.Stat()
		if err != nil {
			f.Close()
			return nil, fmt.Errorf("Error reading '%s': %s", path, err)
		}

		if !fi.IsDir() {
			config, err := DecodeConfig(f)
			f.Close()

			if err != nil {
				return nil, fmt.Errorf("Error decoding '%s': %s", path, err)
			}

			result = MergeConfig(result, config)
			continue
		}

		contents, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			return nil, fmt.Errorf("Error reading '%s': %s", path, err)
		}

		sort.Sort(dirEnts(contents))

		for _, fi := range contents {
			if fi.IsDir() {
				continue
			}

			if !strings.HasSuffix(fi.Name(), ".json") {
				continue
			}

			subpath := filepath.Join(path, fi.Name())
			f, err := os.Open(subpath)
			if err != nil {
				return nil, fmt.Errorf("Error reading '%s': %s", subpath, err)
			}

			config, err := DecodeConfig(f)
			f.Close()

			if err != nil {
				return nil, fmt.Errorf("Error decoding '%s': %s", subpath, err)
			}

			result = MergeConfig(result, config)
		}
	}

	return result, nil
}

// Len ...
func (d dirEnts) Len() int {
	return len(d)
}

// Less ...
func (d dirEnts) Less(i, j int) bool {
	return d[i].Name() < d[j].Name()
}

// Swap ...
func (d dirEnts) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
