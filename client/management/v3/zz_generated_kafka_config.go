package client

const (
	KafkaConfigType                    = "kafkaConfig"
	KafkaConfigFieldBrokerEndpoints    = "brokerEndpoints"
	KafkaConfigFieldCertificate        = "certificate"
	KafkaConfigFieldClientCert         = "clientCert"
	KafkaConfigFieldClientKey          = "clientKey"
	KafkaConfigFieldSaslPassword       = "saslPassword"
	KafkaConfigFieldSaslScramMechanism = "saslScramMechanism"
	KafkaConfigFieldSaslType           = "saslType"
	KafkaConfigFieldSaslUsername       = "saslUsername"
	KafkaConfigFieldTopic              = "topic"
	KafkaConfigFieldZookeeperEndpoint  = "zookeeperEndpoint"
)

type KafkaConfig struct {
	BrokerEndpoints    []string `json:"brokerEndpoints,omitempty" yaml:"broker_endpoints,omitempty"`
	Certificate        string   `json:"certificate,omitempty" yaml:"certificate,omitempty"`
	ClientCert         string   `json:"clientCert,omitempty" yaml:"client_cert,omitempty"`
	ClientKey          string   `json:"clientKey,omitempty" yaml:"client_key,omitempty"`
	SaslPassword       string   `json:"saslPassword,omitempty" yaml:"sasl_password,omitempty"`
	SaslScramMechanism string   `json:"saslScramMechanism,omitempty" yaml:"sasl_scram_mechanism,omitempty"`
	SaslType           string   `json:"saslType,omitempty" yaml:"sasl_type,omitempty"`
	SaslUsername       string   `json:"saslUsername,omitempty" yaml:"sasl_username,omitempty"`
	Topic              string   `json:"topic,omitempty" yaml:"topic,omitempty"`
	ZookeeperEndpoint  string   `json:"zookeeperEndpoint,omitempty" yaml:"zookeeper_endpoint,omitempty"`
}
