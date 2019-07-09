package client

const (
	IngressConfigType              = "ingressConfig"
	IngressConfigFieldExtraArgs    = "extraArgs"
	IngressConfigFieldNodeSelector = "nodeSelector"
	IngressConfigFieldOptions      = "options"
	IngressConfigFieldProvider     = "provider"
)

type IngressConfig struct {
	ExtraArgs    map[string]string `json:"extraArgs,omitempty" yaml:"extra_args,omitempty"`
	NodeSelector map[string]string `json:"nodeSelector,omitempty" yaml:"node_selector,omitempty"`
	Options      map[string]string `json:"options,omitempty" yaml:"options,omitempty"`
	Provider     string            `json:"provider,omitempty" yaml:"provider,omitempty"`
}
