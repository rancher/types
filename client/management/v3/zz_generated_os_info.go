package client

const (
	OSInfoType                 = "osInfo"
	OSInfoFieldDockerVersion   = "dockerVersion"
	OSInfoFieldKernelVersion   = "kernelVersion"
	OSInfoFieldOperatingSystem = "operatingSystem"
)

type OSInfo struct {
	DockerVersion   string `json:"dockerVersion,omitempty" yaml:"docker_version,omitempty"`
	KernelVersion   string `json:"kernelVersion,omitempty" yaml:"kernel_version,omitempty"`
	OperatingSystem string `json:"operatingSystem,omitempty" yaml:"operating_system,omitempty"`
}
