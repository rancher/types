package client

const (
	ContainerType                          = "container"
	ContainerFieldAllowPrivilegeEscalation = "allowPrivilegeEscalation"
	ContainerFieldCapAdd                   = "capAdd"
	ContainerFieldCapDrop                  = "capDrop"
	ContainerFieldCommand                  = "command"
	ContainerFieldEntrypoint               = "entrypoint"
	ContainerFieldEnvironment              = "environment"
	ContainerFieldEnvironmentFrom          = "environmentFrom"
	ContainerFieldHealthcheck              = "healthcheck"
	ContainerFieldImage                    = "image"
	ContainerFieldInitContainer            = "initContainer"
	ContainerFieldName                     = "name"
	ContainerFieldPorts                    = "ports"
	ContainerFieldPostStart                = "postStart"
	ContainerFieldPreStop                  = "preStop"
	ContainerFieldPrivileged               = "privileged"
	ContainerFieldPullPolicy               = "pullPolicy"
	ContainerFieldReadOnly                 = "readOnly"
	ContainerFieldReadycheck               = "readycheck"
	ContainerFieldResources                = "resources"
	ContainerFieldRunAsNonRoot             = "runAsNonRoot"
	ContainerFieldStdin                    = "stdin"
	ContainerFieldStdinOnce                = "stdinOnce"
	ContainerFieldTTY                      = "tty"
	ContainerFieldTerminationMessagePath   = "terminationMessagePath"
	ContainerFieldTerminationMessagePolicy = "terminationMessagePolicy"
	ContainerFieldUid                      = "uid"
	ContainerFieldVolumeMounts             = "volumeMounts"
	ContainerFieldWorkingDir               = "workingDir"
)

type Container struct {
	AllowPrivilegeEscalation *bool             `json:"allowPrivilegeEscalation,omitempty"`
	CapAdd                   []string          `json:"capAdd,omitempty"`
	CapDrop                  []string          `json:"capDrop,omitempty"`
	Command                  []string          `json:"command,omitempty"`
	Entrypoint               []string          `json:"entrypoint,omitempty"`
	Environment              map[string]string `json:"environment,omitempty"`
	EnvironmentFrom          []EnvironmentFrom `json:"environmentFrom,omitempty"`
	Healthcheck              *Probe            `json:"healthcheck,omitempty"`
	Image                    string            `json:"image,omitempty"`
	InitContainer            *bool             `json:"initContainer,omitempty"`
	Name                     string            `json:"name,omitempty"`
	Ports                    []ContainerPort   `json:"ports,omitempty"`
	PostStart                *Handler          `json:"postStart,omitempty"`
	PreStop                  *Handler          `json:"preStop,omitempty"`
	Privileged               *bool             `json:"privileged,omitempty"`
	PullPolicy               string            `json:"pullPolicy,omitempty"`
	ReadOnly                 *bool             `json:"readOnly,omitempty"`
	Readycheck               *Probe            `json:"readycheck,omitempty"`
	Resources                *Resources        `json:"resources,omitempty"`
	RunAsNonRoot             *bool             `json:"runAsNonRoot,omitempty"`
	Stdin                    *bool             `json:"stdin,omitempty"`
	StdinOnce                *bool             `json:"stdinOnce,omitempty"`
	TTY                      *bool             `json:"tty,omitempty"`
	TerminationMessagePath   string            `json:"terminationMessagePath,omitempty"`
	TerminationMessagePolicy string            `json:"terminationMessagePolicy,omitempty"`
	Uid                      *int64            `json:"uid,omitempty"`
	VolumeMounts             []VolumeMount     `json:"volumeMounts,omitempty"`
	WorkingDir               string            `json:"workingDir,omitempty"`
}
