package client

const (
	GitlabApplyInputType              = "gitlabApplyInput"
	GitlabApplyInputFieldClientID     = "clientId"
	GitlabApplyInputFieldClientSecret = "clientSecret"
	GitlabApplyInputFieldCode         = "code"
	GitlabApplyInputFieldHostname     = "hostname"
	GitlabApplyInputFieldRedirectURL  = "redirectUrl"
	GitlabApplyInputFieldTLS          = "tls"
)

type GitlabApplyInput struct {
	ClientID     string `json:"clientId,omitempty" yaml:"client_id,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty" yaml:"client_secret,omitempty"`
	Code         string `json:"code,omitempty" yaml:"code,omitempty"`
	Hostname     string `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	RedirectURL  string `json:"redirectUrl,omitempty" yaml:"redirect_url,omitempty"`
	TLS          bool   `json:"tls,omitempty" yaml:"tls,omitempty"`
}
