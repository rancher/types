package client

const (
	GlobalOpenstackOptsType                             = "globalOpenstackOpts"
	GlobalOpenstackOptsFieldApplicationCredentialID     = "application-credential-id"
	GlobalOpenstackOptsFieldApplicationCredentialName   = "application-credential-name"
	GlobalOpenstackOptsFieldApplicationCredentialSecret = "application-credential-secret"
	GlobalOpenstackOptsFieldAuthURL                     = "auth-url"
	GlobalOpenstackOptsFieldCAFile                      = "ca-file"
	GlobalOpenstackOptsFieldDomainID                    = "domain-id"
	GlobalOpenstackOptsFieldDomainName                  = "domain-name"
	GlobalOpenstackOptsFieldPassword                    = "password"
	GlobalOpenstackOptsFieldRegion                      = "region"
	GlobalOpenstackOptsFieldTenantDomainID              = "tenant-domain-id"
	GlobalOpenstackOptsFieldTenantDomainName            = "tenant-domain-name"
	GlobalOpenstackOptsFieldTenantID                    = "tenant-id"
	GlobalOpenstackOptsFieldTenantName                  = "tenant-name"
	GlobalOpenstackOptsFieldTrustID                     = "trust-id"
	GlobalOpenstackOptsFieldUserDomainID                = "user-domain-id"
	GlobalOpenstackOptsFieldUserDomainName              = "user-domain-name"
	GlobalOpenstackOptsFieldUserID                      = "user-id"
	GlobalOpenstackOptsFieldUsername                    = "username"
)

type GlobalOpenstackOpts struct {
	ApplicationCredentialID     string `json:"application-credential-id,omitempty" yaml:"application-credential-id,omitempty"`
	ApplicationCredentialName   string `json:"application-credential-name,omitempty" yaml:"application-credential-name,omitempty"`
	ApplicationCredentialSecret string `json:"application-credential-secret,omitempty" yaml:"application-credential-secret,omitempty"`
	AuthURL                     string `json:"auth-url,omitempty" yaml:"auth-url,omitempty"`
	CAFile                      string `json:"ca-file,omitempty" yaml:"ca-file,omitempty"`
	DomainID                    string `json:"domain-id,omitempty" yaml:"domain-id,omitempty"`
	DomainName                  string `json:"domain-name,omitempty" yaml:"domain-name,omitempty"`
	Password                    string `json:"password,omitempty" yaml:"password,omitempty"`
	Region                      string `json:"region,omitempty" yaml:"region,omitempty"`
	TenantDomainID              string `json:"tenant-domain-id,omitempty" yaml:"tenant-domain-id,omitempty"`
	TenantDomainName            string `json:"tenant-domain-name,omitempty" yaml:"tenant-domain-name,omitempty"`
	TenantID                    string `json:"tenant-id,omitempty" yaml:"tenant-id,omitempty"`
	TenantName                  string `json:"tenant-name,omitempty" yaml:"tenant-name,omitempty"`
	TrustID                     string `json:"trust-id,omitempty" yaml:"trust-id,omitempty"`
	UserDomainID                string `json:"user-domain-id,omitempty" yaml:"user-domain-id,omitempty"`
	UserDomainName              string `json:"user-domain-name,omitempty" yaml:"user-domain-name,omitempty"`
	UserID                      string `json:"user-id,omitempty" yaml:"user-id,omitempty"`
	Username                    string `json:"username,omitempty" yaml:"username,omitempty"`
}
