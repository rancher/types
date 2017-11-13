package client

const (
	KubeAPIServiceType                       = "kubeAPIService"
	KubeAPIServiceFieldExtraArgs             = "extraArgs"
	KubeAPIServiceFieldImage                 = "image"
	KubeAPIServiceFieldServiceClusterIPRange = "serviceClusterIPRange"
)

type KubeAPIService struct {
	ExtraArgs             []string `json:"extraArgs,omitempty"`
	Image                 string   `json:"image,omitempty"`
	ServiceClusterIPRange string   `json:"serviceClusterIPRange,omitempty"`
}
