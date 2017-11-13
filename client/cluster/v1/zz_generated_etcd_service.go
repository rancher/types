package client

const (
	ETCDServiceType           = "etcdService"
	ETCDServiceFieldExtraArgs = "extraArgs"
	ETCDServiceFieldImage     = "image"
)

type ETCDService struct {
	ExtraArgs []string `json:"extraArgs,omitempty"`
	Image     string   `json:"image,omitempty"`
}
