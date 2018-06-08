package client

const (
	OpenLDAPTestAndApplyInputType                = "openLDAPTestAndApplyInput"
	OpenLDAPTestAndApplyInputFieldOpenLDAPConfig = "openLDAPConfig"
	OpenLDAPTestAndApplyInputFieldPassword       = "password"
	OpenLDAPTestAndApplyInputFieldUsername       = "username"
)

type OpenLDAPTestAndApplyInput struct {
	OpenLDAPConfig *OpenLDAPConfig `json:"openLDAPConfig,omitempty" yaml:"openLDAPConfig,omitempty"`
	Password       string          `json:"password,omitempty" yaml:"password,omitempty"`
	Username       string          `json:"username,omitempty" yaml:"username,omitempty"`
}
