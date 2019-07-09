package client

const (
	CinderPersistentVolumeSourceType           = "cinderPersistentVolumeSource"
	CinderPersistentVolumeSourceFieldFSType    = "fsType"
	CinderPersistentVolumeSourceFieldReadOnly  = "readOnly"
	CinderPersistentVolumeSourceFieldSecretRef = "secretRef"
	CinderPersistentVolumeSourceFieldVolumeID  = "volumeID"
)

type CinderPersistentVolumeSource struct {
	FSType    string           `json:"fsType,omitempty" yaml:"fs_type,omitempty"`
	ReadOnly  bool             `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	SecretRef *SecretReference `json:"secretRef,omitempty" yaml:"secret_ref,omitempty"`
	VolumeID  string           `json:"volumeID,omitempty" yaml:"volume_id,omitempty"`
}
