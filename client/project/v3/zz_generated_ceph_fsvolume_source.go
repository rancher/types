package client

const (
	CephFSVolumeSourceType            = "cephFSVolumeSource"
	CephFSVolumeSourceFieldMonitors   = "monitors"
	CephFSVolumeSourceFieldPath       = "path"
	CephFSVolumeSourceFieldReadOnly   = "readOnly"
	CephFSVolumeSourceFieldSecretFile = "secretFile"
	CephFSVolumeSourceFieldSecretRef  = "secretRef"
	CephFSVolumeSourceFieldUser       = "user"
)

type CephFSVolumeSource struct {
	Monitors   []string              `json:"monitors,omitempty" yaml:"monitors,omitempty"`
	Path       string                `json:"path,omitempty" yaml:"path,omitempty"`
	ReadOnly   bool                  `json:"readOnly,omitempty" yaml:"read_only,omitempty"`
	SecretFile string                `json:"secretFile,omitempty" yaml:"secret_file,omitempty"`
	SecretRef  *LocalObjectReference `json:"secretRef,omitempty" yaml:"secret_ref,omitempty"`
	User       string                `json:"user,omitempty" yaml:"user,omitempty"`
}
