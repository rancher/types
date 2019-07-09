package client

const (
	ServiceMonitorSpecType                   = "serviceMonitorSpec"
	ServiceMonitorSpecFieldEndpoints         = "endpoints"
	ServiceMonitorSpecFieldJobLabel          = "jobLabel"
	ServiceMonitorSpecFieldNamespaceSelector = "namespaceSelector"
	ServiceMonitorSpecFieldPodTargetLabels   = "podTargetLabels"
	ServiceMonitorSpecFieldSampleLimit       = "sampleLimit"
	ServiceMonitorSpecFieldSelector          = "selector"
	ServiceMonitorSpecFieldTargetLabels      = "targetLabels"
)

type ServiceMonitorSpec struct {
	Endpoints         []Endpoint     `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
	JobLabel          string         `json:"jobLabel,omitempty" yaml:"job_label,omitempty"`
	NamespaceSelector []string       `json:"namespaceSelector,omitempty" yaml:"namespace_selector,omitempty"`
	PodTargetLabels   []string       `json:"podTargetLabels,omitempty" yaml:"pod_target_labels,omitempty"`
	SampleLimit       int64          `json:"sampleLimit,omitempty" yaml:"sample_limit,omitempty"`
	Selector          *LabelSelector `json:"selector,omitempty" yaml:"selector,omitempty"`
	TargetLabels      []string       `json:"targetLabels,omitempty" yaml:"target_labels,omitempty"`
}
