package client

const (
	ClusterSpecType                               = "clusterSpec"
	ClusterSpecFieldAzureKubernetesServiceConfig  = "azureKubernetesServiceConfig"
	ClusterSpecFieldGoogleKubernetesEngineConfig  = "googleKubernetesEngineConfig"
	ClusterSpecFieldRancherKubernetesEngineConfig = "rancherKubernetesEngineConfig"
)

type ClusterSpec struct {
	AzureKubernetesServiceConfig  *AzureKubernetesServiceConfig  `json:"azureKubernetesServiceConfig,omitempty"`
	GoogleKubernetesEngineConfig  *GoogleKubernetesEngineConfig  `json:"googleKubernetesEngineConfig,omitempty"`
	RancherKubernetesEngineConfig *RancherKubernetesEngineConfig `json:"rancherKubernetesEngineConfig,omitempty"`
}
