package core

import (
	"fmt"

	lamportclock "github.com/da-moon/coe865-final/pkg/gossip/lamport-clock"
)

// Event ...
type Event interface {
	EventType() EventType
	String() string
}

// EventType ...
type EventType int

const (
	// EventPeerAdded ...
	EventHello EventType = iota
	// Update ...
	EventUpdate
	// Leave ...
	EventLeave
)

// String ...
func (e EventType) String() string {

	switch e {
	case EventHello:
		return "hello"
	case EventUpdate:
		return "update"
	case EventLeave:
		return "leave"
	default:
		return (fmt.Sprintf("unknown event type: %d", e))
	}
}

// GossipEvent ...
type GossipEvent struct {
	LTime   lamportclock.LamportTime
	Name    string
	Payload []byte
}

// UpdateEvent ...
type UpdateEvent GossipEvent

// EventType ...
func (u UpdateEvent) EventType() EventType {

	return EventUpdate
}

// String ...
func (u UpdateEvent) String() string {

	return fmt.Sprintf("update: %s", u.Name)
}

// LeaveEvent ...
type LeaveEvent GossipEvent

// EventType ...
func (u LeaveEvent) EventType() EventType {

	return EventLeave
}

// String ...
func (u LeaveEvent) String() string {

	return fmt.Sprintf("leave: %s", u.Name)
}
