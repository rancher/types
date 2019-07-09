package client

const (
	ConfigMapVolumeSourceType             = "configMapVolumeSource"
	ConfigMapVolumeSourceFieldDefaultMode = "defaultMode"
	ConfigMapVolumeSourceFieldItems       = "items"
	ConfigMapVolumeSourceFieldName        = "name"
	ConfigMapVolumeSourceFieldOptional    = "optional"
)

type ConfigMapVolumeSource struct {
	DefaultMode *int64      `json:"defaultMode,omitempty" yaml:"default_mode,omitempty"`
	Items       []KeyToPath `json:"items,omitempty" yaml:"items,omitempty"`
	Name        string      `json:"name,omitempty" yaml:"name,omitempty"`
	Optional    *bool       `json:"optional,omitempty" yaml:"optional,omitempty"`
}
