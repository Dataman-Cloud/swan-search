package swan

import (
	"fmt"

	"github.com/donovanhide/eventsource"
)

// AddEventsListener adds your self as a listener to events from Marathon
// channel: a EventsChannel used to receive event on
func (r *swanClient) AddEventsListener(filter int) (EventsChannel, error) {
	r.Lock()
	defer r.Unlock()

	// step: someone has asked to start listening to event, we need to register for events
	// if we haven't done so already
	if err := r.registerSSESubscription(); err != nil {
		return nil, err
	}

	channel := make(EventsChannel)
	//r.listeners[channel] = EventsChannelContext{
	//	filter:     filter,
	//	done:       make(chan struct{}, 1),
	//	completion: &sync.WaitGroup{},
	//}
	return channel, nil
}

func (r *swanClient) registerSSESubscription() error {
	// Prevent multiple SSE subscriptions
	if r.subscribedToSSE {
		return nil
	}
	// Get a member from the cluster
	//marathon, err := r.hosts.getMember()
	//if err != nil {
	//	return err
	//}
	fmt.Printf("%s/%s", r.swanAddr, defaultEventsURL)
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
				if err := r.handleEvent(ev.Data()); err != nil {
					// TODO let the user handle this error instead of logging it here
					//r.debugLog.Printf("registerSSESubscription(): failed to handle event: %v\n", err)
					fmt.Printf("registerSSESubscription(): failed to handle event: %v\n", err)
				}
			case err := <-stream.Errors:
				// TODO let the user handle this error instead of logging it here
				//r.debugLog.Printf("registerSSESubscription(): failed to receive event: %v\n", err)
				fmt.Printf("registerSSESubscription(): failed to receive event: %v\n", err)
			}
		}
	}()

	r.subscribedToSSE = true
	return nil
}

func (r *swanClient) handleEvent(content string) error {
	// step: process and decode the event
	fmt.Printf("content:%+v\n", content)
	//eventType := new(EventType)
	//err := json.NewDecoder(strings.NewReader(content)).Decode(eventType)
	//if err != nil {
	//	return fmt.Errorf("failed to decode the event type, content: %s, error: %s", content, err)
	//}

	//// step: check whether event type is handled
	//event, err := GetEvent(eventType.EventType)
	//if err != nil {
	//	return fmt.Errorf("unable to handle event, type: %s, error: %s", eventType.EventType, err)
	//}

	//// step: let's decode message
	//err = json.NewDecoder(strings.NewReader(content)).Decode(event.Event)
	//if err != nil {
	//	return fmt.Errorf("failed to decode the event, id: %d, error: %s", event.ID, err)
	//}

	//r.RLock()
	//defer r.RUnlock()

	//// step: check if anyone is listen for this event
	//for channel, context := range r.listeners {
	//	// step: check if this listener wants this event type
	//	if event.ID&context.filter != 0 {
	//		context.completion.Add(1)
	//		go func(ch EventsChannel, context EventsChannelContext, e *Event) {
	//			defer context.completion.Done()
	//			select {
	//			case ch <- e:
	//			case <-context.done:
	//				// Terminates goroutine.
	//			}
	//		}(channel, context, event)
	//	}
	//}
	return nil
}
