package client

const (
	VolumeMountType                  = "volumeMount"
	VolumeMountFieldMountPath        = "mountPath"
	VolumeMountFieldMountPropagation = "mountPropagation"
	VolumeMountFieldName             = "name"
	VolumeMountFieldReadOnly         = "readOnly"
	VolumeMountFieldSubPath          = "subPath"
)

type VolumeMount struct {
	MountPath        string `json:"mountPath,omitempty" yaml:"mount_path,omitempty"`
	MountPropagation string `json:"mountPropagation,omitempty" yaml:"mount_propagation,omitempty"`
	Name             string `json:"name,omitempty" yaml:"name,omitempty"`
	ReadOnly         bool   `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	SubPath          string `json:"subPath,omitempty" yaml:"sub_path,omitempty"`
}
