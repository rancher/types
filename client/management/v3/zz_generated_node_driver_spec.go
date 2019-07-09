package client

const (
	NodeDriverSpecType                  = "nodeDriverSpec"
	NodeDriverSpecFieldActive           = "active"
	NodeDriverSpecFieldBuiltin          = "builtin"
	NodeDriverSpecFieldChecksum         = "checksum"
	NodeDriverSpecFieldDescription      = "description"
	NodeDriverSpecFieldDisplayName      = "displayName"
	NodeDriverSpecFieldExternalID       = "externalId"
	NodeDriverSpecFieldUIURL            = "uiUrl"
	NodeDriverSpecFieldURL              = "url"
	NodeDriverSpecFieldWhitelistDomains = "whitelistDomains"
)

type NodeDriverSpec struct {
	Active           bool     `json:"active,omitempty" yaml:"active,omitempty"`
	Builtin          bool     `json:"builtin,omitempty" yaml:"builtin,omitempty"`
	Checksum         string   `json:"checksum,omitempty" yaml:"checksum,omitempty"`
	Description      string   `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName      string   `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	ExternalID       string   `json:"externalId,omitempty" yaml:"external_id,omitempty"`
	UIURL            string   `json:"uiUrl,omitempty" yaml:"ui_url,omitempty"`
	URL              string   `json:"url,omitempty" yaml:"url,omitempty"`
	WhitelistDomains []string `json:"whitelistDomains,omitempty" yaml:"whitelist_domains,omitempty"`
}
