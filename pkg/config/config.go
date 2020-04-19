package config

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/da-moon/coe865-final/pkg/jsonutil"
	"github.com/palantir/stacktrace"
)

// Config ...
type Config struct {
	DevelopmentMode            bool               `json:"development_mode" mapstructure:"development_mode"`
	Protocol                   int                `json:"protocol" mapstructure:"protocol"`
	Port                       int                `json:"port" mapstructure:"port"`
	Cron                       string             `json:"cron" mapstructure:"cron"`
	CostEstimatorPath          string             `json:"cost_estimator_path" mapstructure:"cost_estimator_path"`
	LogLevel                   string             `json:"log_level" mapstructure:"log_level"`
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

// AutonomousSystem ...
type AutonomousSystem struct {
	Number       int `json:"number" mapstructure:"number"`
	LinkCapacity int `json:"link_capacity" mapstructure:"link_capacity"`
	Cost         int `json:"cost" mapstructure:"cost"`
}

// SaveAsJSON ...
func (c *Config) SaveAsJSON(path string) error {

	ext := filepath.Ext(filepath.Base(path))
	path = strings.TrimSuffix(path, ext)
	enc, err := jsonutil.EncodeJSONWithIndentation(*c)
	if err != nil {
		err = stacktrace.Propagate(err, "could not encode config file '%s' as json", path)
		return err
	}
	path = path + ".json"
	// checking to see if target exists
	// delete if stat was successful (i.e exists ...)
	// fmt.Println("SaveAsJSON target to stat", path)
	_, err = os.Stat(path)
	if err == nil {
		err = os.Remove(path)
		if err != nil {
			stacktrace.Propagate(err, "could not remove old json config at %s", path)
			// fmt.Println("err", err)
			return err
		}
	}
	sink, err := os.Create(path)
	if err != nil {
		err = stacktrace.Propagate(err, "could not get a file handle at '%s' ", path)
		// cleaning up file that was wrongly created
		os.Remove(path)
		return err
	}
	defer sink.Close()
	buf := bytes.NewBuffer(enc)
	_, err = io.Copy(sink, buf)
	if err != nil {
		err = stacktrace.Propagate(err, "could not copy encoded config data from buffer to '%s' ", path)
		// cleaning up file that was wrongly created
		os.Remove(path)
		return err
	}
	err = sink.Sync()
	if err != nil {
		err = stacktrace.Propagate(err, "could not flush kernel buffer to disk '%s' ", path)
		// cleaning up file that was wrongly created
		os.Remove(path)
		return err
	}
	return nil
}

// ConfigExtension ...
type ConfigExtension int

const (
	// JSON ...
	JSON ConfigExtension = iota
	// TOML ...
	TOML
	// XML ...
	XML
	// CONF ...
	CONF
)

// String ...
func (e ConfigExtension) String() string {
	switch e {
	case JSON:
		return "json"
	case TOML:
		return "toml"
	case XML:
		return "xml"
	case CONF:
		return "conf"
	default:
		return "unknown"
	}
}
