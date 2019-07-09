package client

const (
	ProjectSpecType                               = "projectSpec"
	ProjectSpecFieldClusterID                     = "clusterId"
	ProjectSpecFieldContainerDefaultResourceLimit = "containerDefaultResourceLimit"
	ProjectSpecFieldDescription                   = "description"
	ProjectSpecFieldDisplayName                   = "displayName"
	ProjectSpecFieldEnableProjectMonitoring       = "enableProjectMonitoring"
	ProjectSpecFieldNamespaceDefaultResourceQuota = "namespaceDefaultResourceQuota"
	ProjectSpecFieldResourceQuota                 = "resourceQuota"
)

type ProjectSpec struct {
	ClusterID                     string                  `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`
	ContainerDefaultResourceLimit *ContainerResourceLimit `json:"containerDefaultResourceLimit,omitempty" yaml:"container_default_resource_limit,omitempty"`
	Description                   string                  `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName                   string                  `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	EnableProjectMonitoring       bool                    `json:"enableProjectMonitoring,omitempty" yaml:"enable_project_monitoring,omitempty"`
	NamespaceDefaultResourceQuota *NamespaceResourceQuota `json:"namespaceDefaultResourceQuota,omitempty" yaml:"namespace_default_resource_quota,omitempty"`
	ResourceQuota                 *ProjectResourceQuota   `json:"resourceQuota,omitempty" yaml:"resource_quota,omitempty"`
}
