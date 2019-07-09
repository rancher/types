package client

const (
	ClusterScanSpecType            = "clusterScanSpec"
	ClusterScanSpecFieldClusterID  = "clusterId"
	ClusterScanSpecFieldManual     = "manual"
	ClusterScanSpecFieldScanConfig = "scanConfig"
	ClusterScanSpecFieldScanType   = "scanType"
)

type ClusterScanSpec struct {
	ClusterID  string             `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`
	Manual     bool               `json:"manual,omitempty" yaml:"manual,omitempty"`
	ScanConfig *ClusterScanConfig `json:"scanConfig,omitempty" yaml:"scan_config,omitempty"`
	ScanType   string             `json:"scanType,omitempty" yaml:"scan_type,omitempty"`
}
