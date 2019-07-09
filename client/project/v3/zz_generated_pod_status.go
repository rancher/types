package client

const (
	PodStatusType                       = "podStatus"
	PodStatusFieldConditions            = "conditions"
	PodStatusFieldContainerStatuses     = "containerStatuses"
	PodStatusFieldInitContainerStatuses = "initContainerStatuses"
	PodStatusFieldMessage               = "message"
	PodStatusFieldNodeIp                = "nodeIp"
	PodStatusFieldNominatedNodeName     = "nominatedNodeName"
	PodStatusFieldPhase                 = "phase"
	PodStatusFieldPodIp                 = "podIp"
	PodStatusFieldQOSClass              = "qosClass"
	PodStatusFieldReason                = "reason"
	PodStatusFieldStartTime             = "startTime"
)

type PodStatus struct {
	Conditions            []PodCondition    `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	ContainerStatuses     []ContainerStatus `json:"containerStatuses,omitempty" yaml:"container_statuses,omitempty"`
	InitContainerStatuses []ContainerStatus `json:"initContainerStatuses,omitempty" yaml:"init_container_statuses,omitempty"`
	Message               string            `json:"message,omitempty" yaml:"message,omitempty"`
	NodeIp                string            `json:"nodeIp,omitempty" yaml:"node_ip,omitempty"`
	NominatedNodeName     string            `json:"nominatedNodeName,omitempty" yaml:"nominated_node_name,omitempty"`
	Phase                 string            `json:"phase,omitempty" yaml:"phase,omitempty"`
	PodIp                 string            `json:"podIp,omitempty" yaml:"pod_ip,omitempty"`
	QOSClass              string            `json:"qosClass,omitempty" yaml:"qos_class,omitempty"`
	Reason                string            `json:"reason,omitempty" yaml:"reason,omitempty"`
	StartTime             string            `json:"startTime,omitempty" yaml:"start_time,omitempty"`
}
