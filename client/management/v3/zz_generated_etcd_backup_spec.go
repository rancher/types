package client

const (
	EtcdBackupSpecType              = "etcdBackupSpec"
	EtcdBackupSpecFieldBackupConfig = "backupConfig"
	EtcdBackupSpecFieldClusterID    = "clusterId"
	EtcdBackupSpecFieldFilename     = "filename"
	EtcdBackupSpecFieldManual       = "manual"
)

type EtcdBackupSpec struct {
	BackupConfig *BackupConfig `json:"backupConfig,omitempty" yaml:"backup_config,omitempty"`
	ClusterID    string        `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`
	Filename     string        `json:"filename,omitempty" yaml:"filename,omitempty"`
	Manual       bool          `json:"manual,omitempty" yaml:"manual,omitempty"`
}
