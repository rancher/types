package client

const (
	MachineStatusType            = "machineStatus"
	MachineStatusFieldConditions = "conditions"
)

type MachineStatus struct {
	Conditions []MachineCondition `json:"conditions,omitempty"`
}
