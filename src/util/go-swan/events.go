package swan

// Event is the definition for a event in swan
type Event struct {
	Id      string
	Type    string
	Payload interface{}
}

// EventsChannel is a channel to receive events upon
type EventsChannel chan *Event

// EventType is a wrapper for a swan event
//type EventType struct {
//	Type string `json:"Type"`
//}
