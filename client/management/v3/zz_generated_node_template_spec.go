package client

const (
	NodeTemplateSpecType                          = "nodeTemplateSpec"
	NodeTemplateSpecFieldAuthCertificateAuthority = "authCertificateAuthority"
	NodeTemplateSpecFieldAuthKey                  = "authKey"
	NodeTemplateSpecFieldCloudCredentialID        = "cloudCredentialId"
	NodeTemplateSpecFieldDescription              = "description"
	NodeTemplateSpecFieldDisplayName              = "displayName"
	NodeTemplateSpecFieldDockerVersion            = "dockerVersion"
	NodeTemplateSpecFieldDriver                   = "driver"
	NodeTemplateSpecFieldEngineEnv                = "engineEnv"
	NodeTemplateSpecFieldEngineInsecureRegistry   = "engineInsecureRegistry"
	NodeTemplateSpecFieldEngineInstallURL         = "engineInstallURL"
	NodeTemplateSpecFieldEngineLabel              = "engineLabel"
	NodeTemplateSpecFieldEngineOpt                = "engineOpt"
	NodeTemplateSpecFieldEngineRegistryMirror     = "engineRegistryMirror"
	NodeTemplateSpecFieldEngineStorageDriver      = "engineStorageDriver"
	NodeTemplateSpecFieldUseInternalIPAddress     = "useInternalIpAddress"
)

type NodeTemplateSpec struct {
	AuthCertificateAuthority string            `json:"authCertificateAuthority,omitempty" yaml:"auth_certificate_authority,omitempty"`
	AuthKey                  string            `json:"authKey,omitempty" yaml:"auth_key,omitempty"`
	CloudCredentialID        string            `json:"cloudCredentialId,omitempty" yaml:"cloud_credential_id,omitempty"`
	Description              string            `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName              string            `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	DockerVersion            string            `json:"dockerVersion,omitempty" yaml:"docker_version,omitempty"`
	Driver                   string            `json:"driver,omitempty" yaml:"driver,omitempty"`
	EngineEnv                map[string]string `json:"engineEnv,omitempty" yaml:"engine_env,omitempty"`
	EngineInsecureRegistry   []string          `json:"engineInsecureRegistry,omitempty" yaml:"engine_insecure_registry,omitempty"`
	EngineInstallURL         string            `json:"engineInstallURL,omitempty" yaml:"engine_install_url,omitempty"`
	EngineLabel              map[string]string `json:"engineLabel,omitempty" yaml:"engine_label,omitempty"`
	EngineOpt                map[string]string `json:"engineOpt,omitempty" yaml:"engine_opt,omitempty"`
	EngineRegistryMirror     []string          `json:"engineRegistryMirror,omitempty" yaml:"engine_registry_mirror,omitempty"`
	EngineStorageDriver      string            `json:"engineStorageDriver,omitempty" yaml:"engine_storage_driver,omitempty"`
	UseInternalIPAddress     bool              `json:"useInternalIpAddress,omitempty" yaml:"use_internal_ip_address,omitempty"`
}
