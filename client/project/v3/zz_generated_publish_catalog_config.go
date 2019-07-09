package client

const (
	PublishCatalogConfigType                 = "publishCatalogConfig"
	PublishCatalogConfigFieldCatalogTemplate = "catalogTemplate"
	PublishCatalogConfigFieldGitAuthor       = "gitAuthor"
	PublishCatalogConfigFieldGitBranch       = "gitBranch"
	PublishCatalogConfigFieldGitEmail        = "gitEmail"
	PublishCatalogConfigFieldGitURL          = "gitUrl"
	PublishCatalogConfigFieldPath            = "path"
	PublishCatalogConfigFieldVersion         = "version"
)

type PublishCatalogConfig struct {
	CatalogTemplate string `json:"catalogTemplate,omitempty" yaml:"catalog_template,omitempty"`
	GitAuthor       string `json:"gitAuthor,omitempty" yaml:"git_author,omitempty"`
	GitBranch       string `json:"gitBranch,omitempty" yaml:"git_branch,omitempty"`
	GitEmail        string `json:"gitEmail,omitempty" yaml:"git_email,omitempty"`
	GitURL          string `json:"gitUrl,omitempty" yaml:"git_url,omitempty"`
	Path            string `json:"path,omitempty" yaml:"path,omitempty"`
	Version         string `json:"version,omitempty" yaml:"version,omitempty"`
}
