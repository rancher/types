package client

const (
	HelmChartReferenceType               = "helmChartReference"
	HelmChartReferenceFieldReference     = "reference"
	HelmChartReferenceFieldRepositoryURL = "repositoryUrl"
	HelmChartReferenceFieldVersion       = "version"
)

type HelmChartReference struct {
	Reference     string `json:"reference,omitempty" yaml:"reference,omitempty"`
	RepositoryURL string `json:"repositoryUrl,omitempty" yaml:"repositoryUrl,omitempty"`
	Version       string `json:"version,omitempty" yaml:"version,omitempty"`
}
