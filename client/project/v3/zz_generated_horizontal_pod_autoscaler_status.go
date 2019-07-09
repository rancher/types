package client

const (
	HorizontalPodAutoscalerStatusType                    = "horizontalPodAutoscalerStatus"
	HorizontalPodAutoscalerStatusFieldConditions         = "conditions"
	HorizontalPodAutoscalerStatusFieldCurrentMetrics     = "currentMetrics"
	HorizontalPodAutoscalerStatusFieldCurrentReplicas    = "currentReplicas"
	HorizontalPodAutoscalerStatusFieldDesiredReplicas    = "desiredReplicas"
	HorizontalPodAutoscalerStatusFieldLastScaleTime      = "lastScaleTime"
	HorizontalPodAutoscalerStatusFieldObservedGeneration = "observedGeneration"
)

type HorizontalPodAutoscalerStatus struct {
	Conditions         []HorizontalPodAutoscalerCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	CurrentMetrics     []MetricStatus                     `json:"currentMetrics,omitempty" yaml:"current_metrics,omitempty"`
	CurrentReplicas    int64                              `json:"currentReplicas,omitempty" yaml:"current_replicas,omitempty"`
	DesiredReplicas    int64                              `json:"desiredReplicas,omitempty" yaml:"desired_replicas,omitempty"`
	LastScaleTime      string                             `json:"lastScaleTime,omitempty" yaml:"last_scale_time,omitempty"`
	ObservedGeneration *int64                             `json:"observedGeneration,omitempty" yaml:"observed_generation,omitempty"`
}
