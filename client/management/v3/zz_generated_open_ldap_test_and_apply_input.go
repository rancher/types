package client

const (
	OpenLdapTestAndApplyInputType            = "openLdapTestAndApplyInput"
	OpenLdapTestAndApplyInputFieldLdapConfig = "ldapConfig"
	OpenLdapTestAndApplyInputFieldPassword   = "password"
	OpenLdapTestAndApplyInputFieldUsername   = "username"
)

type OpenLdapTestAndApplyInput struct {
	LdapConfig *LdapConfig `json:"ldapConfig,omitempty" yaml:"ldap_config,omitempty"`
	Password   string      `json:"password,omitempty" yaml:"password,omitempty"`
	Username   string      `json:"username,omitempty" yaml:"username,omitempty"`
}
