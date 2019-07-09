package client

const (
	QueueConfigType                   = "queueConfig"
	QueueConfigFieldBatchSendDeadline = "batchSendDeadline"
	QueueConfigFieldCapacity          = "capacity"
	QueueConfigFieldMaxBackoff        = "maxBackoff"
	QueueConfigFieldMaxRetries        = "maxRetries"
	QueueConfigFieldMaxSamplesPerSend = "maxSamplesPerSend"
	QueueConfigFieldMaxShards         = "maxShards"
	QueueConfigFieldMinBackoff        = "minBackoff"
)

type QueueConfig struct {
	BatchSendDeadline string `json:"batchSendDeadline,omitempty" yaml:"batch_send_deadline,omitempty"`
	Capacity          int64  `json:"capacity,omitempty" yaml:"capacity,omitempty"`
	MaxBackoff        string `json:"maxBackoff,omitempty" yaml:"max_backoff,omitempty"`
	MaxRetries        int64  `json:"maxRetries,omitempty" yaml:"max_retries,omitempty"`
	MaxSamplesPerSend int64  `json:"maxSamplesPerSend,omitempty" yaml:"max_samples_per_send,omitempty"`
	MaxShards         int64  `json:"maxShards,omitempty" yaml:"max_shards,omitempty"`
	MinBackoff        string `json:"minBackoff,omitempty" yaml:"min_backoff,omitempty"`
}
