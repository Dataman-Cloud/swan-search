package swan

import (
// TODO(xychu): add go-swan here
// marathon "github.com/gambol99/go-marathon"
)

type SwanControllerInterface interface {
	//ListApplications(map[string][]string) ([]string, error)
	Applications() ([]*Application, error)
	CreateApplication(*Version) (*Application, error)
	//DeleteApplication(string, bool) (*marathon.DeploymentID, error)
	//UpdateApplication(*marathon.Application, bool) (*marathon.DeploymentID, error)
	//ApplicationBy(string, *marathon.GetAppOpts) (*marathon.Application, error)
	//RestartApplication(string, bool) (*marathon.DeploymentID, error)
	//ApplicationVersions(string) (*marathon.ApplicationVersions, error)
	//ApplicationByVersion(string, string) (*marathon.Application, error)
	//Tasks(string) (*marathon.Tasks, error)
	//KillTasks([]string, *marathon.KillTaskOpts) error
	//Queue() (*marathon.Queue, error)
}
