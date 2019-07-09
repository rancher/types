package client

const (
	ExportOutputType            = "exportOutput"
	ExportOutputFieldYAMLOutput = "yamlOutput"
)

type ExportOutput struct {
	YAMLOutput string `json:"yamlOutput,omitempty" yaml:"yaml_output,omitempty"`
}
