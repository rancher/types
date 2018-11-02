package client

const (
	MultiClusterAppStatusType             = "multiClusterAppStatus"
	MultiClusterAppStatusFieldHealthState = "healthState"
)

type MultiClusterAppStatus struct {
	HealthState string `json:"healthState,omitempty" yaml:"healthState,omitempty"`
}
