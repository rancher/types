package client

const (
	AzureFileVolumeSourceType            = "azureFileVolumeSource"
	AzureFileVolumeSourceFieldReadOnly   = "readOnly"
	AzureFileVolumeSourceFieldSecretName = "secretName"
	AzureFileVolumeSourceFieldShareName  = "shareName"
)

type AzureFileVolumeSource struct {
	ReadOnly   bool   `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	SecretName string `json:"secretName,omitempty" yaml:"secret_name,omitempty"`
	ShareName  string `json:"shareName,omitempty" yaml:"share_name,omitempty"`
}
