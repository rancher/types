package client

const (
	ClusterSpecType                               = "clusterSpec"
	ClusterSpecFieldAzureKubernetesServiceConfig  = "azureKubernetesServiceConfig"
	ClusterSpecFieldDescription                   = "description"
	ClusterSpecFieldDisplayName                   = "displayName"
	ClusterSpecFieldGoogleKubernetesEngineConfig  = "googleKubernetesEngineConfig"
	ClusterSpecFieldRancherKubernetesEngineConfig = "rancherKubernetesEngineConfig"
)

type ClusterSpec struct {
	AzureKubernetesServiceConfig  *AzureKubernetesServiceConfig  `json:"azureKubernetesServiceConfig,omitempty"`
	Description                   string                         `json:"description,omitempty"`
	DisplayName                   string                         `json:"displayName,omitempty"`
	GoogleKubernetesEngineConfig  *GoogleKubernetesEngineConfig  `json:"googleKubernetesEngineConfig,omitempty"`
	RancherKubernetesEngineConfig *RancherKubernetesEngineConfig `json:"rancherKubernetesEngineConfig,omitempty"`
}
