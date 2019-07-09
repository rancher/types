package client

const (
	ContainerStateTerminatedType             = "containerStateTerminated"
	ContainerStateTerminatedFieldContainerID = "containerID"
	ContainerStateTerminatedFieldExitCode    = "exitCode"
	ContainerStateTerminatedFieldFinishedAt  = "finishedAt"
	ContainerStateTerminatedFieldMessage     = "message"
	ContainerStateTerminatedFieldReason      = "reason"
	ContainerStateTerminatedFieldSignal      = "signal"
	ContainerStateTerminatedFieldStartedAt   = "startedAt"
)

type ContainerStateTerminated struct {
	ContainerID string `json:"containerID,omitempty" yaml:"container_id,omitempty"`
	ExitCode    int64  `json:"exitCode,omitempty" yaml:"exit_code,omitempty"`
	FinishedAt  string `json:"finishedAt,omitempty" yaml:"finished_at,omitempty"`
	Message     string `json:"message,omitempty" yaml:"message,omitempty"`
	Reason      string `json:"reason,omitempty" yaml:"reason,omitempty"`
	Signal      int64  `json:"signal,omitempty" yaml:"signal,omitempty"`
	StartedAt   string `json:"startedAt,omitempty" yaml:"started_at,omitempty"`
}
