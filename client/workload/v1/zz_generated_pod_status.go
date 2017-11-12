package client

const (
	PodStatusType                       = "podStatus"
	PodStatusFieldConditions            = "conditions"
	PodStatusFieldContainerStatuses     = "containerStatuses"
	PodStatusFieldHostIP                = "hostIP"
	PodStatusFieldInitContainerStatuses = "initContainerStatuses"
	PodStatusFieldMessage               = "message"
	PodStatusFieldPhase                 = "phase"
	PodStatusFieldPodIP                 = "podIP"
	PodStatusFieldQOSClass              = "qosClass"
	PodStatusFieldReason                = "reason"
	PodStatusFieldStartTime             = "startTime"
)

type PodStatus struct {
	Conditions            []PodCondition    `json:"conditions,omitempty"`
	ContainerStatuses     []ContainerStatus `json:"containerStatuses,omitempty"`
	HostIP                string            `json:"hostIP,omitempty"`
	InitContainerStatuses []ContainerStatus `json:"initContainerStatuses,omitempty"`
	Message               string            `json:"message,omitempty"`
	Phase                 string            `json:"phase,omitempty"`
	PodIP                 string            `json:"podIP,omitempty"`
	QOSClass              string            `json:"qosClass,omitempty"`
	Reason                string            `json:"reason,omitempty"`
	StartTime             string            `json:"startTime,omitempty"`
}
