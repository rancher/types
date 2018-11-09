package client

const (
	GlobalDNSStatusType           = "globalDNSStatus"
	GlobalDNSStatusFieldEndpoints = "endpoints"
)

type GlobalDNSStatus struct {
	Endpoints []string `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
}
