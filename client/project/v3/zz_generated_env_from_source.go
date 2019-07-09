package client

const (
	EnvFromSourceType              = "envFromSource"
	EnvFromSourceFieldConfigMapRef = "configMapRef"
	EnvFromSourceFieldPrefix       = "prefix"
	EnvFromSourceFieldSecretRef    = "secretRef"
)

type EnvFromSource struct {
	ConfigMapRef *ConfigMapEnvSource `json:"configMapRef,omitempty" yaml:"config_map_ref,omitempty"`
	Prefix       string              `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	SecretRef    *SecretEnvSource    `json:"secretRef,omitempty" yaml:"secret_ref,omitempty"`
}
