package client

const (
	ProjectLoggingSpecType                       = "projectLoggingSpec"
	ProjectLoggingSpecFieldCustomTargetConfig    = "customTargetConfig"
	ProjectLoggingSpecFieldDisplayName           = "displayName"
	ProjectLoggingSpecFieldElasticsearchConfig   = "elasticsearchConfig"
	ProjectLoggingSpecFieldFluentForwarderConfig = "fluentForwarderConfig"
	ProjectLoggingSpecFieldKafkaConfig           = "kafkaConfig"
	ProjectLoggingSpecFieldOutputFlushInterval   = "outputFlushInterval"
	ProjectLoggingSpecFieldOutputTags            = "outputTags"
	ProjectLoggingSpecFieldProjectID             = "projectId"
	ProjectLoggingSpecFieldSplunkConfig          = "splunkConfig"
	ProjectLoggingSpecFieldSyslogConfig          = "syslogConfig"
)

type ProjectLoggingSpec struct {
	CustomTargetConfig    *CustomTargetConfig    `json:"customTargetConfig,omitempty" yaml:"custom_target_config,omitempty"`
	DisplayName           string                 `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	ElasticsearchConfig   *ElasticsearchConfig   `json:"elasticsearchConfig,omitempty" yaml:"elasticsearch_config,omitempty"`
	FluentForwarderConfig *FluentForwarderConfig `json:"fluentForwarderConfig,omitempty" yaml:"fluent_forwarder_config,omitempty"`
	KafkaConfig           *KafkaConfig           `json:"kafkaConfig,omitempty" yaml:"kafka_config,omitempty"`
	OutputFlushInterval   int64                  `json:"outputFlushInterval,omitempty" yaml:"output_flush_interval,omitempty"`
	OutputTags            map[string]string      `json:"outputTags,omitempty" yaml:"output_tags,omitempty"`
	ProjectID             string                 `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	SplunkConfig          *SplunkConfig          `json:"splunkConfig,omitempty" yaml:"splunk_config,omitempty"`
	SyslogConfig          *SyslogConfig          `json:"syslogConfig,omitempty" yaml:"syslog_config,omitempty"`
}
