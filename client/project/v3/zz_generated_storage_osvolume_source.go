package client

const (
	StorageOSVolumeSourceType                 = "storageOSVolumeSource"
	StorageOSVolumeSourceFieldFSType          = "fsType"
	StorageOSVolumeSourceFieldReadOnly        = "readOnly"
	StorageOSVolumeSourceFieldSecretRef       = "secretRef"
	StorageOSVolumeSourceFieldVolumeName      = "volumeName"
	StorageOSVolumeSourceFieldVolumeNamespace = "volumeNamespace"
)

type StorageOSVolumeSource struct {
	FSType          string                `json:"fsType,omitempty" yaml:"fs_type,omitempty"`
	ReadOnly        bool                  `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	SecretRef       *LocalObjectReference `json:"secretRef,omitempty" yaml:"secret_ref,omitempty"`
	VolumeName      string                `json:"volumeName,omitempty" yaml:"volume_name,omitempty"`
	VolumeNamespace string                `json:"volumeNamespace,omitempty" yaml:"volume_namespace,omitempty"`
}
