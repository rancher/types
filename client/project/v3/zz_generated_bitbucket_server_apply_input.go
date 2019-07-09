package client

const (
	BitbucketServerApplyInputType               = "bitbucketServerApplyInput"
	BitbucketServerApplyInputFieldHostname      = "hostname"
	BitbucketServerApplyInputFieldOAuthToken    = "oauthToken"
	BitbucketServerApplyInputFieldOAuthVerifier = "oauthVerifier"
	BitbucketServerApplyInputFieldRedirectURL   = "redirectUrl"
	BitbucketServerApplyInputFieldTLS           = "tls"
)

type BitbucketServerApplyInput struct {
	Hostname      string `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	OAuthToken    string `json:"oauthToken,omitempty" yaml:"oauth_token,omitempty"`
	OAuthVerifier string `json:"oauthVerifier,omitempty" yaml:"oauth_verifier,omitempty"`
	RedirectURL   string `json:"redirectUrl,omitempty" yaml:"redirect_url,omitempty"`
	TLS           bool   `json:"tls,omitempty" yaml:"tls,omitempty"`
}
