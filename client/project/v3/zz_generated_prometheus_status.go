package client

const (
	PrometheusStatusType                     = "prometheusStatus"
	PrometheusStatusFieldAvailableReplicas   = "availableReplicas"
	PrometheusStatusFieldPaused              = "paused"
	PrometheusStatusFieldReplicas            = "replicas"
	PrometheusStatusFieldUnavailableReplicas = "unavailableReplicas"
	PrometheusStatusFieldUpdatedReplicas     = "updatedReplicas"
)

type PrometheusStatus struct {
	AvailableReplicas   int64 `json:"availableReplicas,omitempty" yaml:"available_replicas,omitempty"`
	Paused              bool  `json:"paused,omitempty" yaml:"paused,omitempty"`
	Replicas            int64 `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	UnavailableReplicas int64 `json:"unavailableReplicas,omitempty" yaml:"unavailable_replicas,omitempty"`
	UpdatedReplicas     int64 `json:"updatedReplicas,omitempty" yaml:"updated_replicas,omitempty"`
}
