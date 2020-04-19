package config

import (
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/palantir/stacktrace"
)

// DecodeRawConfig ...
func (c *ConfigFactory) DecodeRawConfig(r io.Reader) (*Config, error) {

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		err = stacktrace.Propagate(err, "decode failed due to being unable to read from raw config file")
		return nil, err
	}
	lines := strings.Split(string(buf), "\n")
	index := 0
	var self *RouteController
	connectedRouteControllers := make([]RouteController, 0)
	connectedAutonomousSystems := make([]AutonomousSystem, 0)
	if len(lines) > index {
		// read self config
		self, err = ExtractRouteControllerFromLine(lines[index])
		if err != nil {
			return nil, err
		}
		index = index + 1
	}
	// populating connected rcs slice
	if len(lines) > index {
		parts := strings.Split(strings.TrimSpace(trimComment(lines[index])), " ")
		count, err := strconv.Atoi(parts[0])
		if err != nil {
			err = stacktrace.Propagate(err, "could not convert value in number of connected rcs '%s' to integer", parts[0])
			return nil, err
		}
		index = index + 1
		for i := 0; i < count; i++ {
			// geting a single rc
			rc, err := ExtractRouteControllerFromLine(lines[index])
			if err != nil {
				return nil, err
			}
			if rc != nil {
				connectedRouteControllers = append(connectedRouteControllers, *rc)
			}
			index = index + 1
		}
	}
	// reading connected asns
	if len(lines) > index {
		parts := strings.Split(strings.TrimSpace(trimComment(lines[index])), " ")
		count, err := strconv.Atoi(parts[0])
		if err != nil {
			err = stacktrace.Propagate(err, "could not convert value in number of connected asns '%s' to integer", parts[0])
			return nil, err
		}
		index = index + 1
		for i := 0; i < count; i++ {
			// geting a single autonomous system
			as, err := ExtractAutonomousSystemFromLine(lines[index])
			if err != nil {
				return nil, err
			}
			if as != nil {
				connectedAutonomousSystems = append(connectedAutonomousSystems, *as)
			}
			index = index + 1
		}
	}
	result := c.New(self, connectedRouteControllers, connectedAutonomousSystems)
	return result, nil
}

// ExtractRouteControllerFromLine ...
func ExtractRouteControllerFromLine(input string) (*RouteController, error) {
	result := &RouteController{}
	parts := SanitizeAndSplitLine(input)
	// fmt.Println("parts", parts, len(parts))
	if len(parts) != 3 {
		err := stacktrace.NewError("wrong number of parts '%d' in given line. need 3 parts. possible issue with :space: delimiter", len(parts))
		return nil, err
	}
	id, err := strconv.Atoi(parts[0])
	if err != nil {
		err = stacktrace.Propagate(err, "could not convert value in ID field '%s' to integer", parts[0])
		return nil, err
	}
	asn, err := strconv.Atoi(parts[1])
	if err != nil {
		err = stacktrace.Propagate(err, "could not convert value in Autonomous System Number field '%s' to integer", parts[0])
		return nil, err
	}
	result.ID = id
	result.AutonomousSystemNumber = asn
	result.IP = parts[2]
	return result, nil
}

// ExtractAutonomousSystemFromLine ...
func ExtractAutonomousSystemFromLine(input string) (*AutonomousSystem, error) {
	result := &AutonomousSystem{}
	parts := SanitizeAndSplitLine(input)
	if len(parts) != 3 {
		err := stacktrace.NewError("wrong number of parts '%d' in given line. need 3 parts. possible issue with :space: delimiter")
		return nil, err
	}
	number, err := strconv.Atoi(parts[0])
	if err != nil {
		err = stacktrace.Propagate(err, "could not convert value in number field '%s' to interger", parts[0])
		return nil, err
	}
	linkCapacity, err := strconv.Atoi(parts[1])
	if err != nil {
		err = stacktrace.Propagate(err, "could not convert value in link capacity '%s' to interger", parts[1])
		return nil, err
	}
	cost, err := strconv.Atoi(parts[2])
	if err != nil {
		err = stacktrace.Propagate(err, "could not convert value in cost '%s' to interger", parts[2])
		return nil, err
	}
	result.Number = number
	result.LinkCapacity = linkCapacity
	result.Cost = cost
	return result, nil

}

// SanitizeAndSplitLine ...
func SanitizeAndSplitLine(input string) []string {
	input = strings.TrimSpace(trimComment(input))

	parts := strings.Split(input, " ")
	return parts
}
func trimComment(s string) string {
	result := s
	idx := strings.Index(s, ";")
	if idx != -1 {
		result = result[:idx]
	}
	idx = strings.Index(s, "#")
	if idx != -1 {
		result = result[:idx]
	}
	return result
}
