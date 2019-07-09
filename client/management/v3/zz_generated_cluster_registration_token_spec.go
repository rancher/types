package client

const (
	ClusterRegistrationTokenSpecType           = "clusterRegistrationTokenSpec"
	ClusterRegistrationTokenSpecFieldClusterID = "clusterId"
)

type ClusterRegistrationTokenSpec struct {
	ClusterID string `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`
}
