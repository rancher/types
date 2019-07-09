package client

const (
	ServiceOverrideType               = "serviceOverride"
	ServiceOverrideFieldRegion        = "region"
	ServiceOverrideFieldService       = "service"
	ServiceOverrideFieldSigningMethod = "signing-method"
	ServiceOverrideFieldSigningName   = "signing-name"
	ServiceOverrideFieldSigningRegion = "signing-region"
	ServiceOverrideFieldURL           = "url"
)

type ServiceOverride struct {
	Region        string `json:"region,omitempty" yaml:"region,omitempty"`
	Service       string `json:"service,omitempty" yaml:"service,omitempty"`
	SigningMethod string `json:"signing-method,omitempty" yaml:"signing_method,omitempty"`
	SigningName   string `json:"signing-name,omitempty" yaml:"signing_name,omitempty"`
	SigningRegion string `json:"signing-region,omitempty" yaml:"signing_region,omitempty"`
	URL           string `json:"url,omitempty" yaml:"url,omitempty"`
}
