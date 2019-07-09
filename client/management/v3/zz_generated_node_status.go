package client

const (
	NodeStatusType                   = "nodeStatus"
	NodeStatusFieldAllocatable       = "allocatable"
	NodeStatusFieldCapacity          = "capacity"
	NodeStatusFieldConditions        = "conditions"
	NodeStatusFieldDockerInfo        = "dockerInfo"
	NodeStatusFieldExternalIPAddress = "externalIpAddress"
	NodeStatusFieldHostname          = "hostname"
	NodeStatusFieldIPAddress         = "ipAddress"
	NodeStatusFieldInfo              = "info"
	NodeStatusFieldLimits            = "limits"
	NodeStatusFieldNodeAnnotations   = "nodeAnnotations"
	NodeStatusFieldNodeConfig        = "rkeNode"
	NodeStatusFieldNodeLabels        = "nodeLabels"
	NodeStatusFieldNodeName          = "nodeName"
	NodeStatusFieldNodeTaints        = "nodeTaints"
	NodeStatusFieldRequested         = "requested"
	NodeStatusFieldVolumesAttached   = "volumesAttached"
	NodeStatusFieldVolumesInUse      = "volumesInUse"
)

type NodeStatus struct {
	Allocatable       map[string]string         `json:"allocatable,omitempty" yaml:"allocatable,omitempty"`
	Capacity          map[string]string         `json:"capacity,omitempty" yaml:"capacity,omitempty"`
	Conditions        []NodeCondition           `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	DockerInfo        *DockerInfo               `json:"dockerInfo,omitempty" yaml:"docker_info,omitempty"`
	ExternalIPAddress string                    `json:"externalIpAddress,omitempty" yaml:"external_ip_address,omitempty"`
	Hostname          string                    `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	IPAddress         string                    `json:"ipAddress,omitempty" yaml:"ip_address,omitempty"`
	Info              *NodeInfo                 `json:"info,omitempty" yaml:"info,omitempty"`
	Limits            map[string]string         `json:"limits,omitempty" yaml:"limits,omitempty"`
	NodeAnnotations   map[string]string         `json:"nodeAnnotations,omitempty" yaml:"node_annotations,omitempty"`
	NodeConfig        *RKEConfigNode            `json:"rkeNode,omitempty" yaml:"rke_node,omitempty"`
	NodeLabels        map[string]string         `json:"nodeLabels,omitempty" yaml:"node_labels,omitempty"`
	NodeName          string                    `json:"nodeName,omitempty" yaml:"node_name,omitempty"`
	NodeTaints        []Taint                   `json:"nodeTaints,omitempty" yaml:"node_taints,omitempty"`
	Requested         map[string]string         `json:"requested,omitempty" yaml:"requested,omitempty"`
	VolumesAttached   map[string]AttachedVolume `json:"volumesAttached,omitempty" yaml:"volumes_attached,omitempty"`
	VolumesInUse      []string                  `json:"volumesInUse,omitempty" yaml:"volumes_in_use,omitempty"`
}
