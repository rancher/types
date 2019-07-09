package client

const (
	StorageOSPersistentVolumeSourceType                 = "storageOSPersistentVolumeSource"
	StorageOSPersistentVolumeSourceFieldFSType          = "fsType"
	StorageOSPersistentVolumeSourceFieldReadOnly        = "readOnly"
	StorageOSPersistentVolumeSourceFieldSecretRef       = "secretRef"
	StorageOSPersistentVolumeSourceFieldVolumeName      = "volumeName"
	StorageOSPersistentVolumeSourceFieldVolumeNamespace = "volumeNamespace"
)

type StorageOSPersistentVolumeSource struct {
	FSType          string           `json:"fsType,omitempty" yaml:"fs_type,omitempty"`
	ReadOnly        bool             `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	SecretRef       *ObjectReference `json:"secretRef,omitempty" yaml:"secret_ref,omitempty"`
	VolumeName      string           `json:"volumeName,omitempty" yaml:"volume_name,omitempty"`
	VolumeNamespace string           `json:"volumeNamespace,omitempty" yaml:"volume_namespace,omitempty"`
}
