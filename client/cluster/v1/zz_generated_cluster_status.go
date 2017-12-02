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
	ClusterStatusFieldLimits              = "limits"
	ClusterStatusFieldRequested           = "requested"
	ClusterStatusFieldServiceAccountToken = "serviceAccountToken"
)

type ClusterStatus struct {
	APIEndpoint         string                   `json:"apiEndpoint,omitempty"`
	Allocatable         map[string]string        `json:"allocatable,omitempty"`
	AppliedSpec         *ClusterSpec             `json:"appliedSpec,omitempty"`
	CACert              string                   `json:"caCert,omitempty"`
	Capacity            map[string]string        `json:"capacity,omitempty"`
	ComponentStatuses   []ClusterComponentStatus `json:"componentStatuses,omitempty"`
	Conditions          []ClusterCondition       `json:"conditions,omitempty"`
	Limits              map[string]string        `json:"limits,omitempty"`
	Requested           map[string]string        `json:"requested,omitempty"`
	ServiceAccountToken string                   `json:"serviceAccountToken,omitempty"`
}
