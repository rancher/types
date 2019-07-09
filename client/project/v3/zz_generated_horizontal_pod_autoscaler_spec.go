package client

const (
	HorizontalPodAutoscalerSpecType                = "horizontalPodAutoscalerSpec"
	HorizontalPodAutoscalerSpecFieldMaxReplicas    = "maxReplicas"
	HorizontalPodAutoscalerSpecFieldMetrics        = "metrics"
	HorizontalPodAutoscalerSpecFieldMinReplicas    = "minReplicas"
	HorizontalPodAutoscalerSpecFieldScaleTargetRef = "scaleTargetRef"
)

type HorizontalPodAutoscalerSpec struct {
	MaxReplicas    int64                        `json:"maxReplicas,omitempty" yaml:"max_replicas,omitempty"`
	Metrics        []Metric                     `json:"metrics,omitempty" yaml:"metrics,omitempty"`
	MinReplicas    *int64                       `json:"minReplicas,omitempty" yaml:"min_replicas,omitempty"`
	ScaleTargetRef *CrossVersionObjectReference `json:"scaleTargetRef,omitempty" yaml:"scale_target_ref,omitempty"`
}
