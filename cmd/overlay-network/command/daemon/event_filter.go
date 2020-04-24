package daemon

import (
	"strings"

	"github.com/da-moon/coe865-final/pkg/gossip/core"
)

// EventFilter ...
type EventFilter struct {
	Event string
	Name  string
}

// ParseEventFilter a string with the event type filters and
// parses it into a series of EventFilters if it can.
func ParseEventFilter(v string) []EventFilter {

	// No filter translates to stream all
	if v == "" {
		v = "*"
	}
	events := strings.Split(v, ",")
	results := make([]EventFilter, 0, len(events))
	for _, event := range events {
		var result EventFilter
		var name string
		if strings.HasPrefix(event, "hello:") {
			name = event[len("hello:"):]
			event = "hello"
		} else if strings.HasPrefix(event, "update:") {
			name = event[len("update:"):]
			event = "update"
		} else if strings.HasPrefix(event, "leave:") {
			name = event[len("leave:"):]
			event = "leave"
		}
		result.Event = event
		result.Name = name
		results = append(results, result)
	}
	return results
}

// Invoke ...
func (s *EventFilter) Invoke(e core.Event) bool {

	if s.Event == "*" {
		return true
	}
	if e.EventType().String() != s.Event {
		return false
	}
	if s.Event == "hello" && s.Name != "" {
		event, ok := e.(core.HelloEvent)
		if !ok {
			return false
		}
		if event.Name != s.Name {
			return false
		}
	}
	if s.Event == "update" && s.Name != "" {
		event, ok := e.(core.UpdateEvent)
		if !ok {
			return false
		}
		if event.Name != s.Name {
			return false
		}
	}
	if s.Event == "leave" && s.Name != "" {
		event, ok := e.(core.LeaveEvent)
		if !ok {
			return false
		}
		if event.Name != s.Name {
			return false
		}
	}
	return true
}

// Valid checks if this is a valid agent event script.
func (s *EventFilter) Valid() bool {

	switch s.Event {
	case "hello":
	case "update":
	case "leave":
	case "*":
	default:
		return false
	}
	return true
}
