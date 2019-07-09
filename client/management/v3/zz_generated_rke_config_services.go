package client

const (
	RKEConfigServicesType                = "rkeConfigServices"
	RKEConfigServicesFieldEtcd           = "etcd"
	RKEConfigServicesFieldKubeAPI        = "kubeApi"
	RKEConfigServicesFieldKubeController = "kubeController"
	RKEConfigServicesFieldKubelet        = "kubelet"
	RKEConfigServicesFieldKubeproxy      = "kubeproxy"
	RKEConfigServicesFieldScheduler      = "scheduler"
)

type RKEConfigServices struct {
	Etcd           *ETCDService           `json:"etcd,omitempty" yaml:"etcd,omitempty"`
	KubeAPI        *KubeAPIService        `json:"kubeApi,omitempty" yaml:"kube_api,omitempty"`
	KubeController *KubeControllerService `json:"kubeController,omitempty" yaml:"kube_controller,omitempty"`
	Kubelet        *KubeletService        `json:"kubelet,omitempty" yaml:"kubelet,omitempty"`
	Kubeproxy      *KubeproxyService      `json:"kubeproxy,omitempty" yaml:"kubeproxy,omitempty"`
	Scheduler      *SchedulerService      `json:"scheduler,omitempty" yaml:"scheduler,omitempty"`
}
