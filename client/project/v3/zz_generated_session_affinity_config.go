package client

const (
	SessionAffinityConfigType          = "sessionAffinityConfig"
	SessionAffinityConfigFieldClientIP = "clientIP"
)

type SessionAffinityConfig struct {
	ClientIP *ClientIPConfig `json:"clientIP,omitempty" yaml:"client_ip,omitempty"`
}
