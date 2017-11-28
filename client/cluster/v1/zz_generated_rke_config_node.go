package client

const (
	RKEConfigNodeType                  = "rkeConfigNode"
	RKEConfigNodeFieldAddress          = "address"
	RKEConfigNodeFieldDockerSocket     = "dockerSocket"
	RKEConfigNodeFieldHostnameOverride = "hostnameOverride"
	RKEConfigNodeFieldInternalAddress  = "internalAddress"
	RKEConfigNodeFieldRole             = "role"
	RKEConfigNodeFieldUser             = "user"
)

type RKEConfigNode struct {
	Address          string   `json:"address,omitempty"`
	DockerSocket     string   `json:"dockerSocket,omitempty"`
	HostnameOverride string   `json:"hostnameOverride,omitempty"`
	InternalAddress  string   `json:"internalAddress,omitempty"`
	Role             []string `json:"role,omitempty"`
	User             string   `json:"user,omitempty"`
}
