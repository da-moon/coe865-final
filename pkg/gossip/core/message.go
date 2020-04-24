package core

import (
	"time"

	lamportclock "github.com/da-moon/coe865-final/pkg/gossip/lamport-clock"
)

const (
	queryFlagAck uint32 = 1 << iota
	queryFlagNoBroadcast
)

// messageType are the types of gossip messages Serf will send along
// memberlist.
type messageType uint8

const (
	messageLeaveType messageType = iota
	messageJoinType
)

// messageQuery is used for query events
type messageQuery struct {
	// Event lamport time
	LTime       lamportclock.LamportTime
	MessageType messageType
	// Query ID, randomly generated
	ID uint32
	// Used to set the number of duplicate relayed responses
	RelayFactor uint8
	// Used to provide various flags
	Flags uint32
	// Maximum time between delivery and response
	Timeout time.Duration
	// Query name
	Name string
	// Query payload
	Payload []byte
}

// Ack checks if the ack flag is set
func (m *messageQuery) Ack() bool {

	return (m.Flags & queryFlagAck) != 0
}

// NoBroadcast checks if the no broadcast flag is set
func (m *messageQuery) NoBroadcast() bool {

	return (m.Flags & queryFlagNoBroadcast) != 0
}
func decodeMessage(buf []byte, out interface{}) error {

	return nil
}
func encodeMessage(t messageType, msg interface{}) ([]byte, error) {

	return nil, nil
}

// messageQueryResponse is used to respond to a query
type messageQueryResponse struct {
	// Event lamport time
	LTime       lamportclock.LamportTime
	MessageType messageType
	// Query ID
	ID uint32
	// Node name
	From string
	// Used to provide various flags
	Flags uint32
	// Optional response payload
	Payload []byte
}

// Ack checks if the ack flag is set
func (m *messageQueryResponse) Ack() bool {

	return (m.Flags & queryFlagAck) != 0
}
