package client

const (
	NodeSystemInfoType                         = "nodeSystemInfo"
	NodeSystemInfoFieldArchitecture            = "architecture"
	NodeSystemInfoFieldBootID                  = "bootID"
	NodeSystemInfoFieldContainerRuntimeVersion = "containerRuntimeVersion"
	NodeSystemInfoFieldKernelVersion           = "kernelVersion"
	NodeSystemInfoFieldKubeProxyVersion        = "kubeProxyVersion"
	NodeSystemInfoFieldKubeletVersion          = "kubeletVersion"
	NodeSystemInfoFieldMachineID               = "machineID"
	NodeSystemInfoFieldOSImage                 = "osImage"
	NodeSystemInfoFieldOperatingSystem         = "operatingSystem"
	NodeSystemInfoFieldSystemUUID              = "systemUUID"
)

type NodeSystemInfo struct {
	Architecture            string `json:"architecture,omitempty" yaml:"architecture,omitempty"`
	BootID                  string `json:"bootID,omitempty" yaml:"boot_id,omitempty"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion,omitempty" yaml:"container_runtime_version,omitempty"`
	KernelVersion           string `json:"kernelVersion,omitempty" yaml:"kernel_version,omitempty"`
	KubeProxyVersion        string `json:"kubeProxyVersion,omitempty" yaml:"kube_proxy_version,omitempty"`
	KubeletVersion          string `json:"kubeletVersion,omitempty" yaml:"kubelet_version,omitempty"`
	MachineID               string `json:"machineID,omitempty" yaml:"machine_id,omitempty"`
	OSImage                 string `json:"osImage,omitempty" yaml:"os_image,omitempty"`
	OperatingSystem         string `json:"operatingSystem,omitempty" yaml:"operating_system,omitempty"`
	SystemUUID              string `json:"systemUUID,omitempty" yaml:"system_uuid,omitempty"`
}
