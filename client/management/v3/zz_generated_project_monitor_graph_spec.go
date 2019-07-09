package client

const (
	ProjectMonitorGraphSpecType                        = "projectMonitorGraphSpec"
	ProjectMonitorGraphSpecFieldDescription            = "description"
	ProjectMonitorGraphSpecFieldDetailsMetricsSelector = "detailsMetricsSelector"
	ProjectMonitorGraphSpecFieldDisplayResourceType    = "displayResourceType"
	ProjectMonitorGraphSpecFieldGraphType              = "graphType"
	ProjectMonitorGraphSpecFieldMetricsSelector        = "metricsSelector"
	ProjectMonitorGraphSpecFieldPriority               = "priority"
	ProjectMonitorGraphSpecFieldProjectID              = "projectId"
	ProjectMonitorGraphSpecFieldResourceType           = "resourceType"
	ProjectMonitorGraphSpecFieldYAxis                  = "yAxis"
)

type ProjectMonitorGraphSpec struct {
	Description            string            `json:"description,omitempty" yaml:"description,omitempty"`
	DetailsMetricsSelector map[string]string `json:"detailsMetricsSelector,omitempty" yaml:"details_metrics_selector,omitempty"`
	DisplayResourceType    string            `json:"displayResourceType,omitempty" yaml:"display_resource_type,omitempty"`
	GraphType              string            `json:"graphType,omitempty" yaml:"graph_type,omitempty"`
	MetricsSelector        map[string]string `json:"metricsSelector,omitempty" yaml:"metrics_selector,omitempty"`
	Priority               int64             `json:"priority,omitempty" yaml:"priority,omitempty"`
	ProjectID              string            `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	ResourceType           string            `json:"resourceType,omitempty" yaml:"resource_type,omitempty"`
	YAxis                  *YAxis            `json:"yAxis,omitempty" yaml:"y_axis,omitempty"`
}
