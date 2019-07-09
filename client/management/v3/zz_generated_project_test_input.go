package client

const (
	ProjectTestInputType                       = "projectTestInput"
	ProjectTestInputFieldCustomTargetConfig    = "customTargetConfig"
	ProjectTestInputFieldElasticsearchConfig   = "elasticsearchConfig"
	ProjectTestInputFieldFluentForwarderConfig = "fluentForwarderConfig"
	ProjectTestInputFieldKafkaConfig           = "kafkaConfig"
	ProjectTestInputFieldOutputTags            = "outputTags"
	ProjectTestInputFieldProjectName           = "projectId"
	ProjectTestInputFieldSplunkConfig          = "splunkConfig"
	ProjectTestInputFieldSyslogConfig          = "syslogConfig"
)

type ProjectTestInput struct {
	CustomTargetConfig    *CustomTargetConfig    `json:"customTargetConfig,omitempty" yaml:"custom_target_config,omitempty"`
	ElasticsearchConfig   *ElasticsearchConfig   `json:"elasticsearchConfig,omitempty" yaml:"elasticsearch_config,omitempty"`
	FluentForwarderConfig *FluentForwarderConfig `json:"fluentForwarderConfig,omitempty" yaml:"fluent_forwarder_config,omitempty"`
	KafkaConfig           *KafkaConfig           `json:"kafkaConfig,omitempty" yaml:"kafka_config,omitempty"`
	OutputTags            map[string]string      `json:"outputTags,omitempty" yaml:"output_tags,omitempty"`
	ProjectName           string                 `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	SplunkConfig          *SplunkConfig          `json:"splunkConfig,omitempty" yaml:"splunk_config,omitempty"`
	SyslogConfig          *SyslogConfig          `json:"syslogConfig,omitempty" yaml:"syslog_config,omitempty"`
}
