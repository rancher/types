package client

const (
	CephFSPersistentVolumeSourceType            = "cephFSPersistentVolumeSource"
	CephFSPersistentVolumeSourceFieldMonitors   = "monitors"
	CephFSPersistentVolumeSourceFieldPath       = "path"
	CephFSPersistentVolumeSourceFieldReadOnly   = "readOnly"
	CephFSPersistentVolumeSourceFieldSecretFile = "secretFile"
	CephFSPersistentVolumeSourceFieldSecretRef  = "secretRef"
	CephFSPersistentVolumeSourceFieldUser       = "user"
)

type CephFSPersistentVolumeSource struct {
	Monitors   []string         `json:"monitors,omitempty" yaml:"monitors,omitempty"`
	Path       string           `json:"path,omitempty" yaml:"path,omitempty"`
	ReadOnly   bool             `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	SecretFile string           `json:"secretFile,omitempty" yaml:"secret_file,omitempty"`
	SecretRef  *SecretReference `json:"secretRef,omitempty" yaml:"secret_ref,omitempty"`
	User       string           `json:"user,omitempty" yaml:"user,omitempty"`
}
