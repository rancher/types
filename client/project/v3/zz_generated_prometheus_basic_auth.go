package client

const (
	PrometheusBasicAuthType          = "prometheusBasicAuth"
	PrometheusBasicAuthFieldPassword = "password"
	PrometheusBasicAuthFieldUsername = "username"
)

type PrometheusBasicAuth struct {
	Password *SecretKeySelector `json:"password,omitempty" yaml:"password,omitempty"`
	Username *SecretKeySelector `json:"username,omitempty" yaml:"username,omitempty"`
}
