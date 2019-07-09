package client

const (
	ClusterSpecType                                     = "clusterSpec"
	ClusterSpecFieldAmazonElasticContainerServiceConfig = "amazonElasticContainerServiceConfig"
	ClusterSpecFieldAzureKubernetesServiceConfig        = "azureKubernetesServiceConfig"
	ClusterSpecFieldClusterTemplateAnswers              = "answers"
	ClusterSpecFieldClusterTemplateID                   = "clusterTemplateId"
	ClusterSpecFieldClusterTemplateRevisionID           = "clusterTemplateRevisionId"
	ClusterSpecFieldDefaultClusterRoleForProjectMembers = "defaultClusterRoleForProjectMembers"
	ClusterSpecFieldDefaultPodSecurityPolicyTemplateID  = "defaultPodSecurityPolicyTemplateId"
	ClusterSpecFieldDescription                         = "description"
	ClusterSpecFieldDesiredAgentImage                   = "desiredAgentImage"
	ClusterSpecFieldDesiredAuthImage                    = "desiredAuthImage"
	ClusterSpecFieldDisplayName                         = "displayName"
	ClusterSpecFieldDockerRootDir                       = "dockerRootDir"
	ClusterSpecFieldEnableClusterAlerting               = "enableClusterAlerting"
	ClusterSpecFieldEnableClusterMonitoring             = "enableClusterMonitoring"
	ClusterSpecFieldEnableNetworkPolicy                 = "enableNetworkPolicy"
	ClusterSpecFieldGenericEngineConfig                 = "genericEngineConfig"
	ClusterSpecFieldGoogleKubernetesEngineConfig        = "googleKubernetesEngineConfig"
	ClusterSpecFieldImportedConfig                      = "importedConfig"
	ClusterSpecFieldInternal                            = "internal"
	ClusterSpecFieldLocalClusterAuthEndpoint            = "localClusterAuthEndpoint"
	ClusterSpecFieldRancherKubernetesEngineConfig       = "rancherKubernetesEngineConfig"
)

type ClusterSpec struct {
	AmazonElasticContainerServiceConfig map[string]interface{}         `json:"amazonElasticContainerServiceConfig,omitempty" yaml:"amazon_elastic_container_service_config,omitempty"`
	AzureKubernetesServiceConfig        map[string]interface{}         `json:"azureKubernetesServiceConfig,omitempty" yaml:"azure_kubernetes_service_config,omitempty"`
	ClusterTemplateAnswers              *Answer                        `json:"answers,omitempty" yaml:"answers,omitempty"`
	ClusterTemplateID                   string                         `json:"clusterTemplateId,omitempty" yaml:"cluster_template_id,omitempty"`
	ClusterTemplateRevisionID           string                         `json:"clusterTemplateRevisionId,omitempty" yaml:"cluster_template_revision_id,omitempty"`
	DefaultClusterRoleForProjectMembers string                         `json:"defaultClusterRoleForProjectMembers,omitempty" yaml:"default_cluster_role_for_project_members,omitempty"`
	DefaultPodSecurityPolicyTemplateID  string                         `json:"defaultPodSecurityPolicyTemplateId,omitempty" yaml:"default_pod_security_policy_template_id,omitempty"`
	Description                         string                         `json:"description,omitempty" yaml:"description,omitempty"`
	DesiredAgentImage                   string                         `json:"desiredAgentImage,omitempty" yaml:"desired_agent_image,omitempty"`
	DesiredAuthImage                    string                         `json:"desiredAuthImage,omitempty" yaml:"desired_auth_image,omitempty"`
	DisplayName                         string                         `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	DockerRootDir                       string                         `json:"dockerRootDir,omitempty" yaml:"docker_root_dir,omitempty"`
	EnableClusterAlerting               bool                           `json:"enableClusterAlerting,omitempty" yaml:"enable_cluster_alerting,omitempty"`
	EnableClusterMonitoring             bool                           `json:"enableClusterMonitoring,omitempty" yaml:"enable_cluster_monitoring,omitempty"`
	EnableNetworkPolicy                 *bool                          `json:"enableNetworkPolicy,omitempty" yaml:"enable_network_policy,omitempty"`
	GenericEngineConfig                 map[string]interface{}         `json:"genericEngineConfig,omitempty" yaml:"generic_engine_config,omitempty"`
	GoogleKubernetesEngineConfig        map[string]interface{}         `json:"googleKubernetesEngineConfig,omitempty" yaml:"google_kubernetes_engine_config,omitempty"`
	ImportedConfig                      *ImportedConfig                `json:"importedConfig,omitempty" yaml:"imported_config,omitempty"`
	Internal                            bool                           `json:"internal,omitempty" yaml:"internal,omitempty"`
	LocalClusterAuthEndpoint            *LocalClusterAuthEndpoint      `json:"localClusterAuthEndpoint,omitempty" yaml:"local_cluster_auth_endpoint,omitempty"`
	RancherKubernetesEngineConfig       *RancherKubernetesEngineConfig `json:"rancherKubernetesEngineConfig,omitempty" yaml:"rancher_kubernetes_engine_config,omitempty"`
}
