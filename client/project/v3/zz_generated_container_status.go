package client

const (
	ContainerStatusType                      = "containerStatus"
	ContainerStatusFieldContainerID          = "containerID"
	ContainerStatusFieldImage                = "image"
	ContainerStatusFieldImageID              = "imageID"
	ContainerStatusFieldLastTerminationState = "lastState"
	ContainerStatusFieldName                 = "name"
	ContainerStatusFieldReady                = "ready"
	ContainerStatusFieldRestartCount         = "restartCount"
	ContainerStatusFieldState                = "state"
)

type ContainerStatus struct {
	ContainerID          string          `json:"containerID,omitempty" yaml:"container_id,omitempty"`
	Image                string          `json:"image,omitempty" yaml:"image,omitempty"`
	ImageID              string          `json:"imageID,omitempty" yaml:"image_id,omitempty"`
	LastTerminationState *ContainerState `json:"lastState,omitempty" yaml:"last_state,omitempty"`
	Name                 string          `json:"name,omitempty" yaml:"name,omitempty"`
	Ready                bool            `json:"ready,omitempty" yaml:"ready,omitempty"`
	RestartCount         int64           `json:"restartCount,omitempty" yaml:"restart_count,omitempty"`
	State                *ContainerState `json:"state,omitempty" yaml:"state,omitempty"`
}
