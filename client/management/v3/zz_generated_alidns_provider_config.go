package client

const (
	AlidnsProviderConfigType           = "alidnsProviderConfig"
	AlidnsProviderConfigFieldAccessKey = "accessKey"
	AlidnsProviderConfigFieldSecretKey = "secretKey"
)

type AlidnsProviderConfig struct {
	AccessKey string `json:"accessKey,omitempty" yaml:"access_key,omitempty"`
	SecretKey string `json:"secretKey,omitempty" yaml:"secret_key,omitempty"`
}
