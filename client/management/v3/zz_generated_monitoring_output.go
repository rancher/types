package client

const (
	MonitoringOutputType                  = "monitoringOutput"
	MonitoringOutputFieldAnswers          = "answers"
	MonitoringOutputFieldAnswersSetString = "answersAnswersSetString"
	MonitoringOutputFieldVersion          = "version"
)

type MonitoringOutput struct {
	Answers          map[string]string `json:"answers,omitempty" yaml:"answers,omitempty"`
	AnswersSetString map[string]string `json:"answersAnswersSetString,omitempty" yaml:"answersAnswersSetString,omitempty"`
	Version          string            `json:"version,omitempty" yaml:"version,omitempty"`
}
