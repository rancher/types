package client

const (
	AuthAppInputType                = "authAppInput"
	AuthAppInputFieldClientID       = "clientId"
	AuthAppInputFieldClientSecret   = "clientSecret"
	AuthAppInputFieldCode           = "code"
	AuthAppInputFieldHost           = "host"
	AuthAppInputFieldInheritGlobal  = "inheritGlobal"
	AuthAppInputFieldRedirectURL    = "redirectUrl"
	AuthAppInputFieldSourceCodeType = "sourceCodeType"
	AuthAppInputFieldTLS            = "tls"
)

type AuthAppInput struct {
	ClientID       string `json:"clientId,omitempty" yaml:"client_id,omitempty"`
	ClientSecret   string `json:"clientSecret,omitempty" yaml:"client_secret,omitempty"`
	Code           string `json:"code,omitempty" yaml:"code,omitempty"`
	Host           string `json:"host,omitempty" yaml:"host,omitempty"`
	InheritGlobal  bool   `json:"inheritGlobal,omitempty" yaml:"inherit_global,omitempty"`
	RedirectURL    string `json:"redirectUrl,omitempty" yaml:"redirect_url,omitempty"`
	SourceCodeType string `json:"sourceCodeType,omitempty" yaml:"source_code_type,omitempty"`
	TLS            bool   `json:"tls,omitempty" yaml:"tls,omitempty"`
}
