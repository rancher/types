package client

const (
	RollingUpdateType           = "rollingUpdate"
	RollingUpdateFieldBatchSize = "batchSize"
	RollingUpdateFieldInterval  = "interval"
)

type RollingUpdate struct {
	BatchSize int64 `json:"batchSize,omitempty" yaml:"batch_size,omitempty"`
	Interval  int64 `json:"interval,omitempty" yaml:"interval,omitempty"`
}
