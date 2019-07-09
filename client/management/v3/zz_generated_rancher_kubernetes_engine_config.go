package client

const (
	RancherKubernetesEngineConfigType                     = "rancherKubernetesEngineConfig"
	RancherKubernetesEngineConfigFieldAddonJobTimeout     = "addonJobTimeout"
	RancherKubernetesEngineConfigFieldAddons              = "addons"
	RancherKubernetesEngineConfigFieldAddonsInclude       = "addonsInclude"
	RancherKubernetesEngineConfigFieldAuthentication      = "authentication"
	RancherKubernetesEngineConfigFieldAuthorization       = "authorization"
	RancherKubernetesEngineConfigFieldBastionHost         = "bastionHost"
	RancherKubernetesEngineConfigFieldCloudProvider       = "cloudProvider"
	RancherKubernetesEngineConfigFieldClusterName         = "clusterName"
	RancherKubernetesEngineConfigFieldDNS                 = "dns"
	RancherKubernetesEngineConfigFieldIgnoreDockerVersion = "ignoreDockerVersion"
	RancherKubernetesEngineConfigFieldIngress             = "ingress"
	RancherKubernetesEngineConfigFieldMonitoring          = "monitoring"
	RancherKubernetesEngineConfigFieldNetwork             = "network"
	RancherKubernetesEngineConfigFieldNodes               = "nodes"
	RancherKubernetesEngineConfigFieldPrefixPath          = "prefixPath"
	RancherKubernetesEngineConfigFieldPrivateRegistries   = "privateRegistries"
	RancherKubernetesEngineConfigFieldRestore             = "restore"
	RancherKubernetesEngineConfigFieldRotateCertificates  = "rotateCertificates"
	RancherKubernetesEngineConfigFieldSSHAgentAuth        = "sshAgentAuth"
	RancherKubernetesEngineConfigFieldSSHCertPath         = "sshCertPath"
	RancherKubernetesEngineConfigFieldSSHKeyPath          = "sshKeyPath"
	RancherKubernetesEngineConfigFieldServices            = "services"
	RancherKubernetesEngineConfigFieldVersion             = "kubernetesVersion"
)

type RancherKubernetesEngineConfig struct {
	AddonJobTimeout     int64               `json:"addonJobTimeout,omitempty" yaml:"addon_job_timeout,omitempty"`
	Addons              string              `json:"addons,omitempty" yaml:"addons,omitempty"`
	AddonsInclude       []string            `json:"addonsInclude,omitempty" yaml:"addons_include,omitempty"`
	Authentication      *AuthnConfig        `json:"authentication,omitempty" yaml:"authentication,omitempty"`
	Authorization       *AuthzConfig        `json:"authorization,omitempty" yaml:"authorization,omitempty"`
	BastionHost         *BastionHost        `json:"bastionHost,omitempty" yaml:"bastion_host,omitempty"`
	CloudProvider       *CloudProvider      `json:"cloudProvider,omitempty" yaml:"cloud_provider,omitempty"`
	ClusterName         string              `json:"clusterName,omitempty" yaml:"cluster_name,omitempty"`
	DNS                 *DNSConfig          `json:"dns,omitempty" yaml:"dns,omitempty"`
	IgnoreDockerVersion bool                `json:"ignoreDockerVersion,omitempty" yaml:"ignore_docker_version,omitempty"`
	Ingress             *IngressConfig      `json:"ingress,omitempty" yaml:"ingress,omitempty"`
	Monitoring          *MonitoringConfig   `json:"monitoring,omitempty" yaml:"monitoring,omitempty"`
	Network             *NetworkConfig      `json:"network,omitempty" yaml:"network,omitempty"`
	Nodes               []RKEConfigNode     `json:"nodes,omitempty" yaml:"nodes,omitempty"`
	PrefixPath          string              `json:"prefixPath,omitempty" yaml:"prefix_path,omitempty"`
	PrivateRegistries   []PrivateRegistry   `json:"privateRegistries,omitempty" yaml:"private_registries,omitempty"`
	Restore             *RestoreConfig      `json:"restore,omitempty" yaml:"restore,omitempty"`
	RotateCertificates  *RotateCertificates `json:"rotateCertificates,omitempty" yaml:"rotate_certificates,omitempty"`
	SSHAgentAuth        bool                `json:"sshAgentAuth,omitempty" yaml:"ssh_agent_auth,omitempty"`
	SSHCertPath         string              `json:"sshCertPath,omitempty" yaml:"ssh_cert_path,omitempty"`
	SSHKeyPath          string              `json:"sshKeyPath,omitempty" yaml:"ssh_key_path,omitempty"`
	Services            *RKEConfigServices  `json:"services,omitempty" yaml:"services,omitempty"`
	Version             string              `json:"kubernetesVersion,omitempty" yaml:"kubernetes_version,omitempty"`
}
