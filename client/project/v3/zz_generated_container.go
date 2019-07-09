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
	ContainerFieldExitCode                 = "exitCode"
	ContainerFieldImage                    = "image"
	ContainerFieldImagePullPolicy          = "imagePullPolicy"
	ContainerFieldInitContainer            = "initContainer"
	ContainerFieldLivenessProbe            = "livenessProbe"
	ContainerFieldName                     = "name"
	ContainerFieldPorts                    = "ports"
	ContainerFieldPostStart                = "postStart"
	ContainerFieldPreStop                  = "preStop"
	ContainerFieldPrivileged               = "privileged"
	ContainerFieldProcMount                = "procMount"
	ContainerFieldReadOnly                 = "readOnly"
	ContainerFieldReadinessProbe           = "readinessProbe"
	ContainerFieldResources                = "resources"
	ContainerFieldRestartCount             = "restartCount"
	ContainerFieldRunAsGroup               = "runAsGroup"
	ContainerFieldRunAsNonRoot             = "runAsNonRoot"
	ContainerFieldState                    = "state"
	ContainerFieldStdin                    = "stdin"
	ContainerFieldStdinOnce                = "stdinOnce"
	ContainerFieldTTY                      = "tty"
	ContainerFieldTerminationMessagePath   = "terminationMessagePath"
	ContainerFieldTerminationMessagePolicy = "terminationMessagePolicy"
	ContainerFieldTransitioning            = "transitioning"
	ContainerFieldTransitioningMessage     = "transitioningMessage"
	ContainerFieldUid                      = "uid"
	ContainerFieldVolumeDevices            = "volumeDevices"
	ContainerFieldVolumeMounts             = "volumeMounts"
	ContainerFieldWorkingDir               = "workingDir"
)

type Container struct {
	AllowPrivilegeEscalation *bool                 `json:"allowPrivilegeEscalation,omitempty" yaml:"allow_privilege_escalation,omitempty"`
	CapAdd                   []string              `json:"capAdd,omitempty" yaml:"cap_add,omitempty"`
	CapDrop                  []string              `json:"capDrop,omitempty" yaml:"cap_drop,omitempty"`
	Command                  []string              `json:"command,omitempty" yaml:"command,omitempty"`
	Entrypoint               []string              `json:"entrypoint,omitempty" yaml:"entrypoint,omitempty"`
	Environment              map[string]string     `json:"environment,omitempty" yaml:"environment,omitempty"`
	EnvironmentFrom          []EnvironmentFrom     `json:"environmentFrom,omitempty" yaml:"environment_from,omitempty"`
	ExitCode                 *int64                `json:"exitCode,omitempty" yaml:"exit_code,omitempty"`
	Image                    string                `json:"image,omitempty" yaml:"image,omitempty"`
	ImagePullPolicy          string                `json:"imagePullPolicy,omitempty" yaml:"image_pull_policy,omitempty"`
	InitContainer            bool                  `json:"initContainer,omitempty" yaml:"init_container,omitempty"`
	LivenessProbe            *Probe                `json:"livenessProbe,omitempty" yaml:"liveness_probe,omitempty"`
	Name                     string                `json:"name,omitempty" yaml:"name,omitempty"`
	Ports                    []ContainerPort       `json:"ports,omitempty" yaml:"ports,omitempty"`
	PostStart                *Handler              `json:"postStart,omitempty" yaml:"post_start,omitempty"`
	PreStop                  *Handler              `json:"preStop,omitempty" yaml:"pre_stop,omitempty"`
	Privileged               *bool                 `json:"privileged,omitempty" yaml:"privileged,omitempty"`
	ProcMount                string                `json:"procMount,omitempty" yaml:"proc_mount,omitempty"`
	ReadOnly                 *bool                 `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	ReadinessProbe           *Probe                `json:"readinessProbe,omitempty" yaml:"readiness_probe,omitempty"`
	Resources                *ResourceRequirements `json:"resources,omitempty" yaml:"resources,omitempty"`
	RestartCount             int64                 `json:"restartCount,omitempty" yaml:"restart_count,omitempty"`
	RunAsGroup               *int64                `json:"runAsGroup,omitempty" yaml:"run_as_group,omitempty"`
	RunAsNonRoot             *bool                 `json:"runAsNonRoot,omitempty" yaml:"run_as_non_root,omitempty"`
	State                    string                `json:"state,omitempty" yaml:"state,omitempty"`
	Stdin                    bool                  `json:"stdin,omitempty" yaml:"stdin,omitempty"`
	StdinOnce                bool                  `json:"stdinOnce,omitempty" yaml:"stdin_once,omitempty"`
	TTY                      bool                  `json:"tty,omitempty" yaml:"tty,omitempty"`
	TerminationMessagePath   string                `json:"terminationMessagePath,omitempty" yaml:"termination_message_path,omitempty"`
	TerminationMessagePolicy string                `json:"terminationMessagePolicy,omitempty" yaml:"termination_message_policy,omitempty"`
	Transitioning            string                `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage     string                `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`
	Uid                      *int64                `json:"uid,omitempty" yaml:"uid,omitempty"`
	VolumeDevices            []VolumeDevice        `json:"volumeDevices,omitempty" yaml:"volume_devices,omitempty"`
	VolumeMounts             []VolumeMount         `json:"volumeMounts,omitempty" yaml:"volume_mounts,omitempty"`
	WorkingDir               string                `json:"workingDir,omitempty" yaml:"working_dir,omitempty"`
}
