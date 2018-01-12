package client

const (
	MachineSpecType                   = "machineSpec"
	MachineSpecFieldClusterId         = "clusterId"
	MachineSpecFieldDescription       = "description"
	MachineSpecFieldDisplayName       = "displayName"
	MachineSpecFieldMachineTemplateId = "machineTemplateId"
	MachineSpecFieldPodCidr           = "podCidr"
	MachineSpecFieldProviderId        = "providerId"
	MachineSpecFieldRequestedHostname = "requestedHostname"
	MachineSpecFieldRole              = "role"
	MachineSpecFieldTaints            = "taints"
	MachineSpecFieldUnschedulable     = "unschedulable"
)

type MachineSpec struct {
	ClusterId         string   `json:"clusterId,omitempty"`
	Description       string   `json:"description,omitempty"`
	DisplayName       string   `json:"displayName,omitempty"`
	MachineTemplateId string   `json:"machineTemplateId,omitempty"`
	PodCidr           string   `json:"podCidr,omitempty"`
	ProviderId        string   `json:"providerId,omitempty"`
	RequestedHostname string   `json:"requestedHostname,omitempty"`
	Role              []string `json:"role,omitempty"`
	Taints            []Taint  `json:"taints,omitempty"`
	Unschedulable     *bool    `json:"unschedulable,omitempty"`
}
