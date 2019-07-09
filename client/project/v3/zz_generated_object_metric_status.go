package client

const (
	ObjectMetricStatusType                 = "objectMetricStatus"
	ObjectMetricStatusFieldCurrent         = "current"
	ObjectMetricStatusFieldDescribedObject = "describedObject"
	ObjectMetricStatusFieldMetric          = "metric"
)

type ObjectMetricStatus struct {
	Current         *MetricValueStatus           `json:"current,omitempty" yaml:"current,omitempty"`
	DescribedObject *CrossVersionObjectReference `json:"describedObject,omitempty" yaml:"described_object,omitempty"`
	Metric          *MetricIdentifier            `json:"metric,omitempty" yaml:"metric,omitempty"`
}
