package client

const (
	KubernetesInfoType                  = "kubernetesInfo"
	KubernetesInfoFieldKubeProxyVersion = "kubeProxyVersion"
	KubernetesInfoFieldKubeletVersion   = "kubeletVersion"
)

type KubernetesInfo struct {
	KubeProxyVersion string `json:"kubeProxyVersion,omitempty" yaml:"kube_proxy_version,omitempty"`
	KubeletVersion   string `json:"kubeletVersion,omitempty" yaml:"kubelet_version,omitempty"`
}
