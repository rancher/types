package client

const (
	TargetStatusType                  = "targetStatus"
	TargetStatusFieldChartReleaseName = "chartReleaseName"
	TargetStatusFieldHealthState      = "healthState"
)

type TargetStatus struct {
	ChartReleaseName string `json:"chartReleaseName,omitempty" yaml:"chartReleaseName,omitempty"`
	HealthState      string `json:"healthState,omitempty" yaml:"healthState,omitempty"`
}
