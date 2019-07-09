package client

const (
	MultiClusterAppSpecType                      = "multiClusterAppSpec"
	MultiClusterAppSpecFieldAnswers              = "answers"
	MultiClusterAppSpecFieldMembers              = "members"
	MultiClusterAppSpecFieldRevisionHistoryLimit = "revisionHistoryLimit"
	MultiClusterAppSpecFieldRoles                = "roles"
	MultiClusterAppSpecFieldTargets              = "targets"
	MultiClusterAppSpecFieldTemplateVersionID    = "templateVersionId"
	MultiClusterAppSpecFieldUpgradeStrategy      = "upgradeStrategy"
)

type MultiClusterAppSpec struct {
	Answers              []Answer         `json:"answers,omitempty" yaml:"answers,omitempty"`
	Members              []Member         `json:"members,omitempty" yaml:"members,omitempty"`
	RevisionHistoryLimit int64            `json:"revisionHistoryLimit,omitempty" yaml:"revision_history_limit,omitempty"`
	Roles                []string         `json:"roles,omitempty" yaml:"roles,omitempty"`
	Targets              []Target         `json:"targets,omitempty" yaml:"targets,omitempty"`
	TemplateVersionID    string           `json:"templateVersionId,omitempty" yaml:"template_version_id,omitempty"`
	UpgradeStrategy      *UpgradeStrategy `json:"upgradeStrategy,omitempty" yaml:"upgrade_strategy,omitempty"`
}
