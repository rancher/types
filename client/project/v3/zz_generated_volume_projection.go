package client

const (
	VolumeProjectionType                     = "volumeProjection"
	VolumeProjectionFieldConfigMap           = "configMap"
	VolumeProjectionFieldDownwardAPI         = "downwardAPI"
	VolumeProjectionFieldSecret              = "secret"
	VolumeProjectionFieldServiceAccountToken = "serviceAccountToken"
)

type VolumeProjection struct {
	ConfigMap           *ConfigMapProjection           `json:"configMap,omitempty" yaml:"config_map,omitempty"`
	DownwardAPI         *DownwardAPIProjection         `json:"downwardAPI,omitempty" yaml:"downward_api,omitempty"`
	Secret              *SecretProjection              `json:"secret,omitempty" yaml:"secret,omitempty"`
	ServiceAccountToken *ServiceAccountTokenProjection `json:"serviceAccountToken,omitempty" yaml:"service_account_token,omitempty"`
}
