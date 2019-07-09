package client

const (
	ResourceFieldSelectorType               = "resourceFieldSelector"
	ResourceFieldSelectorFieldContainerName = "containerName"
	ResourceFieldSelectorFieldDivisor       = "divisor"
	ResourceFieldSelectorFieldResource      = "resource"
)

type ResourceFieldSelector struct {
	ContainerName string `json:"containerName,omitempty" yaml:"container_name,omitempty"`
	Divisor       string `json:"divisor,omitempty" yaml:"divisor,omitempty"`
	Resource      string `json:"resource,omitempty" yaml:"resource,omitempty"`
}
