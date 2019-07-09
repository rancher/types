package client

const (
	UpgradeStrategyType               = "upgradeStrategy"
	UpgradeStrategyFieldRollingUpdate = "rollingUpdate"
)

type UpgradeStrategy struct {
	RollingUpdate *RollingUpdate `json:"rollingUpdate,omitempty" yaml:"rolling_update,omitempty"`
}
