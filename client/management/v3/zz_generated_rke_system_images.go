package client

const (
	RKESystemImagesType                           = "rkeSystemImages"
	RKESystemImagesFieldAlpine                    = "alpine"
	RKESystemImagesFieldCalicoCNI                 = "calicoCni"
	RKESystemImagesFieldCalicoControllers         = "calicoControllers"
	RKESystemImagesFieldCalicoCtl                 = "calicoCtl"
	RKESystemImagesFieldCalicoNode                = "calicoNode"
	RKESystemImagesFieldCanalCNI                  = "canalCni"
	RKESystemImagesFieldCanalFlannel              = "canalFlannel"
	RKESystemImagesFieldCanalNode                 = "canalNode"
	RKESystemImagesFieldCertDownloader            = "certDownloader"
	RKESystemImagesFieldCoreDNS                   = "coredns"
	RKESystemImagesFieldCoreDNSAutoscaler         = "corednsAutoscaler"
	RKESystemImagesFieldDNSmasq                   = "dnsmasq"
	RKESystemImagesFieldEtcd                      = "etcd"
	RKESystemImagesFieldFlannel                   = "flannel"
	RKESystemImagesFieldFlannelCNI                = "flannelCni"
	RKESystemImagesFieldIngress                   = "ingress"
	RKESystemImagesFieldIngressBackend            = "ingressBackend"
	RKESystemImagesFieldKubeDNS                   = "kubedns"
	RKESystemImagesFieldKubeDNSAutoscaler         = "kubednsAutoscaler"
	RKESystemImagesFieldKubeDNSSidecar            = "kubednsSidecar"
	RKESystemImagesFieldKubernetes                = "kubernetes"
	RKESystemImagesFieldKubernetesServicesSidecar = "kubernetesServicesSidecar"
	RKESystemImagesFieldMetricsServer             = "metricsServer"
	RKESystemImagesFieldNginxProxy                = "nginxProxy"
	RKESystemImagesFieldPodInfraContainer         = "podInfraContainer"
	RKESystemImagesFieldWeaveCNI                  = "weaveCni"
	RKESystemImagesFieldWeaveNode                 = "weaveNode"
)

type RKESystemImages struct {
	Alpine                    string `json:"alpine,omitempty" yaml:"alpine,omitempty"`
	CalicoCNI                 string `json:"calicoCni,omitempty" yaml:"calico_cni,omitempty"`
	CalicoControllers         string `json:"calicoControllers,omitempty" yaml:"calico_controllers,omitempty"`
	CalicoCtl                 string `json:"calicoCtl,omitempty" yaml:"calico_ctl,omitempty"`
	CalicoNode                string `json:"calicoNode,omitempty" yaml:"calico_node,omitempty"`
	CanalCNI                  string `json:"canalCni,omitempty" yaml:"canal_cni,omitempty"`
	CanalFlannel              string `json:"canalFlannel,omitempty" yaml:"canal_flannel,omitempty"`
	CanalNode                 string `json:"canalNode,omitempty" yaml:"canal_node,omitempty"`
	CertDownloader            string `json:"certDownloader,omitempty" yaml:"cert_downloader,omitempty"`
	CoreDNS                   string `json:"coredns,omitempty" yaml:"coredns,omitempty"`
	CoreDNSAutoscaler         string `json:"corednsAutoscaler,omitempty" yaml:"coredns_autoscaler,omitempty"`
	DNSmasq                   string `json:"dnsmasq,omitempty" yaml:"dnsmasq,omitempty"`
	Etcd                      string `json:"etcd,omitempty" yaml:"etcd,omitempty"`
	Flannel                   string `json:"flannel,omitempty" yaml:"flannel,omitempty"`
	FlannelCNI                string `json:"flannelCni,omitempty" yaml:"flannel_cni,omitempty"`
	Ingress                   string `json:"ingress,omitempty" yaml:"ingress,omitempty"`
	IngressBackend            string `json:"ingressBackend,omitempty" yaml:"ingress_backend,omitempty"`
	KubeDNS                   string `json:"kubedns,omitempty" yaml:"kubedns,omitempty"`
	KubeDNSAutoscaler         string `json:"kubednsAutoscaler,omitempty" yaml:"kubedns_autoscaler,omitempty"`
	KubeDNSSidecar            string `json:"kubednsSidecar,omitempty" yaml:"kubedns_sidecar,omitempty"`
	Kubernetes                string `json:"kubernetes,omitempty" yaml:"kubernetes,omitempty"`
	KubernetesServicesSidecar string `json:"kubernetesServicesSidecar,omitempty" yaml:"kubernetes_services_sidecar,omitempty"`
	MetricsServer             string `json:"metricsServer,omitempty" yaml:"metrics_server,omitempty"`
	NginxProxy                string `json:"nginxProxy,omitempty" yaml:"nginx_proxy,omitempty"`
	PodInfraContainer         string `json:"podInfraContainer,omitempty" yaml:"pod_infra_container,omitempty"`
	WeaveCNI                  string `json:"weaveCni,omitempty" yaml:"weave_cni,omitempty"`
	WeaveNode                 string `json:"weaveNode,omitempty" yaml:"weave_node,omitempty"`
}
