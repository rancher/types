package client

const (
	WindowsSystemImagesType                    = "windowsSystemImages"
	WindowsSystemImagesFieldCalicoCNIBinaries  = "calicoCniBinaries"
	WindowsSystemImagesFieldCanalCNIBinaries   = "canalCniBinaries"
	WindowsSystemImagesFieldFlannelCNIBinaries = "flannelCniBinaries"
	WindowsSystemImagesFieldKubeletPause       = "kubeletPause"
	WindowsSystemImagesFieldKubernetesBinaries = "kubernetesBinaries"
	WindowsSystemImagesFieldNginxProxy         = "nginxProxy"
)

type WindowsSystemImages struct {
	CalicoCNIBinaries  string `json:"calicoCniBinaries,omitempty" yaml:"calico_cni_binaries,omitempty"`
	CanalCNIBinaries   string `json:"canalCniBinaries,omitempty" yaml:"canal_cni_binaries,omitempty"`
	FlannelCNIBinaries string `json:"flannelCniBinaries,omitempty" yaml:"flannel_cni_binaries,omitempty"`
	KubeletPause       string `json:"kubeletPause,omitempty" yaml:"kubelet_pause,omitempty"`
	KubernetesBinaries string `json:"kubernetesBinaries,omitempty" yaml:"kubernetes_binaries,omitempty"`
	NginxProxy         string `json:"nginxProxy,omitempty" yaml:"nginx_proxy,omitempty"`
}
