package client

const (
	ClusterNodeStatusType                 = "clusterNodeStatus"
	ClusterNodeStatusFieldAddresses       = "addresses"
	ClusterNodeStatusFieldAllocatable     = "allocatable"
	ClusterNodeStatusFieldCapacity        = "capacity"
	ClusterNodeStatusFieldConditions      = "conditions"
	ClusterNodeStatusFieldDaemonEndpoints = "daemonEndpoints"
	ClusterNodeStatusFieldImages          = "images"
	ClusterNodeStatusFieldLimits          = "limits"
	ClusterNodeStatusFieldNodeInfo        = "nodeInfo"
	ClusterNodeStatusFieldPhase           = "phase"
	ClusterNodeStatusFieldRequested       = "requested"
	ClusterNodeStatusFieldVolumesAttached = "volumesAttached"
	ClusterNodeStatusFieldVolumesInUse    = "volumesInUse"
)

type ClusterNodeStatus struct {
	Addresses       []NodeAddress        `json:"addresses,omitempty"`
	Allocatable     map[string]string    `json:"allocatable,omitempty"`
	Capacity        map[string]string    `json:"capacity,omitempty"`
	Conditions      []NodeCondition      `json:"conditions,omitempty"`
	DaemonEndpoints *NodeDaemonEndpoints `json:"daemonEndpoints,omitempty"`
	Images          []ContainerImage     `json:"images,omitempty"`
	Limits          map[string]string    `json:"limits,omitempty"`
	NodeInfo        *NodeSystemInfo      `json:"nodeInfo,omitempty"`
	Phase           string               `json:"phase,omitempty"`
	Requested       map[string]string    `json:"requested,omitempty"`
	VolumesAttached []AttachedVolume     `json:"volumesAttached,omitempty"`
	VolumesInUse    []string             `json:"volumesInUse,omitempty"`
}
