package client

const (
	AciNetworkProviderType                        = "aciNetworkProvider"
	AciNetworkProviderFieldAEP                    = "aep"
	AciNetworkProviderFieldApicHosts              = "apicHosts"
	AciNetworkProviderFieldApicRefreshTime        = "apicRefreshTime"
	AciNetworkProviderFieldApicUserCrt            = "apicUserCrt"
	AciNetworkProviderFieldApicUserKey            = "apicUserKey"
	AciNetworkProviderFieldApicUserName           = "apicUserName"
	AciNetworkProviderFieldControllerLogLevel     = "controllerLogLevel"
	AciNetworkProviderFieldDropLogEnable          = "dropLogEnable"
	AciNetworkProviderFieldDynamicExternalSubnet  = "dynamicExternalSubnet"
	AciNetworkProviderFieldEncapType              = "encapType"
	AciNetworkProviderFieldHostAgentLogLevel      = "hostAgentLogLevel"
	AciNetworkProviderFieldImagePullPolicy        = "imagePullPolicy"
	AciNetworkProviderFieldImagePullSecret        = "imagePullSecret"
	AciNetworkProviderFieldInfraVlan              = "infraVlan"
	AciNetworkProviderFieldInstallIstio           = "installIstio"
	AciNetworkProviderFieldIstioProfile           = "istioProfile"
	AciNetworkProviderFieldKubeAPIVlan            = "kubeAPIVlan"
	AciNetworkProviderFieldL3Out                  = "l3Out"
	AciNetworkProviderFieldL3OutExternalNetworks  = "l3OutExternalNetworks"
	AciNetworkProviderFieldMcastRangeEnd          = "mcastRangeEnd"
	AciNetworkProviderFieldMcastRangeStart        = "mcastRangeStart"
	AciNetworkProviderFieldNodeSubnet             = "nodeSubnet"
	AciNetworkProviderFieldOVSMemoryLimit         = "ovsMemoryLimit"
	AciNetworkProviderFieldOpflexAgentLogLevel    = "opflexLogLevel"
	AciNetworkProviderFieldPBRTrackingNonSnat     = "pbrTrackingNonSnat"
	AciNetworkProviderFieldServiceGraphSubnet     = "serviceGraphSubnet"
	AciNetworkProviderFieldServiceMonitorInterval = "serviceMonitorInterval"
	AciNetworkProviderFieldServiceVlan            = "serviceVlan"
	AciNetworkProviderFieldStaticExternalSubnet   = "staticExternalSubnet"
	AciNetworkProviderFieldSystemIdentifier       = "systemId"
	AciNetworkProviderFieldTenant                 = "tenant"
	AciNetworkProviderFieldToken                  = "token"
	AciNetworkProviderFieldVRFName                = "vrfName"
	AciNetworkProviderFieldVRFTenant              = "vrfTenant"
	AciNetworkProviderFieldVmmController          = "vmmController"
	AciNetworkProviderFieldVmmDomain              = "vmmDomain"
)

type AciNetworkProvider struct {
	AEP                    string   `json:"aep,omitempty" yaml:"aep,omitempty"`
	ApicHosts              []string `json:"apicHosts,omitempty" yaml:"apicHosts,omitempty"`
	ApicRefreshTime        string   `json:"apicRefreshTime,omitempty" yaml:"apicRefreshTime,omitempty"`
	ApicUserCrt            string   `json:"apicUserCrt,omitempty" yaml:"apicUserCrt,omitempty"`
	ApicUserKey            string   `json:"apicUserKey,omitempty" yaml:"apicUserKey,omitempty"`
	ApicUserName           string   `json:"apicUserName,omitempty" yaml:"apicUserName,omitempty"`
	ControllerLogLevel     string   `json:"controllerLogLevel,omitempty" yaml:"controllerLogLevel,omitempty"`
	DropLogEnable          string   `json:"dropLogEnable,omitempty" yaml:"dropLogEnable,omitempty"`
	DynamicExternalSubnet  string   `json:"dynamicExternalSubnet,omitempty" yaml:"dynamicExternalSubnet,omitempty"`
	EncapType              string   `json:"encapType,omitempty" yaml:"encapType,omitempty"`
	HostAgentLogLevel      string   `json:"hostAgentLogLevel,omitempty" yaml:"hostAgentLogLevel,omitempty"`
	ImagePullPolicy        string   `json:"imagePullPolicy,omitempty" yaml:"imagePullPolicy,omitempty"`
	ImagePullSecret        string   `json:"imagePullSecret,omitempty" yaml:"imagePullSecret,omitempty"`
	InfraVlan              string   `json:"infraVlan,omitempty" yaml:"infraVlan,omitempty"`
	InstallIstio           string   `json:"installIstio,omitempty" yaml:"installIstio,omitempty"`
	IstioProfile           string   `json:"istioProfile,omitempty" yaml:"istioProfile,omitempty"`
	KubeAPIVlan            string   `json:"kubeAPIVlan,omitempty" yaml:"kubeAPIVlan,omitempty"`
	L3Out                  string   `json:"l3Out,omitempty" yaml:"l3Out,omitempty"`
	L3OutExternalNetworks  []string `json:"l3OutExternalNetworks,omitempty" yaml:"l3OutExternalNetworks,omitempty"`
	McastRangeEnd          string   `json:"mcastRangeEnd,omitempty" yaml:"mcastRangeEnd,omitempty"`
	McastRangeStart        string   `json:"mcastRangeStart,omitempty" yaml:"mcastRangeStart,omitempty"`
	NodeSubnet             string   `json:"nodeSubnet,omitempty" yaml:"nodeSubnet,omitempty"`
	OVSMemoryLimit         string   `json:"ovsMemoryLimit,omitempty" yaml:"ovsMemoryLimit,omitempty"`
	OpflexAgentLogLevel    string   `json:"opflexLogLevel,omitempty" yaml:"opflexLogLevel,omitempty"`
	PBRTrackingNonSnat     string   `json:"pbrTrackingNonSnat,omitempty" yaml:"pbrTrackingNonSnat,omitempty"`
	ServiceGraphSubnet     string   `json:"serviceGraphSubnet,omitempty" yaml:"serviceGraphSubnet,omitempty"`
	ServiceMonitorInterval string   `json:"serviceMonitorInterval,omitempty" yaml:"serviceMonitorInterval,omitempty"`
	ServiceVlan            string   `json:"serviceVlan,omitempty" yaml:"serviceVlan,omitempty"`
	StaticExternalSubnet   string   `json:"staticExternalSubnet,omitempty" yaml:"staticExternalSubnet,omitempty"`
	SystemIdentifier       string   `json:"systemId,omitempty" yaml:"systemId,omitempty"`
	Tenant                 string   `json:"tenant,omitempty" yaml:"tenant,omitempty"`
	Token                  string   `json:"token,omitempty" yaml:"token,omitempty"`
	VRFName                string   `json:"vrfName,omitempty" yaml:"vrfName,omitempty"`
	VRFTenant              string   `json:"vrfTenant,omitempty" yaml:"vrfTenant,omitempty"`
	VmmController          string   `json:"vmmController,omitempty" yaml:"vmmController,omitempty"`
	VmmDomain              string   `json:"vmmDomain,omitempty" yaml:"vmmDomain,omitempty"`
}
