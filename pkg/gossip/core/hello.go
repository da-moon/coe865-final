package core

import (
	"fmt"

	"github.com/da-moon/coe865-final/pkg/gossip/sentry"
	"github.com/da-moon/coe865-final/pkg/jsonutil"
	"github.com/palantir/stacktrace"
)

// HelloPayload ...
type HelloPayload struct {
	YourAddr string
}

// HelloEvent ...
type HelloEvent GossipEvent

// EventType ...
func (u HelloEvent) EventType() EventType {
	return EventHello
}

// String ...
func (u HelloEvent) String() string {
	return fmt.Sprintf("hello: %s", u.Name)
}

// AgentHelloEvent ...
func (c *Config) AgentHelloEvent(address string) (*sentry.SignedMessage, error) {
	c.Logger.Printf("[INFO] core: generating HelloMessage for address '%v'", address)
	payload, err := jsonutil.EncodeJSON(&HelloPayload{
		YourAddr: address,
	})
	if err != nil {
		err = stacktrace.Propagate(err, "could not generate hello event for address '%s'", address)
		return nil, err
	}
	evt := HelloEvent{
		LTime:   c.eventClock.Time(),
		Name:    EventHello.String(),
		Payload: payload,
	}
	result, err := c.sentry.NewMessage(evt)
	if err != nil {
		err = stacktrace.Propagate(err, "sentry could not sign message %v", evt)
		return nil, err
	}
	c.ExternalEventCh <- evt
	return result, nil
}
