package client

const (
	NodeDriverStatusType                             = "nodeDriverStatus"
	NodeDriverStatusFieldAppliedChecksum             = "appliedChecksum"
	NodeDriverStatusFieldAppliedDockerMachineVersion = "appliedDockerMachineVersion"
	NodeDriverStatusFieldAppliedURL                  = "appliedURL"
	NodeDriverStatusFieldConditions                  = "conditions"
)

type NodeDriverStatus struct {
	AppliedChecksum             string      `json:"appliedChecksum,omitempty" yaml:"applied_checksum,omitempty"`
	AppliedDockerMachineVersion string      `json:"appliedDockerMachineVersion,omitempty" yaml:"applied_docker_machine_version,omitempty"`
	AppliedURL                  string      `json:"appliedURL,omitempty" yaml:"applied_url,omitempty"`
	Conditions                  []Condition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}
