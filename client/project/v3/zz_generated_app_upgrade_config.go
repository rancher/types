package client

const (
	AppUpgradeConfigType              = "appUpgradeConfig"
	AppUpgradeConfigFieldAnswers      = "answers"
	AppUpgradeConfigFieldExternalID   = "externalId"
	AppUpgradeConfigFieldFiles        = "files"
	AppUpgradeConfigFieldForceUpgrade = "forceUpgrade"
	AppUpgradeConfigFieldValuesYaml   = "valuesYaml"
)

type AppUpgradeConfig struct {
	Answers      map[string]string `json:"answers,omitempty" yaml:"answers,omitempty"`
	ExternalID   string            `json:"externalId,omitempty" yaml:"external_id,omitempty"`
	Files        map[string]string `json:"files,omitempty" yaml:"files,omitempty"`
	ForceUpgrade bool              `json:"forceUpgrade,omitempty" yaml:"force_upgrade,omitempty"`
	ValuesYaml   string            `json:"valuesYaml,omitempty" yaml:"values_yaml,omitempty"`
}
