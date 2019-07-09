package client

const (
	KontainerDriverSpecType                  = "kontainerDriverSpec"
	KontainerDriverSpecFieldActive           = "active"
	KontainerDriverSpecFieldBuiltIn          = "builtIn"
	KontainerDriverSpecFieldChecksum         = "checksum"
	KontainerDriverSpecFieldUIURL            = "uiUrl"
	KontainerDriverSpecFieldURL              = "url"
	KontainerDriverSpecFieldWhitelistDomains = "whitelistDomains"
)

type KontainerDriverSpec struct {
	Active           bool     `json:"active,omitempty" yaml:"active,omitempty"`
	BuiltIn          bool     `json:"builtIn,omitempty" yaml:"built_in,omitempty"`
	Checksum         string   `json:"checksum,omitempty" yaml:"checksum,omitempty"`
	UIURL            string   `json:"uiUrl,omitempty" yaml:"ui_url,omitempty"`
	URL              string   `json:"url,omitempty" yaml:"url,omitempty"`
	WhitelistDomains []string `json:"whitelistDomains,omitempty" yaml:"whitelist_domains,omitempty"`
}
