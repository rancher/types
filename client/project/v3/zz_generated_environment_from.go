package client

const (
	EnvironmentFromType            = "environmentFrom"
	EnvironmentFromFieldOptional   = "optional"
	EnvironmentFromFieldPrefix     = "prefix"
	EnvironmentFromFieldSource     = "source"
	EnvironmentFromFieldSourceKey  = "sourceKey"
	EnvironmentFromFieldSourceName = "sourceName"
	EnvironmentFromFieldTargetKey  = "targetKey"
)

type EnvironmentFrom struct {
	Optional   bool   `json:"optional,omitempty" yaml:"optional,omitempty"`
	Prefix     string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Source     string `json:"source,omitempty" yaml:"source,omitempty"`
	SourceKey  string `json:"sourceKey,omitempty" yaml:"source_key,omitempty"`
	SourceName string `json:"sourceName,omitempty" yaml:"source_name,omitempty"`
	TargetKey  string `json:"targetKey,omitempty" yaml:"target_key,omitempty"`
}
