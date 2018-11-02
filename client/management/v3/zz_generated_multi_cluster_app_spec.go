package client

const (
	MultiClusterAppSpecType                  = "multiClusterAppSpec"
	MultiClusterAppSpecFieldAnswers          = "answers"
	MultiClusterAppSpecFieldHelmChart        = "helmChart"
	MultiClusterAppSpecFieldReleaseName      = "releaseName"
	MultiClusterAppSpecFieldReleaseNamespace = "releaseNamespace"
	MultiClusterAppSpecFieldTargets          = "targets"
	MultiClusterAppSpecFieldTemplate         = "template"
)

type MultiClusterAppSpec struct {
	Answers          map[string]string   `json:"answers,omitempty" yaml:"answers,omitempty"`
	HelmChart        *HelmChartReference `json:"helmChart,omitempty" yaml:"helmChart,omitempty"`
	ReleaseName      string              `json:"releaseName,omitempty" yaml:"releaseName,omitempty"`
	ReleaseNamespace string              `json:"releaseNamespace,omitempty" yaml:"releaseNamespace,omitempty"`
	Targets          []Target            `json:"targets,omitempty" yaml:"targets,omitempty"`
	Template         *TemplateReference  `json:"template,omitempty" yaml:"template,omitempty"`
}
