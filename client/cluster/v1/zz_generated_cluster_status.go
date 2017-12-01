package client

const (
	ClusterStatusType                     = "clusterStatus"
	ClusterStatusFieldAPIEndpoint         = "apiEndpoint"
	ClusterStatusFieldAllocatable         = "allocatable"
	ClusterStatusFieldAppliedSpec         = "appliedSpec"
	ClusterStatusFieldCACert              = "caCert"
	ClusterStatusFieldCapacity            = "capacity"
	ClusterStatusFieldComponentStatuses   = "componentStatuses"
	ClusterStatusFieldConditions          = "conditions"
	ClusterStatusFieldServiceAccountToken = "serviceAccountToken"
)

type ClusterStatus struct {
	APIEndpoint              string                   `json:"apiEndpoint,omitempty"`
	Allocatable              map[string]string        `json:"allocatable,omitempty"`
	AppliedSpec              *ClusterSpec             `json:"appliedSpec,omitempty"`
	CACert                   string                   `json:"caCert,omitempty"`
	Capacity                 map[string]string        `json:"capacity,omitempty"`
	ComponentStatuses        []ClusterComponentStatus `json:"componentStatuses,omitempty"`
	Conditions               []ClusterCondition       `json:"conditions,omitempty"`
	ServiceAccountSecretName string                   `json:"serviceAccountSecretName,omitempty"`
}
