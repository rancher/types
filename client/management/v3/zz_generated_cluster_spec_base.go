package client

const (
	ClusterSpecBaseType                                     = "clusterSpecBase"
	ClusterSpecBaseFieldDefaultClusterRoleForProjectMembers = "defaultClusterRoleForProjectMembers"
	ClusterSpecBaseFieldDefaultPodSecurityPolicyTemplateID  = "defaultPodSecurityPolicyTemplateId"
	ClusterSpecBaseFieldDesiredAgentImage                   = "desiredAgentImage"
	ClusterSpecBaseFieldDesiredAuthImage                    = "desiredAuthImage"
	ClusterSpecBaseFieldDockerRootDir                       = "dockerRootDir"
	ClusterSpecBaseFieldEnableClusterAlerting               = "enableClusterAlerting"
	ClusterSpecBaseFieldEnableClusterMonitoring             = "enableClusterMonitoring"
	ClusterSpecBaseFieldEnableNetworkPolicy                 = "enableNetworkPolicy"
	ClusterSpecBaseFieldLocalClusterAuthEndpoint            = "localClusterAuthEndpoint"
	ClusterSpecBaseFieldRancherKubernetesEngineConfig       = "rancherKubernetesEngineConfig"
)

type ClusterSpecBase struct {
	DefaultClusterRoleForProjectMembers string                         `json:"defaultClusterRoleForProjectMembers,omitempty" yaml:"default_cluster_role_for_project_members,omitempty"`
	DefaultPodSecurityPolicyTemplateID  string                         `json:"defaultPodSecurityPolicyTemplateId,omitempty" yaml:"default_pod_security_policy_template_id,omitempty"`
	DesiredAgentImage                   string                         `json:"desiredAgentImage,omitempty" yaml:"desired_agent_image,omitempty"`
	DesiredAuthImage                    string                         `json:"desiredAuthImage,omitempty" yaml:"desired_auth_image,omitempty"`
	DockerRootDir                       string                         `json:"dockerRootDir,omitempty" yaml:"docker_root_dir,omitempty"`
	EnableClusterAlerting               bool                           `json:"enableClusterAlerting,omitempty" yaml:"enable_cluster_alerting,omitempty"`
	EnableClusterMonitoring             bool                           `json:"enableClusterMonitoring,omitempty" yaml:"enable_cluster_monitoring,omitempty"`
	EnableNetworkPolicy                 *bool                          `json:"enableNetworkPolicy,omitempty" yaml:"enable_network_policy,omitempty"`
	LocalClusterAuthEndpoint            *LocalClusterAuthEndpoint      `json:"localClusterAuthEndpoint,omitempty" yaml:"local_cluster_auth_endpoint,omitempty"`
	RancherKubernetesEngineConfig       *RancherKubernetesEngineConfig `json:"rancherKubernetesEngineConfig,omitempty" yaml:"rancher_kubernetes_engine_config,omitempty"`
}
