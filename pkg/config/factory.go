package config

import (
	"sync"

	"github.com/da-moon/coe865-final/pkg/gossip/swarm"
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
	MinPeers        int    `json:"min_peers" mapstructure:"min_peers"`
	MaxPeers        int    `json:"max_peers" mapstructure:"max_peers"`
}

// DefaultConfigFactory ...
func DefaultConfigFactory() *ConfigFactory {

	result := &ConfigFactory{
		LogLevel:        "INFO",
		DevelopmentMode: false,
		Protocol:        ProtocolVersion,
		Cron:            "@every 3s",
		Port:            DefaultPort,
		MinPeers:        swarm.DefaultMinPeers,
		MaxPeers:        swarm.DefaultMaxPeers,
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
	result.MaxPeers = c.MaxPeers
	result.MinPeers = c.MinPeers
	return result
}

// Init ...
func (c *ConfigFactory) Init() {

	c.Do(func() {
		if len(c.LogLevel) == 0 {
			c.LogLevel = "INFO"
		}
		if c.Protocol == 0 {
			c.Protocol = ProtocolVersion
		}
		if len(c.Cron) == 0 {
			c.Cron = "@every 3s"
		}
		if c.Port == 0 {
			c.Port = DefaultPort
		}
		if c.MinPeers == 0 {
			c.MinPeers = swarm.DefaultMinPeers
		}
		if c.MaxPeers == 0 {
			c.MaxPeers = swarm.DefaultMaxPeers
		}

	})
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
	if b.MinPeers > 0 {
		result.MinPeers = b.MinPeers
	}
	if b.MaxPeers > 0 {
		result.MinPeers = b.MinPeers
	}
	if b.Cron != "" {
		result.Cron = b.Cron
	}
	result.DevelopmentMode = b.DevelopmentMode
	return &result
}
