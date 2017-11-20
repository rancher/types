package client

const (
	ProjectSpecType             = "projectSpec"
	ProjectSpecFieldClusterName = "clusterName"
	ProjectSpecFieldDisplayName = "displayName"
)

type ProjectSpec struct {
	ClusterName string `json:"clusterName,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}
