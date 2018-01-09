package client

const (
	ClusterStatusType                          = "clusterStatus"
	ClusterStatusFieldAPIEndpoint              = "apiEndpoint"
	ClusterStatusFieldAllocatable              = "allocatable"
	ClusterStatusFieldCapacity                 = "capacity"
	ClusterStatusFieldComponentStatuses        = "componentStatuses"
	ClusterStatusFieldConditions               = "conditions"
	ClusterStatusFieldLimits                   = "limits"
	ClusterStatusFieldRequested                = "requested"
	ClusterStatusFieldServiceAccountSecretName = "serviceAccountSecretName"
)

type ClusterStatus struct {
	APIEndpoint              string                   `json:"apiEndpoint,omitempty"`
	Allocatable              map[string]string        `json:"allocatable,omitempty"`
	Capacity                 map[string]string        `json:"capacity,omitempty"`
	ComponentStatuses        []ClusterComponentStatus `json:"componentStatuses,omitempty"`
	Conditions               []ClusterCondition       `json:"conditions,omitempty"`
	Limits                   map[string]string        `json:"limits,omitempty"`
	Requested                map[string]string        `json:"requested,omitempty"`
	ServiceAccountSecretName string                   `json:"serviceAccountSecretName,omitempty"`
}
