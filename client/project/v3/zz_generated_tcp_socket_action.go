package client

const (
    TCPSocketActionType = "tcpSocketAction"
	TCPSocketActionFieldHost = "host"
	TCPSocketActionFieldPort = "port"
)

type TCPSocketAction struct {
        Host string `json:"host,omitempty" yaml:"host,omitempty"`
        Port intstr.IntOrString `json:"port,omitempty" yaml:"port,omitempty"`
}

