package client

const (
	KubeproxyServiceType           = "kubeproxyService"
	KubeproxyServiceFieldExtraArgs = "extraArgs"
	KubeproxyServiceFieldImage     = "image"
)

type KubeproxyService struct {
	ExtraArgs []string `json:"extraArgs,omitempty"`
	Image     string   `json:"image,omitempty"`
}
