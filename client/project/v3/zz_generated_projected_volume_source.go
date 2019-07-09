package client

const (
	ProjectedVolumeSourceType             = "projectedVolumeSource"
	ProjectedVolumeSourceFieldDefaultMode = "defaultMode"
	ProjectedVolumeSourceFieldSources     = "sources"
)

type ProjectedVolumeSource struct {
	DefaultMode *int64             `json:"defaultMode,omitempty" yaml:"default_mode,omitempty"`
	Sources     []VolumeProjection `json:"sources,omitempty" yaml:"sources,omitempty"`
}
