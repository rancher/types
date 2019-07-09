package client

const (
	EnvFromType            = "envFrom"
	EnvFromFieldSourceKey  = "sourceKey"
	EnvFromFieldSourceName = "sourceName"
	EnvFromFieldTargetKey  = "targetKey"
)

type EnvFrom struct {
	SourceKey  string `json:"sourceKey,omitempty" yaml:"source_key,omitempty"`
	SourceName string `json:"sourceName,omitempty" yaml:"source_name,omitempty"`
	TargetKey  string `json:"targetKey,omitempty" yaml:"target_key,omitempty"`
}
