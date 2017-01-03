package swan

import (
	"errors"
	"fmt"
)

// EventsChannel is a channel to receive events upon
type EventsChannel chan *Event

// EventType is a wrapper for a swan event
//type EventType struct {
//	Type string `json:"Type"`
//}

type Event struct {
	ID    string
	Event string
	Data  interface{}
}

type TaskInfo struct {
	Ip     string
	TaskId string
	Port   string
	Type   string
}

func GetEvent(eventType string) (*Event, error) {
	event := new(Event)
	switch eventType {
	case "task_rm":
		event.Data = new(TaskInfo)
	case "task_add":
		event.Data = new(TaskInfo)
	case "app_add":
		event.Data = new(Application)
	case "app_rm":
		event.Data = new(Application)
	default:
		return nil, errors.New(fmt.Sprintf("no event is found for eventType:%s", eventType))
	}
	return event, nil
}
