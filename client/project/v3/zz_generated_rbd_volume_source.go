package client

const (
	RBDVolumeSourceType              = "rbdVolumeSource"
	RBDVolumeSourceFieldCephMonitors = "monitors"
	RBDVolumeSourceFieldFSType       = "fsType"
	RBDVolumeSourceFieldKeyring      = "keyring"
	RBDVolumeSourceFieldRBDImage     = "image"
	RBDVolumeSourceFieldRBDPool      = "pool"
	RBDVolumeSourceFieldRadosUser    = "user"
	RBDVolumeSourceFieldReadOnly     = "readOnly"
	RBDVolumeSourceFieldSecretRef    = "secretRef"
)

type RBDVolumeSource struct {
	CephMonitors []string              `json:"monitors,omitempty" yaml:"monitors,omitempty"`
	FSType       string                `json:"fsType,omitempty" yaml:"fs_type,omitempty"`
	Keyring      string                `json:"keyring,omitempty" yaml:"keyring,omitempty"`
	RBDImage     string                `json:"image,omitempty" yaml:"image,omitempty"`
	RBDPool      string                `json:"pool,omitempty" yaml:"pool,omitempty"`
	RadosUser    string                `json:"user,omitempty" yaml:"user,omitempty"`
	ReadOnly     bool                  `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	SecretRef    *LocalObjectReference `json:"secretRef,omitempty" yaml:"secret_ref,omitempty"`
}
