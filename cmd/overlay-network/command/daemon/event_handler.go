package daemon

import (
	"github.com/da-moon/coe865-final/pkg/gossip/core"
)

// EventHandler is a handler that does things when events happen.
type EventHandler interface {
	HandleEvent(core.Event)
}
