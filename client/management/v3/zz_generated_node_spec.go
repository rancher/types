package client

const (
	NodeSpecType                          = "nodeSpec"
	NodeSpecFieldControlPlane             = "controlPlane"
	NodeSpecFieldCurrentNodeAnnotations   = "currentNodeAnnotations"
	NodeSpecFieldCurrentNodeLabels        = "currentNodeLabels"
	NodeSpecFieldCustomConfig             = "customConfig"
	NodeSpecFieldDescription              = "description"
	NodeSpecFieldDesiredNodeAnnotations   = "desiredNodeAnnotations"
	NodeSpecFieldDesiredNodeLabels        = "desiredNodeLabels"
	NodeSpecFieldDesiredNodeUnschedulable = "desiredNodeUnschedulable"
	NodeSpecFieldDisplayName              = "displayName"
	NodeSpecFieldEtcd                     = "etcd"
	NodeSpecFieldImported                 = "imported"
	NodeSpecFieldNodeDrainInput           = "nodeDrainInput"
	NodeSpecFieldNodePoolID               = "nodePoolId"
	NodeSpecFieldNodeTemplateID           = "nodeTemplateId"
	NodeSpecFieldPodCidr                  = "podCidr"
	NodeSpecFieldProviderId               = "providerId"
	NodeSpecFieldRequestedHostname        = "requestedHostname"
	NodeSpecFieldTaints                   = "taints"
	NodeSpecFieldUnschedulable            = "unschedulable"
	NodeSpecFieldWorker                   = "worker"
)

type NodeSpec struct {
	ControlPlane             bool              `json:"controlPlane,omitempty" yaml:"control_plane,omitempty"`
	CurrentNodeAnnotations   map[string]string `json:"currentNodeAnnotations,omitempty" yaml:"current_node_annotations,omitempty"`
	CurrentNodeLabels        map[string]string `json:"currentNodeLabels,omitempty" yaml:"current_node_labels,omitempty"`
	CustomConfig             *CustomConfig     `json:"customConfig,omitempty" yaml:"custom_config,omitempty"`
	Description              string            `json:"description,omitempty" yaml:"description,omitempty"`
	DesiredNodeAnnotations   map[string]string `json:"desiredNodeAnnotations,omitempty" yaml:"desired_node_annotations,omitempty"`
	DesiredNodeLabels        map[string]string `json:"desiredNodeLabels,omitempty" yaml:"desired_node_labels,omitempty"`
	DesiredNodeUnschedulable string            `json:"desiredNodeUnschedulable,omitempty" yaml:"desired_node_unschedulable,omitempty"`
	DisplayName              string            `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	Etcd                     bool              `json:"etcd,omitempty" yaml:"etcd,omitempty"`
	Imported                 bool              `json:"imported,omitempty" yaml:"imported,omitempty"`
	NodeDrainInput           *NodeDrainInput   `json:"nodeDrainInput,omitempty" yaml:"node_drain_input,omitempty"`
	NodePoolID               string            `json:"nodePoolId,omitempty" yaml:"node_pool_id,omitempty"`
	NodeTemplateID           string            `json:"nodeTemplateId,omitempty" yaml:"node_template_id,omitempty"`
	PodCidr                  string            `json:"podCidr,omitempty" yaml:"pod_cidr,omitempty"`
	ProviderId               string            `json:"providerId,omitempty" yaml:"provider_id,omitempty"`
	RequestedHostname        string            `json:"requestedHostname,omitempty" yaml:"requested_hostname,omitempty"`
	Taints                   []Taint           `json:"taints,omitempty" yaml:"taints,omitempty"`
	Unschedulable            bool              `json:"unschedulable,omitempty" yaml:"unschedulable,omitempty"`
	Worker                   bool              `json:"worker,omitempty" yaml:"worker,omitempty"`
}
