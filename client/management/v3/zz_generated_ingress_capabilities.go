package client

const (
	IngressCapabilitiesType                      = "ingressCapabilities"
	IngressCapabilitiesFieldCustomDefaultBackend = "customDefaultBackend"
	IngressCapabilitiesFieldIngressProvider      = "ingressProvider"
)

type IngressCapabilities struct {
	CustomDefaultBackend bool   `json:"customDefaultBackend,omitempty" yaml:"custom_default_backend,omitempty"`
	IngressProvider      string `json:"ingressProvider,omitempty" yaml:"ingress_provider,omitempty"`
}
