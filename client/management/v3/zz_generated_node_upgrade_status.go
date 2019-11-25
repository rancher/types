package client

const (
	NodeUpgradeStatusType       = "nodeUpgradeStatus"
	NodeUpgradeStatusFieldNodes = "nodes"
	NodeUpgradeStatusFieldState = "state"
	NodeUpgradeStatusFieldToken = "token"
)

type NodeUpgradeStatus struct {
	Nodes map[string]bool `json:"nodes,omitempty" yaml:"nodes,omitempty"`
	State string          `json:"state,omitempty" yaml:"state,omitempty"`
	Token string          `json:"token,omitempty" yaml:"token,omitempty"`
}
