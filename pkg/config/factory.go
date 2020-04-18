package config

import (
	"os"
	"path/filepath"
)

// ProtocolVersion ...
const ProtocolVersion = 1

// ConfigFactory is used to add extra metadata
// the server needs to a config struct
// parsed from raw config file
type ConfigFactory struct {
	DevelopmentMode   bool   `json:"development_mode" mapstructure:"development_mode"`
	Protocol          int    `json:"protocol" mapstructure:"protocol"`
	Port              int    `json:"port" mapstructure:"port"`
	Cron              string `json:"cron" mapstructure:"cron"`
	CostEstimatorPath string `json:"cost_estimator_path" mapstructure:"cost_estimator_path"`
	LogLevel          string `json:"log_level" mapstructure:"log_level"`
}

// DefaultConfigFactory
// returns a config factory with default values
func DefaultConfigFactory() *ConfigFactory {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	result := &ConfigFactory{
		LogLevel:          "INFO",
		DevelopmentMode:   false,
		Protocol:          ProtocolVersion,
		CostEstimatorPath: filepath.Join(path, "cost-estimator"),
		Cron:              "@every 10s",
		Port:              1450,
	}
	return result
}

// New returns a new config struct
func (c *ConfigFactory) New(
	self *RouteController,
	connectedRouteControllers []RouteController,
	connectedAutonomousSystems []AutonomousSystem,
) *Config {
	result := &Config{
		Self:                       self,
		ConnectedRouteControllers:  connectedRouteControllers,
		ConnectedAutonomousSystems: connectedAutonomousSystems,
	}
	result.DevelopmentMode = c.DevelopmentMode
	result.Protocol = c.Protocol
	result.Port = c.Port
	result.Cron = c.Cron
	result.CostEstimatorPath = c.CostEstimatorPath
	result.LogLevel = c.LogLevel
	return result
}

// MergeFactory ...
func MergeFactory(a, b *ConfigFactory) *ConfigFactory {
	// fmt.Println("MergeFactory")
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
	if b.Port > 0 {
		// fmt.Println("MergeFactory Port Override", "original", result.Port, "new", b.Port)
		result.Port = b.Port
	}
	if b.Cron != "" {
		result.Cron = b.Cron
	}
	result.DevelopmentMode = b.DevelopmentMode

	return &result
}
