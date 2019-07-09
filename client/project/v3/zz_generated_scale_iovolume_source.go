package client

const (
	ScaleIOVolumeSourceType                  = "scaleIOVolumeSource"
	ScaleIOVolumeSourceFieldFSType           = "fsType"
	ScaleIOVolumeSourceFieldGateway          = "gateway"
	ScaleIOVolumeSourceFieldProtectionDomain = "protectionDomain"
	ScaleIOVolumeSourceFieldReadOnly         = "readOnly"
	ScaleIOVolumeSourceFieldSSLEnabled       = "sslEnabled"
	ScaleIOVolumeSourceFieldSecretRef        = "secretRef"
	ScaleIOVolumeSourceFieldStorageMode      = "storageMode"
	ScaleIOVolumeSourceFieldStoragePool      = "storagePool"
	ScaleIOVolumeSourceFieldSystem           = "system"
	ScaleIOVolumeSourceFieldVolumeName       = "volumeName"
)

type ScaleIOVolumeSource struct {
	FSType           string                `json:"fsType,omitempty" yaml:"fs_type,omitempty"`
	Gateway          string                `json:"gateway,omitempty" yaml:"gateway,omitempty"`
	ProtectionDomain string                `json:"protectionDomain,omitempty" yaml:"protection_domain,omitempty"`
	ReadOnly         bool                  `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	SSLEnabled       bool                  `json:"sslEnabled,omitempty" yaml:"ssl_enabled,omitempty"`
	SecretRef        *LocalObjectReference `json:"secretRef,omitempty" yaml:"secret_ref,omitempty"`
	StorageMode      string                `json:"storageMode,omitempty" yaml:"storage_mode,omitempty"`
	StoragePool      string                `json:"storagePool,omitempty" yaml:"storage_pool,omitempty"`
	System           string                `json:"system,omitempty" yaml:"system,omitempty"`
	VolumeName       string                `json:"volumeName,omitempty" yaml:"volume_name,omitempty"`
}
