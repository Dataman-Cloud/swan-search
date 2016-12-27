package swan

import (
	global "github.com/Dataman-Cloud/borgsphere/src/util/config"

	"golang.org/x/net/context"
)

const (
	AppStatusOK     = 0
	AppBackendError = 22001
	AppUnauthorized = 22002
)

type Setter interface {
	Set(string, interface{})
}

// returns the Store associated with this context
func FromContext(c context.Context) SwanControllerInterface {
	return c.Value(global.SwanHandlerKey()).(SwanControllerInterface)
}

func ToContext(c Setter, swanControllerIml SwanControllerInterface) {
	c.Set(global.SwanHandlerKey(), swanControllerIml)
}
