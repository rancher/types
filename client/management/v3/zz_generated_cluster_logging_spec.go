package client

const (
	ClusterLoggingSpecType                        = "clusterLoggingSpec"
	ClusterLoggingSpecFieldClusterID              = "clusterId"
	ClusterLoggingSpecFieldCustomTargetConfig     = "customTargetConfig"
	ClusterLoggingSpecFieldDisplayName            = "displayName"
	ClusterLoggingSpecFieldElasticsearchConfig    = "elasticsearchConfig"
	ClusterLoggingSpecFieldFluentForwarderConfig  = "fluentForwarderConfig"
	ClusterLoggingSpecFieldIncludeSystemComponent = "includeSystemComponent"
	ClusterLoggingSpecFieldKafkaConfig            = "kafkaConfig"
	ClusterLoggingSpecFieldOutputFlushInterval    = "outputFlushInterval"
	ClusterLoggingSpecFieldOutputTags             = "outputTags"
	ClusterLoggingSpecFieldSplunkConfig           = "splunkConfig"
	ClusterLoggingSpecFieldSyslogConfig           = "syslogConfig"
)

type ClusterLoggingSpec struct {
	ClusterID              string                 `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`
	CustomTargetConfig     *CustomTargetConfig    `json:"customTargetConfig,omitempty" yaml:"custom_target_config,omitempty"`
	DisplayName            string                 `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	ElasticsearchConfig    *ElasticsearchConfig   `json:"elasticsearchConfig,omitempty" yaml:"elasticsearch_config,omitempty"`
	FluentForwarderConfig  *FluentForwarderConfig `json:"fluentForwarderConfig,omitempty" yaml:"fluent_forwarder_config,omitempty"`
	IncludeSystemComponent *bool                  `json:"includeSystemComponent,omitempty" yaml:"include_system_component,omitempty"`
	KafkaConfig            *KafkaConfig           `json:"kafkaConfig,omitempty" yaml:"kafka_config,omitempty"`
	OutputFlushInterval    int64                  `json:"outputFlushInterval,omitempty" yaml:"output_flush_interval,omitempty"`
	OutputTags             map[string]string      `json:"outputTags,omitempty" yaml:"output_tags,omitempty"`
	SplunkConfig           *SplunkConfig          `json:"splunkConfig,omitempty" yaml:"splunk_config,omitempty"`
	SyslogConfig           *SyslogConfig          `json:"syslogConfig,omitempty" yaml:"syslog_config,omitempty"`
}
