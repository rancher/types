package client

const (
	TopologySelectorTermType                       = "topologySelectorTerm"
	TopologySelectorTermFieldMatchLabelExpressions = "matchLabelExpressions"
)

type TopologySelectorTerm struct {
	MatchLabelExpressions []TopologySelectorLabelRequirement `json:"matchLabelExpressions,omitempty" yaml:"match_label_expressions,omitempty"`
}
