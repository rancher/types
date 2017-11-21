package client

const (
	GoogleKubernetesEngineConfigType                    = "googleKubernetesEngineConfig"
	GoogleKubernetesEngineConfigFieldClusterIpv4Cidr    = "clusterIpv4Cidr"
	GoogleKubernetesEngineConfigFieldCredentialPath     = "credentialPath"
	GoogleKubernetesEngineConfigFieldDescription        = "description"
	GoogleKubernetesEngineConfigFieldDiskSizeGb         = "diskSizeGb"
	GoogleKubernetesEngineConfigFieldEnableAlphaFeature = "enableAlphaFeature"
	GoogleKubernetesEngineConfigFieldLabels             = "labels"
	GoogleKubernetesEngineConfigFieldMachineType        = "machineType"
	GoogleKubernetesEngineConfigFieldMasterVersion      = "masterVersion"
	GoogleKubernetesEngineConfigFieldNodeCount          = "nodeCount"
	GoogleKubernetesEngineConfigFieldNodeVersion        = "nodeVersion"
	GoogleKubernetesEngineConfigFieldProjectID          = "projectId"
	GoogleKubernetesEngineConfigFieldZone               = "zone"
)

type GoogleKubernetesEngineConfig struct {
	ClusterIpv4Cidr    string            `json:"clusterIpv4Cidr,omitempty"`
	CredentialPath     string            `json:"credentialPath,omitempty"`
	Description        string            `json:"description,omitempty"`
	DiskSizeGb         *int64            `json:"diskSizeGb,omitempty"`
	EnableAlphaFeature *bool             `json:"enableAlphaFeature,omitempty"`
	Labels             map[string]string `json:"labels,omitempty"`
	MachineType        string            `json:"machineType,omitempty"`
	MasterVersion      string            `json:"masterVersion,omitempty"`
	NodeCount          *int64            `json:"nodeCount,omitempty"`
	NodeVersion        string            `json:"nodeVersion,omitempty"`
	ProjectID          string            `json:"projectId,omitempty"`
	Zone               string            `json:"zone,omitempty"`
}
