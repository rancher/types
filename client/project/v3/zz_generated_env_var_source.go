package client

const (
	EnvVarSourceType                  = "envVarSource"
	EnvVarSourceFieldConfigMapKeyRef  = "configMapKeyRef"
	EnvVarSourceFieldFieldRef         = "fieldRef"
	EnvVarSourceFieldResourceFieldRef = "resourceFieldRef"
	EnvVarSourceFieldSecretKeyRef     = "secretKeyRef"
)

type EnvVarSource struct {
	ConfigMapKeyRef  *ConfigMapKeySelector  `json:"configMapKeyRef,omitempty" yaml:"config_map_key_ref,omitempty"`
	FieldRef         *ObjectFieldSelector   `json:"fieldRef,omitempty" yaml:"field_ref,omitempty"`
	ResourceFieldRef *ResourceFieldSelector `json:"resourceFieldRef,omitempty" yaml:"resource_field_ref,omitempty"`
	SecretKeyRef     *SecretKeySelector     `json:"secretKeyRef,omitempty" yaml:"secret_key_ref,omitempty"`
}
