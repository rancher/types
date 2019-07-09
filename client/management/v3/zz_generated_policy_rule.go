package client

const (
	PolicyRuleType                 = "policyRule"
	PolicyRuleFieldAPIGroups       = "apiGroups"
	PolicyRuleFieldNonResourceURLs = "nonResourceURLs"
	PolicyRuleFieldResourceNames   = "resourceNames"
	PolicyRuleFieldResources       = "resources"
	PolicyRuleFieldVerbs           = "verbs"
)

type PolicyRule struct {
	APIGroups       []string `json:"apiGroups,omitempty" yaml:"api_groups,omitempty"`
	NonResourceURLs []string `json:"nonResourceURLs,omitempty" yaml:"non_resource_ur_ls,omitempty"`
	ResourceNames   []string `json:"resourceNames,omitempty" yaml:"resource_names,omitempty"`
	Resources       []string `json:"resources,omitempty" yaml:"resources,omitempty"`
	Verbs           []string `json:"verbs,omitempty" yaml:"verbs,omitempty"`
}
