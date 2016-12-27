package swan

type Application struct {
	ID              string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version         *Version `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	ProposedVersion *Version `protobuf:"bytes,4,opt,name=proposedVersion" json:"proposedVersion,omitempty"`
	ClusterId       string   `protobuf:"bytes,5,opt,name=clusterId,proto3" json:"clusterId,omitempty"`
	State           string   `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	CreatedAt       int64    `protobuf:"varint,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt       int64    `protobuf:"varint,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

type Version struct {
	ID                string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PerviousVersionID string            `protobuf:"bytes,2,opt,name=perviousVersionID,proto3" json:"perviousVersionID,omitempty"`
	Command           string            `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
	Cpus              float64           `protobuf:"fixed64,4,opt,name=cpus,proto3" json:"cpus,omitempty"`
	Mem               float64           `protobuf:"fixed64,5,opt,name=mem,proto3" json:"mem,omitempty"`
	Disk              float64           `protobuf:"fixed64,6,opt,name=disk,proto3" json:"disk,omitempty"`
	Instances         int32             `protobuf:"varint,7,opt,name=instances,proto3" json:"instances,omitempty"`
	RunAs             string            `protobuf:"bytes,8,opt,name=runAs,proto3" json:"runAs,omitempty"`
	Container         *Container        `protobuf:"bytes,9,opt,name=container" json:"container,omitempty"`
	Labels            map[string]string `protobuf:"bytes,10,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	HealthChecks      []*HealthCheck    `protobuf:"bytes,11,rep,name=healthChecks" json:"healthChecks,omitempty"`
	Env               map[string]string `protobuf:"bytes,12,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	KillPolicy        *KillPolicy       `protobuf:"bytes,13,opt,name=killPolicy" json:"killPolicy,omitempty"`
	UpdatePolicy      *UpdatePolicy     `protobuf:"bytes,14,opt,name=updatePolicy" json:"updatePolicy,omitempty"`
	Constraints       []string          `protobuf:"bytes,15,rep,name=constraints" json:"constraints,omitempty"`
	Uris              []string          `protobuf:"bytes,16,rep,name=uris" json:"uris,omitempty"`
	Ip                []string          `protobuf:"bytes,17,rep,name=ip" json:"ip,omitempty"`
	Mode              string            `protobuf:"bytes,18,opt,name=mode,proto3" json:"mode,omitempty"`
	AppId             string            `protobuf:"bytes,19,opt,name=appId,proto3" json:"appId,omitempty"`
}

type Container struct {
	Type    string    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Docker  *Docker   `protobuf:"bytes,2,opt,name=docker" json:"docker,omitempty"`
	Volumes []*Volume `protobuf:"bytes,3,rep,name=volumes" json:"volumes,omitempty"`
}

type Docker struct {
	ForcePullImage bool           `protobuf:"varint,1,opt,name=forcePullImage,proto3" json:"forcePullImage,omitempty"`
	Image          string         `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Network        string         `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	Parameters     []*Parameter   `protobuf:"bytes,4,rep,name=parameters" json:"parameters,omitempty"`
	PortMappings   []*PortMapping `protobuf:"bytes,5,rep,name=portMappings" json:"portMappings,omitempty"`
	Privileged     bool           `protobuf:"varint,6,opt,name=privileged,proto3" json:"privileged,omitempty"`
}

type Parameter struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

type PortMapping struct {
	ContainerPort int32  `protobuf:"varint,1,opt,name=containerPort,proto3" json:"containerPort,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Protocol      string `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

type Volume struct {
	ContainerPath string `protobuf:"bytes,1,opt,name=containerPath,proto3" json:"containerPath,omitempty"`
	HostPath      string `protobuf:"bytes,2,opt,name=hostPath,proto3" json:"hostPath,omitempty"`
	Mode          string `protobuf:"bytes,3,opt,name=mode,proto3" json:"mode,omitempty"`
}

type HealthCheck struct {
	ID                     string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address                string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	TaskID                 string   `protobuf:"bytes,3,opt,name=taskid,proto3" json:"taskid,omitempty"`
	AppID                  string   `protobuf:"bytes,4,opt,name=appid,proto3" json:"appid,omitempty"`
	Protocol               string   `protobuf:"bytes,5,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Port                   int32    `protobuf:"varint,6,opt,name=port,proto3" json:"port,omitempty"`
	PortIndex              int32    `protobuf:"varint,7,opt,name=portIndex,proto3" json:"portIndex,omitempty"`
	PortName               string   `protobuf:"bytes,8,opt,name=portName,proto3" json:"portName,omitempty"`
	Command                *Command `protobuf:"bytes,9,opt,name=command" json:"command,omitempty"`
	Path                   string   `protobuf:"bytes,10,opt,name=path,proto3" json:"path,omitempty"`
	MaxConsecutiveFailures uint32   `protobuf:"varint,11,opt,name=maxConsecutiveFailures,proto3" json:"maxConsecutiveFailures,omitempty"`
	GracePeriodSeconds     float64  `protobuf:"fixed64,12,opt,name=gracePeriodSeconds,proto3" json:"gracePeriodSeconds,omitempty"`
	IntervalSeconds        float64  `protobuf:"fixed64,13,opt,name=intervalSeconds,proto3" json:"intervalSeconds,omitempty"`
	TimeoutSeconds         float64  `protobuf:"fixed64,14,opt,name=timeoutSeconds,proto3" json:"timeoutSeconds,omitempty"`
}

type Command struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

type UpdatePolicy struct {
	UpdateDelay  int32  `protobuf:"varint,1,opt,name=updateDelay,proto3" json:"updateDelay,omitempty"`
	MaxRetries   int32  `protobuf:"varint,2,opt,name=maxRetries,proto3" json:"maxRetries,omitempty"`
	MaxFailovers int32  `protobuf:"varint,3,opt,name=maxFailovers,proto3" json:"maxFailovers,omitempty"`
	Action       string `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
}

type KillPolicy struct {
	Duration int64 `protobuf:"varint,1,opt,name=duration,proto3" json:"duration,omitempty"`
}
