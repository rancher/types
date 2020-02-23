package client

const (
	NodeUpgradeStrategyType                       = "nodeUpgradeStrategy"
	NodeUpgradeStrategyFieldDrain                 = "drain"
	NodeUpgradeStrategyFieldDrainInput            = "nodeDrainInput"
	NodeUpgradeStrategyFieldMaxUnavailableControl = "maxUnavailableControl"
	NodeUpgradeStrategyFieldMaxUnavailableWorker  = "maxUnavailableWorker"
)

type NodeUpgradeStrategy struct {
	Drain                 bool            `json:"drain,omitempty" yaml:"drain,omitempty"`
	DrainInput            *NodeDrainInput `json:"nodeDrainInput,omitempty" yaml:"nodeDrainInput,omitempty"`
	MaxUnavailableControl string          `json:"maxUnavailableControl,omitempty" yaml:"maxUnavailableControl,omitempty"`
	MaxUnavailableWorker  string          `json:"maxUnavailableWorker,omitempty" yaml:"maxUnavailableWorker,omitempty"`
}
