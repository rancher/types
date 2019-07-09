package client

const (
	SamlConfigTestInputType                  = "samlConfigTestInput"
	SamlConfigTestInputFieldFinalRedirectURL = "finalRedirectUrl"
)

type SamlConfigTestInput struct {
	FinalRedirectURL string `json:"finalRedirectUrl,omitempty" yaml:"final_redirect_url,omitempty"`
}
