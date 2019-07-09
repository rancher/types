package client

const (
	PortworxVolumeSourceType          = "portworxVolumeSource"
	PortworxVolumeSourceFieldFSType   = "fsType"
	PortworxVolumeSourceFieldReadOnly = "readOnly"
	PortworxVolumeSourceFieldVolumeID = "volumeID"
)

type PortworxVolumeSource struct {
	FSType   string `json:"fsType,omitempty" yaml:"fs_type,omitempty"`
	ReadOnly bool   `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	VolumeID string `json:"volumeID,omitempty" yaml:"volume_id,omitempty"`
}
