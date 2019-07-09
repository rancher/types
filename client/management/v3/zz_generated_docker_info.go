package client

const (
	DockerInfoType                    = "dockerInfo"
	DockerInfoFieldArchitecture       = "architecture"
	DockerInfoFieldCgroupDriver       = "cgroupDriver"
	DockerInfoFieldDebug              = "debug"
	DockerInfoFieldDockerRootDir      = "dockerRootDir"
	DockerInfoFieldDriver             = "driver"
	DockerInfoFieldExperimentalBuild  = "experimentalBuild"
	DockerInfoFieldHTTPProxy          = "httpProxy"
	DockerInfoFieldHTTPSProxy         = "httpsProxy"
	DockerInfoFieldIndexServerAddress = "indexServerAddress"
	DockerInfoFieldKernelVersion      = "kernelVersion"
	DockerInfoFieldLabels             = "labels"
	DockerInfoFieldLoggingDriver      = "loggingDriver"
	DockerInfoFieldName               = "name"
	DockerInfoFieldNoProxy            = "noProxy"
	DockerInfoFieldOSType             = "osType"
	DockerInfoFieldOperatingSystem    = "operatingSystem"
	DockerInfoFieldServerVersion      = "serverVersion"
)

type DockerInfo struct {
	Architecture       string   `json:"architecture,omitempty" yaml:"architecture,omitempty"`
	CgroupDriver       string   `json:"cgroupDriver,omitempty" yaml:"cgroup_driver,omitempty"`
	Debug              bool     `json:"debug,omitempty" yaml:"debug,omitempty"`
	DockerRootDir      string   `json:"dockerRootDir,omitempty" yaml:"docker_root_dir,omitempty"`
	Driver             string   `json:"driver,omitempty" yaml:"driver,omitempty"`
	ExperimentalBuild  bool     `json:"experimentalBuild,omitempty" yaml:"experimental_build,omitempty"`
	HTTPProxy          string   `json:"httpProxy,omitempty" yaml:"http_proxy,omitempty"`
	HTTPSProxy         string   `json:"httpsProxy,omitempty" yaml:"https_proxy,omitempty"`
	IndexServerAddress string   `json:"indexServerAddress,omitempty" yaml:"index_server_address,omitempty"`
	KernelVersion      string   `json:"kernelVersion,omitempty" yaml:"kernel_version,omitempty"`
	Labels             []string `json:"labels,omitempty" yaml:"labels,omitempty"`
	LoggingDriver      string   `json:"loggingDriver,omitempty" yaml:"logging_driver,omitempty"`
	Name               string   `json:"name,omitempty" yaml:"name,omitempty"`
	NoProxy            string   `json:"noProxy,omitempty" yaml:"no_proxy,omitempty"`
	OSType             string   `json:"osType,omitempty" yaml:"os_type,omitempty"`
	OperatingSystem    string   `json:"operatingSystem,omitempty" yaml:"operating_system,omitempty"`
	ServerVersion      string   `json:"serverVersion,omitempty" yaml:"server_version,omitempty"`
}
