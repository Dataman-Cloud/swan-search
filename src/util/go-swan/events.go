package swan

import (
	"github.com/donovanhide/eventsource"
)

// EventsChannel is a channel to receive events upon
type EventsChannel chan *eventsource.Event

// EventType is a wrapper for a swan event
//type EventType struct {
//	Type string `json:"Type"`
//}
