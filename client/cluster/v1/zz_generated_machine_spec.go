package client

const (
	MachineSpecType                          = "machineSpec"
	MachineSpecFieldAmazonEC2Config          = "amazonEc2Config"
	MachineSpecFieldAuthCertificateAuthority = "authCertificateAuthority"
	MachineSpecFieldAuthKey                  = "authKey"
	MachineSpecFieldAzureConfig              = "azureConfig"
	MachineSpecFieldClusterId                = "clusterId"
	MachineSpecFieldDescription              = "description"
	MachineSpecFieldDigitalOceanConfig       = "digitalOceanConfig"
	MachineSpecFieldDisplayName              = "displayName"
	MachineSpecFieldDockerVersion            = "dockerVersion"
	MachineSpecFieldDriver                   = "driver"
	MachineSpecFieldEngineEnv                = "engineEnv"
	MachineSpecFieldEngineInsecureRegistry   = "engineInsecureRegistry"
	MachineSpecFieldEngineInstallURL         = "engineInstallURL"
	MachineSpecFieldEngineLabel              = "engineLabel"
	MachineSpecFieldEngineOpt                = "engineOpt"
	MachineSpecFieldEngineRegistryMirror     = "engineRegistryMirror"
	MachineSpecFieldEngineStorageDriver      = "engineStorageDriver"
	MachineSpecFieldExternalID               = "externalId"
	MachineSpecFieldHostname                 = "hostname"
	MachineSpecFieldMachineTemplateId        = "machineTemplateId"
)

type MachineSpec struct {
	AmazonEC2Config          *AmazonEC2Config    `json:"amazonEc2Config,omitempty"`
	AuthCertificateAuthority string              `json:"authCertificateAuthority,omitempty"`
	AuthKey                  string              `json:"authKey,omitempty"`
	AzureConfig              *AzureConfig        `json:"azureConfig,omitempty"`
	ClusterId                string              `json:"clusterId,omitempty"`
	Description              string              `json:"description,omitempty"`
	DigitalOceanConfig       *DigitalOceanConfig `json:"digitalOceanConfig,omitempty"`
	DisplayName              string              `json:"displayName,omitempty"`
	DockerVersion            string              `json:"dockerVersion,omitempty"`
	Driver                   string              `json:"driver,omitempty"`
	EngineEnv                map[string]string   `json:"engineEnv,omitempty"`
	EngineInsecureRegistry   []string            `json:"engineInsecureRegistry,omitempty"`
	EngineInstallURL         string              `json:"engineInstallURL,omitempty"`
	EngineLabel              map[string]string   `json:"engineLabel,omitempty"`
	EngineOpt                map[string]string   `json:"engineOpt,omitempty"`
	EngineRegistryMirror     []string            `json:"engineRegistryMirror,omitempty"`
	EngineStorageDriver      string              `json:"engineStorageDriver,omitempty"`
	ExternalID               string              `json:"externalId,omitempty"`
	Hostname                 string              `json:"hostname,omitempty"`
	MachineTemplateId        string              `json:"machineTemplateId,omitempty"`
}
