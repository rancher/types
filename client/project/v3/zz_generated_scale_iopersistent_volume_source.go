package client

const (
	ScaleIOPersistentVolumeSourceType                  = "scaleIOPersistentVolumeSource"
	ScaleIOPersistentVolumeSourceFieldFSType           = "fsType"
	ScaleIOPersistentVolumeSourceFieldGateway          = "gateway"
	ScaleIOPersistentVolumeSourceFieldProtectionDomain = "protectionDomain"
	ScaleIOPersistentVolumeSourceFieldReadOnly         = "readOnly"
	ScaleIOPersistentVolumeSourceFieldSSLEnabled       = "sslEnabled"
	ScaleIOPersistentVolumeSourceFieldSecretRef        = "secretRef"
	ScaleIOPersistentVolumeSourceFieldStorageMode      = "storageMode"
	ScaleIOPersistentVolumeSourceFieldStoragePool      = "storagePool"
	ScaleIOPersistentVolumeSourceFieldSystem           = "system"
	ScaleIOPersistentVolumeSourceFieldVolumeName       = "volumeName"
)

type ScaleIOPersistentVolumeSource struct {
	FSType           string           `json:"fsType,omitempty" yaml:"fs_type,omitempty"`
	Gateway          string           `json:"gateway,omitempty" yaml:"gateway,omitempty"`
	ProtectionDomain string           `json:"protectionDomain,omitempty" yaml:"protection_domain,omitempty"`
	ReadOnly         bool             `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	SSLEnabled       bool             `json:"sslEnabled,omitempty" yaml:"ssl_enabled,omitempty"`
	SecretRef        *SecretReference `json:"secretRef,omitempty" yaml:"secret_ref,omitempty"`
	StorageMode      string           `json:"storageMode,omitempty" yaml:"storage_mode,omitempty"`
	StoragePool      string           `json:"storagePool,omitempty" yaml:"storage_pool,omitempty"`
	System           string           `json:"system,omitempty" yaml:"system,omitempty"`
	VolumeName       string           `json:"volumeName,omitempty" yaml:"volume_name,omitempty"`
}
