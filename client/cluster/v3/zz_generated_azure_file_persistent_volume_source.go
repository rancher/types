package client

const (
	AzureFilePersistentVolumeSourceType                 = "azureFilePersistentVolumeSource"
	AzureFilePersistentVolumeSourceFieldReadOnly        = "readOnly"
	AzureFilePersistentVolumeSourceFieldSecretName      = "secretName"
	AzureFilePersistentVolumeSourceFieldSecretNamespace = "secretNamespace"
	AzureFilePersistentVolumeSourceFieldShareName       = "shareName"
)

type AzureFilePersistentVolumeSource struct {
	ReadOnly        bool   `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	SecretName      string `json:"secretName,omitempty" yaml:"secret_name,omitempty"`
	SecretNamespace string `json:"secretNamespace,omitempty" yaml:"secret_namespace,omitempty"`
	ShareName       string `json:"shareName,omitempty" yaml:"share_name,omitempty"`
}
