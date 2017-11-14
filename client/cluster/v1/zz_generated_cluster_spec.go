package client

const (
	ClusterSpecType                              = "clusterSpec"
	ClusterSpecFieldAzureKubernetesServiceConfig = "azureKubernetesServiceConfig"
	ClusterSpecFieldGoogleKubernetesEngineConfig = "googleKubernetesEngineConfig"
	ClusterSpecFieldRKEConfig                    = "rkeConfig"
)

type ClusterSpec struct {
	AzureKubernetesServiceConfig *AzureKubernetesServiceConfig `json:"azureKubernetesServiceConfig,omitempty"`
	GoogleKubernetesEngineConfig *GoogleKubernetesEngineConfig `json:"googleKubernetesEngineConfig,omitempty"`
	RKEConfig                    *RKEConfig                    `json:"rkeConfig,omitempty"`
}
