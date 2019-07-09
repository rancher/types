package client

const (
	SamlLoginOutputType                = "samlLoginOutput"
	SamlLoginOutputFieldIdpRedirectURL = "idpRedirectUrl"
)

type SamlLoginOutput struct {
	IdpRedirectURL string `json:"idpRedirectUrl,omitempty" yaml:"idp_redirect_url,omitempty"`
}
