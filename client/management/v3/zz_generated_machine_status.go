package client

const (
	MachineStatusType                 = "machineStatus"
	MachineStatusFieldAllocatable     = "allocatable"
	MachineStatusFieldCapacity        = "capacity"
	MachineStatusFieldConditions      = "conditions"
	MachineStatusFieldHostname        = "hostname"
	MachineStatusFieldIPAddress       = "ipAddress"
	MachineStatusFieldInfo            = "info"
	MachineStatusFieldLimits          = "limits"
	MachineStatusFieldNodeName        = "nodeName"
	MachineStatusFieldPhase           = "phase"
	MachineStatusFieldRequested       = "requested"
	MachineStatusFieldVolumesAttached = "volumesAttached"
	MachineStatusFieldVolumesInUse    = "volumesInUse"
)

type MachineStatus struct {
	Allocatable     map[string]string         `json:"allocatable,omitempty"`
	Capacity        map[string]string         `json:"capacity,omitempty"`
	Conditions      []NodeCondition           `json:"conditions,omitempty"`
	Hostname        string                    `json:"hostname,omitempty"`
	IPAddress       string                    `json:"ipAddress,omitempty"`
	Info            *NodeInfo                 `json:"info,omitempty"`
	Limits          map[string]string         `json:"limits,omitempty"`
	NodeName        string                    `json:"nodeName,omitempty"`
	Phase           string                    `json:"phase,omitempty"`
	Requested       map[string]string         `json:"requested,omitempty"`
	VolumesAttached map[string]AttachedVolume `json:"volumesAttached,omitempty"`
	VolumesInUse    []string                  `json:"volumesInUse,omitempty"`
}
