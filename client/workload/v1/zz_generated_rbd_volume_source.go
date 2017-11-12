package client

const (
	RBDVolumeSourceType              = "rbdVolumeSource"
	RBDVolumeSourceFieldCephMonitors = "cephMonitors"
	RBDVolumeSourceFieldFSType       = "fsType"
	RBDVolumeSourceFieldKeyring      = "keyring"
	RBDVolumeSourceFieldRBDImage     = "rbdImage"
	RBDVolumeSourceFieldRBDPool      = "rbdPool"
	RBDVolumeSourceFieldRadosUser    = "radosUser"
	RBDVolumeSourceFieldReadOnly     = "readOnly"
	RBDVolumeSourceFieldSecretRef    = "secretRef"
)

type RBDVolumeSource struct {
	CephMonitors []string              `json:"cephMonitors,omitempty"`
	FSType       string                `json:"fsType,omitempty"`
	Keyring      string                `json:"keyring,omitempty"`
	RBDImage     string                `json:"rbdImage,omitempty"`
	RBDPool      string                `json:"rbdPool,omitempty"`
	RadosUser    string                `json:"radosUser,omitempty"`
	ReadOnly     bool                  `json:"readOnly,omitempty"`
	SecretRef    *LocalObjectReference `json:"secretRef,omitempty"`
}
