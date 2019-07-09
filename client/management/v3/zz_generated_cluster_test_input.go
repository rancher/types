package client

const (
	ClusterTestInputType                       = "clusterTestInput"
	ClusterTestInputFieldClusterName           = "clusterId"
	ClusterTestInputFieldCustomTargetConfig    = "customTargetConfig"
	ClusterTestInputFieldElasticsearchConfig   = "elasticsearchConfig"
	ClusterTestInputFieldFluentForwarderConfig = "fluentForwarderConfig"
	ClusterTestInputFieldKafkaConfig           = "kafkaConfig"
	ClusterTestInputFieldOutputTags            = "outputTags"
	ClusterTestInputFieldSplunkConfig          = "splunkConfig"
	ClusterTestInputFieldSyslogConfig          = "syslogConfig"
)

type ClusterTestInput struct {
	ClusterName           string                 `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`
	CustomTargetConfig    *CustomTargetConfig    `json:"customTargetConfig,omitempty" yaml:"custom_target_config,omitempty"`
	ElasticsearchConfig   *ElasticsearchConfig   `json:"elasticsearchConfig,omitempty" yaml:"elasticsearch_config,omitempty"`
	FluentForwarderConfig *FluentForwarderConfig `json:"fluentForwarderConfig,omitempty" yaml:"fluent_forwarder_config,omitempty"`
	KafkaConfig           *KafkaConfig           `json:"kafkaConfig,omitempty" yaml:"kafka_config,omitempty"`
	OutputTags            map[string]string      `json:"outputTags,omitempty" yaml:"output_tags,omitempty"`
	SplunkConfig          *SplunkConfig          `json:"splunkConfig,omitempty" yaml:"splunk_config,omitempty"`
	SyslogConfig          *SyslogConfig          `json:"syslogConfig,omitempty" yaml:"syslog_config,omitempty"`
}
