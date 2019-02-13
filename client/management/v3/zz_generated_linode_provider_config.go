package client

const (
	LinodeProviderConfigType            = "linodeProviderConfig"
	LinodeProviderConfigFieldAccessKey  = "accessKey"
	LinodeProviderConfigFieldRootDomain = "rootDomain"
	LinodeProviderConfigFieldSecretKey  = "secretKey"
)

type LinodeProviderConfig struct {
	AccessKey  string `json:"accessKey,omitempty" yaml:"accessKey,omitempty"`
	RootDomain string `json:"rootDomain,omitempty" yaml:"rootDomain,omitempty"`
	SecretKey  string `json:"secretKey,omitempty" yaml:"secretKey,omitempty"`
}
