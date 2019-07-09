package client

const (
	MetricValueStatusType              = "metricValueStatus"
	MetricValueStatusFieldAverageValue = "averageValue"
	MetricValueStatusFieldUtilization  = "utilization"
	MetricValueStatusFieldValue        = "value"
)

type MetricValueStatus struct {
	AverageValue string `json:"averageValue,omitempty" yaml:"average_value,omitempty"`
	Utilization  *int64 `json:"utilization,omitempty" yaml:"utilization,omitempty"`
	Value        string `json:"value,omitempty" yaml:"value,omitempty"`
}
