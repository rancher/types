package client

const (
	CinderVolumeSourceType           = "cinderVolumeSource"
	CinderVolumeSourceFieldFSType    = "fsType"
	CinderVolumeSourceFieldReadOnly  = "readOnly"
	CinderVolumeSourceFieldSecretRef = "secretRef"
	CinderVolumeSourceFieldVolumeID  = "volumeID"
)

type CinderVolumeSource struct {
	FSType    string                `json:"fsType,omitempty" yaml:"fs_type,omitempty"`
	ReadOnly  bool                  `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" yaml:"secret_ref,omitempty"`
	VolumeID  string                `json:"volumeID,omitempty" yaml:"volume_id,omitempty"`
}
