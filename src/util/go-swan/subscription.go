package swan

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/donovanhide/eventsource"
)

// AddEventsListener adds your self as a listener to events from Marathon
// channel: a EventsChannel used to receive event on
func (r *swanClient) AddEventsListener() (EventsChannel, error) {
	r.Lock()
	defer r.Unlock()

	channel := make(EventsChannel)
	if err := r.registerSSESubscription(channel); err != nil {
		return nil, err
	}

	return channel, nil
}

func (r *swanClient) registerSSESubscription(channel EventsChannel) error {
	// Prevent multiple SSE subscriptions

	url, err := r.hosts.getMember()
	if err != nil {
		return err
	}

	request, err := r.apiRequest("GET", fmt.Sprintf("%s/%s", url, defaultEventsURL), nil)
	if err != nil {
		return err
	}

	// Try to connect to stream, reusing the http client settings
	stream, err := eventsource.SubscribeWith("", r.httpClient, request)
	if err != nil {
		fmt.Println("err when event request to /events")
		return err
	}

	go func() {
		for {
			select {
			case ev := <-stream.Events:
				//if err := r.handleEvent(ev.Id(), ev.Event(), ev.Data()); err != nil {
				//	// TODO let the user handle this error instead of logging it here
				//r.debugLog.Printf("registerSSESubscription(): failed to handle event: %v\n", err)
				//fmt.Printf("registerSSESubscription(): failed to handle event: %v\n", err)
				//}
				event, err := GetEvent(ev.Event())
				if err != nil {
					fmt.Errorf("failed to handle event:%s", err)
				}
				event.ID = ev.Id()
				event.Event = ev.Event()
				err = json.NewDecoder(strings.NewReader(ev.Data())).Decode(event.Data)
				if err != nil {
					fmt.Errorf("failed to decode the event, eventType: %d, error: %s", event.Event, err)
				}
				channel <- event
			case err := <-stream.Errors:
				// TODO let the user handle this error instead of logging it here
				//r.debugLog.Printf("registerSSESubscription(): failed to receive event: %v\n", err)
				fmt.Errorf("registerSSESubscription(): failed to receive event: %s", err)
			}
		}
	}()
	return nil
}
