package swan

import (
	"time"
)

type Version struct {
	ID                string
	AppId             string
	PerviousVersionID string
	Command           string
	Cpus              float64
	Mem               float64
	Disk              float64
	Instances         int32
	RunAs             string
	Container         *Container
	Labels            map[string]string
	HealthChecks      []*HealthCheck
	Env               map[string]string
	KillPolicy        *KillPolicy
	UpdatePolicy      *UpdatePolicy
	Constraints       []string
	Uris              []string
	Ip                []string
	Mode              string
}

type Container struct {
	Type    string
	Docker  *Docker
	Volumes []*Volume
}

type Docker struct {
	ForcePullImage bool
	Image          string
	Network        string
	Parameters     []*Parameter
	PortMappings   []*PortMapping
	Privileged     bool
}

type Parameter struct {
	Key   string
	Value string
}

type PortMapping struct {
	ContainerPort int32
	Name          string
	Protocol      string
}

type Volume struct {
	ContainerPath string
	HostPath      string
	Mode          string
}

type KillPolicy struct {
	Duration int64
}

type UpdatePolicy struct {
	UpdateDelay  int32
	MaxRetries   int32
	MaxFailovers int32
	Action       string
}

type HealthCheck struct {
	ID                  string
	Address             string
	TaskID              string
	AppID               string
	Protocol            string
	PortName            string
	Command             *Command
	Path                string
	ConsecutiveFailures uint32
	GracePeriodSeconds  float64
	IntervalSeconds     float64
	TimeoutSeconds      float64
}

type Command struct {
	Value string
}

type Application struct {
	ID               string    `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Instances        int       `json:"instances,omitempty"`
	UpdatedInstances int       `json:"updatedInstances,omitempty"`
	RunningInstances int       `json:"runningInstances"`
	RunAs            string    `json:"runAs,omitempty"`
	ClusterId        string    `json:"clusterId,omitempty"`
	Status           string    `json:"status,omitempty"`
	Created          time.Time `json:"created,omitempty"`
	Updated          time.Time `json:"updated,omitempty"`
	Mode             string    `json:"mode,omitempty"`
	State            string    `json:"state"`

	// use task for compatability now, should be slot here
	Tasks    []*Task  `json:"tasks,omitempty"`
	Versions []string `json:"versions,omitempty"`
	IP       []string `json:"ip,omitempty"`

	// current version related info
	Labels      map[string]string `json:"labels,omitempty"`
	Env         map[string]string `json:"env,omitempty"`
	Constraints []string          `json:"constraints,omitempty"`
	Uris        []string          `json:"uris,omitempty"`
	//HealthChecks      []*types.HealthCheck
	//KillPolicy        *types.KillPolicy
	//UpdatePolicy      *types.UpdatePolicy
}

// use task for compatability now, should be slot here
// and together with task history
type Task struct {
	ID        string `json:"id,omitempty"`
	AppId     string `json:"appId,omitempty"`
	VersionId string `json:"versionId,omitempty"`

	Status string `json:"status,omitempty"`

	OfferId       string `json:"offerId,omitempty"`
	AgentId       string `json:"agentId,omitempty"`
	AgentHostname string `json:"agentHostname,omitempty"`

	Cpu  float64 `json:"cpu,omitempty"`
	Mem  float64 `json:"mem,omitempty"`
	Disk float64 `json:"disk,omitempty"`

	History []*TaskHistory `json:"history,omitempty"`

	IP string `json:"ip,omitempty"`

	Created time.Time `json:"created,omitempty"`

	Image   string `json:"image,omitempty"`
	Healthy bool   `json:"healthy,omitempty"`
}

type TaskHistory struct {
	ID        string `json:"id,omitempty"`
	AppId     string `json:"appId,omitempty"`
	VersionId string `json:"versionId,omitempty"`

	OfferId       string `json:"offerId,omitempty"`
	AgentId       string `json:"agentId,omitempty"`
	AgentHostname string `json:"agentHostname,omitempty"`

	Cpu  float64 `json:"cpu,omitempty"`
	Mem  float64 `json:"mem,omitempty"`
	Disk float64 `json:"disk,omitempty"`

	State  string `json:"state,omitempty"`
	Reason string `json:"Reason,omitempty"`
	Stdout string `json:"stdout,omitempty"`
	Stderr string `json:"stderr,omitempty"`
}

type Stats struct {
	AppCount  int `json:"appCount,omitempty"`
	TaskCount int `json:"taskCount,omitempty"`

	CpuTotalOffered  float64 `json:"cpuTotalOffered,omitempty"`
	MemTotalOffered  float64 `json:"memTotalOffered,omitempty"`
	DiskTotalOffered float64 `json:"diskTotalOffered,omitempty"`

	CpuTotalUsed  float64 `json:"cpuTotalUsed,omitempty"`
	MemTotalUsed  float64 `json:"memTotalUsed,omitempty"`
	DiskTotalUsed float64 `json:"diskTotalUsed,omitempty"`

	AppStats map[string]int `json:"appStats,omitempty"`
}

// AddLabel adds a label to the application
//		name:	the name of the label
//		value: value for this label
func (v *Version) AddLabel(name, value string) *Version {
	if v.Labels == nil {
		v.EmptyLabels()
	}
	v.Labels[name] = value

	return v
}

// EmptyLabels explicitly empties the labels -- use this if you need to empty
// the labels of an application that already has labels set (setting labels to nil will
// keep the current value)
func (v *Version) EmptyLabels() *Version {
	v.Labels = map[string]string{}

	return v
}
