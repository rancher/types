package client

const (
	KubeControllerServiceType                       = "kubeControllerService"
	KubeControllerServiceFieldClusterCIDR           = "clusterCIDR"
	KubeControllerServiceFieldExtraArgs             = "extraArgs"
	KubeControllerServiceFieldImage                 = "image"
	KubeControllerServiceFieldServiceClusterIPRange = "serviceClusterIPRange"
)

type KubeControllerService struct {
	ClusterCIDR           string   `json:"clusterCIDR,omitempty"`
	ExtraArgs             []string `json:"extraArgs,omitempty"`
	Image                 string   `json:"image,omitempty"`
	ServiceClusterIPRange string   `json:"serviceClusterIPRange,omitempty"`
}
