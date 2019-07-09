package client

const (
	NodeDaemonEndpointsType                 = "nodeDaemonEndpoints"
	NodeDaemonEndpointsFieldKubeletEndpoint = "kubeletEndpoint"
)

type NodeDaemonEndpoints struct {
	KubeletEndpoint *DaemonEndpoint `json:"kubeletEndpoint,omitempty" yaml:"kubelet_endpoint,omitempty"`
}
