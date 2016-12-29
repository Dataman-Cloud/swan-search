package swan

import (
	"fmt"

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

	request, err := r.apiRequest("GET", fmt.Sprintf("%s/%s", r.swanAddr, defaultEventsURL), nil)
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
				channel <- &ev
			case err := <-stream.Errors:
				// TODO let the user handle this error instead of logging it here
				//r.debugLog.Printf("registerSSESubscription(): failed to receive event: %v\n", err)
				fmt.Printf("registerSSESubscription(): failed to receive event: %v\n", err)
			}
		}
	}()
	return nil
}
