package config

import (
	"sync"
)

// ProtocolVersion ...
const ProtocolVersion = 1

// ConfigFactory is used to add extra metadata
// the server needs to a config struct
// parsed from raw config file
type ConfigFactory struct {
	sync.Once
	DevelopmentMode bool   `json:"development_mode" mapstructure:"development_mode"`
	Protocol        int    `json:"protocol" mapstructure:"protocol"`
	Port            int    `json:"port" mapstructure:"port"`
	Cron            string `json:"cron" mapstructure:"cron"`
	LogLevel        string `json:"log_level" mapstructure:"log_level"`
}

// DefaultConfigFactory ...
func DefaultConfigFactory() *ConfigFactory {

	result := &ConfigFactory{
		LogLevel:        "INFO",
		DevelopmentMode: false,
		Protocol:        ProtocolVersion,
		Cron:            "@every 3s",
		Port:            DefaultPort,
	}
	return result
}

// New returns a new config struct
func (c *ConfigFactory) New(self *RouteController, connectedRouteControllers []RouteController, connectedAutonomousSystems []AutonomousSystem) *Config {

	result := &Config{
		Self:                       self,
		ConnectedRouteControllers:  connectedRouteControllers,
		ConnectedAutonomousSystems: connectedAutonomousSystems,
	}
	result.DevelopmentMode = c.DevelopmentMode
	result.Protocol = c.Protocol
	result.Port = c.Port
	result.Cron = c.Cron
	result.LogLevel = c.LogLevel
	return result
}

// MergeFactory ...
func MergeFactory(a, b *ConfigFactory) *ConfigFactory {
	result := *a
	if b.LogLevel != "" {
		result.LogLevel = b.LogLevel
	}
	if b.Protocol > 0 {
		result.Protocol = b.Protocol
	}
	if b.Port > 0 {
		result.Port = b.Port
	}
	if b.Cron != "" {
		result.Cron = b.Cron
	}
	result.DevelopmentMode = b.DevelopmentMode
	return &result
}
