package client

const (
	SamlConfigTestOutputType                = "samlConfigTestOutput"
	SamlConfigTestOutputFieldIdpRedirectURL = "idpRedirectUrl"
)

type SamlConfigTestOutput struct {
	IdpRedirectURL string `json:"idpRedirectUrl,omitempty" yaml:"idp_redirect_url,omitempty"`
}
