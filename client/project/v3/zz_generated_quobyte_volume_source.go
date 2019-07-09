package client

const (
	QuobyteVolumeSourceType          = "quobyteVolumeSource"
	QuobyteVolumeSourceFieldGroup    = "group"
	QuobyteVolumeSourceFieldReadOnly = "readOnly"
	QuobyteVolumeSourceFieldRegistry = "registry"
	QuobyteVolumeSourceFieldUser     = "user"
	QuobyteVolumeSourceFieldVolume   = "volume"
)

type QuobyteVolumeSource struct {
	Group    string `json:"group,omitempty" yaml:"group,omitempty"`
	ReadOnly bool   `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	Registry string `json:"registry,omitempty" yaml:"registry,omitempty"`
	User     string `json:"user,omitempty" yaml:"user,omitempty"`
	Volume   string `json:"volume,omitempty" yaml:"volume,omitempty"`
}
