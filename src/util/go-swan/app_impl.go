package swan

import (
	"fmt"
	"time"

	log "github.com/Dataman-Cloud/borgsphere/src/util/log"
	//marathon "github.com/gambol99/go-marathon"
)

const (
	ListSlaveContainerUrl    = "http://%s:%s/containers"
	MesosContainerNameFromat = "mesos-%s.%s"

	defaultHttpRequestTimeout = time.Second * 5
)

type SwanControllerImpl struct {
	Client Swan

	// use SwanUrl for PoC
	SwanUrl string
}

type SlaveContainer struct {
	ContainerID string `json:"container_id"`
	ExecutorID  string `json:"executor_id"`
	FrameworkID string `json:"framework_id"`
	Source      string `json:"source"`
}

func NewSwanControllerImpl(swanAddr string) *SwanControllerImpl {
	swanURL := fmt.Sprintf("http://%s", swanAddr)
	client, err := NewClient(swanURL)
	if err != nil {
		log.L.Error("Failed to create a client for marathon, error: %s", err)
	}

	return &SwanControllerImpl{
		Client:  client,
		SwanUrl: swanAddr,
	}
}

func (app *SwanControllerImpl) Applications() ([]*Application, error) {
	return app.Client.Applications()
}

func (app *SwanControllerImpl) CreateApplication(version *Version) (*Application, error) {
	return app.Client.CreateApplication(version)
}
