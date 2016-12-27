package swan

const (
	defaultEventsURL = "/event"

	/* --- api related constants --- */
	swanAPIVersion      = "v_beta"
	swanAPIEventStream  = swanAPIVersion + "/events"
	swanAPISubscription = swanAPIVersion + "/eventSubscriptions"
	swanAPIApps         = swanAPIVersion + "/apps"
	swanAPITasks        = swanAPIVersion + "/tasks"
	swanAPIPing         = "ping"
)
