package client

const (
	ServiceMonitorSpecType                   = "serviceMonitorSpec"
	ServiceMonitorSpecFieldEndpoints         = "endpoints"
	ServiceMonitorSpecFieldJobLabel          = "jobLabel"
	ServiceMonitorSpecFieldNamespaceSelector = "namespaceSelector"
	ServiceMonitorSpecFieldSelector          = "selector"
	ServiceMonitorSpecFieldTargetLabels      = "targetLabels"
)

type ServiceMonitorSpec struct {
	Endpoints         []Endpoint     `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
	JobLabel          string         `json:"jobLabel,omitempty" yaml:"jobLabel,omitempty"`
	NamespaceSelector []string       `json:"namespaceSelector,omitempty" yaml:"namespaceSelector,omitempty"`
	Selector          *LabelSelector `json:"selector,omitempty" yaml:"selector,omitempty"`
	TargetLabels      []string       `json:"targetLabels,omitempty" yaml:"targetLabels,omitempty"`
}
