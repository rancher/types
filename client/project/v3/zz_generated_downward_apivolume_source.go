package client

const (
	DownwardAPIVolumeSourceType             = "downwardAPIVolumeSource"
	DownwardAPIVolumeSourceFieldDefaultMode = "defaultMode"
	DownwardAPIVolumeSourceFieldItems       = "items"
)

type DownwardAPIVolumeSource struct {
	DefaultMode *int64                  `json:"defaultMode,omitempty" yaml:"default_mode,omitempty"`
	Items       []DownwardAPIVolumeFile `json:"items,omitempty" yaml:"items,omitempty"`
}
