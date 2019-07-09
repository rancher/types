package client

const (
	ClusterTemplateRevisionSpecType                   = "clusterTemplateRevisionSpec"
	ClusterTemplateRevisionSpecFieldClusterConfig     = "clusterConfig"
	ClusterTemplateRevisionSpecFieldClusterTemplateID = "clusterTemplateId"
	ClusterTemplateRevisionSpecFieldEnabled           = "enabled"
	ClusterTemplateRevisionSpecFieldQuestions         = "questions"
)

type ClusterTemplateRevisionSpec struct {
	ClusterConfig     *ClusterSpecBase `json:"clusterConfig,omitempty" yaml:"cluster_config,omitempty"`
	ClusterTemplateID string           `json:"clusterTemplateId,omitempty" yaml:"cluster_template_id,omitempty"`
	Enabled           *bool            `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Questions         []Question       `json:"questions,omitempty" yaml:"questions,omitempty"`
}
