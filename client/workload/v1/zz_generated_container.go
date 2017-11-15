package client

const (
	ContainerType                          = "container"
	ContainerFieldAdd                      = "add"
	ContainerFieldAllowPrivilegeEscalation = "allowPrivilegeEscalation"
	ContainerFieldArgs                     = "args"
	ContainerFieldCommand                  = "command"
	ContainerFieldDrop                     = "drop"
	ContainerFieldEnvironment              = "environment"
	ContainerFieldEnvironmentFrom          = "environmentFrom"
	ContainerFieldImage                    = "image"
	ContainerFieldImagePullPolicy          = "imagePullPolicy"
	ContainerFieldInitContainer            = "initContainer"
	ContainerFieldLivenessProbe            = "livenessProbe"
	ContainerFieldName                     = "name"
	ContainerFieldPorts                    = "ports"
	ContainerFieldPostStart                = "postStart"
	ContainerFieldPreStop                  = "preStop"
	ContainerFieldPrivileged               = "privileged"
	ContainerFieldReadOnlyRootFilesystem   = "readOnlyRootFilesystem"
	ContainerFieldReadinessProbe           = "readinessProbe"
	ContainerFieldResources                = "resources"
	ContainerFieldRunAsNonRoot             = "runAsNonRoot"
	ContainerFieldRunAsUser                = "runAsUser"
	ContainerFieldScheduling               = "scheduling"
	ContainerFieldStdin                    = "stdin"
	ContainerFieldStdinOnce                = "stdinOnce"
	ContainerFieldTTY                      = "tty"
	ContainerFieldTerminationMessagePath   = "terminationMessagePath"
	ContainerFieldTerminationMessagePolicy = "terminationMessagePolicy"
	ContainerFieldVolumeMounts             = "volumeMounts"
	ContainerFieldWorkingDir               = "workingDir"
)

type Container struct {
	Add                      []string          `json:"add,omitempty"`
	AllowPrivilegeEscalation *bool             `json:"allowPrivilegeEscalation,omitempty"`
	Args                     []string          `json:"args,omitempty"`
	Command                  []string          `json:"command,omitempty"`
	Drop                     []string          `json:"drop,omitempty"`
	Environment              map[string]string `json:"environment,omitempty"`
	EnvironmentFrom          []EnvironmentFrom `json:"environmentFrom,omitempty"`
	Image                    string            `json:"image,omitempty"`
	ImagePullPolicy          string            `json:"imagePullPolicy,omitempty"`
	InitContainer            bool              `json:"initContainer,omitempty"`
	LivenessProbe            *Probe            `json:"livenessProbe,omitempty"`
	Name                     string            `json:"name,omitempty"`
	Ports                    []ContainerPort   `json:"ports,omitempty"`
	PostStart                *Handler          `json:"postStart,omitempty"`
	PreStop                  *Handler          `json:"preStop,omitempty"`
	Privileged               *bool             `json:"privileged,omitempty"`
	ReadOnlyRootFilesystem   *bool             `json:"readOnlyRootFilesystem,omitempty"`
	ReadinessProbe           *Probe            `json:"readinessProbe,omitempty"`
	Resources                *Resources        `json:"resources,omitempty"`
	RunAsNonRoot             *bool             `json:"runAsNonRoot,omitempty"`
	RunAsUser                *int64            `json:"runAsUser,omitempty"`
	Scheduling               *Scheduling       `json:"scheduling,omitempty"`
	Stdin                    bool              `json:"stdin,omitempty"`
	StdinOnce                bool              `json:"stdinOnce,omitempty"`
	TTY                      bool              `json:"tty,omitempty"`
	TerminationMessagePath   string            `json:"terminationMessagePath,omitempty"`
	TerminationMessagePolicy string            `json:"terminationMessagePolicy,omitempty"`
	VolumeMounts             []VolumeMount     `json:"volumeMounts,omitempty"`
	WorkingDir               string            `json:"workingDir,omitempty"`
}
