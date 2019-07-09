package client

const (
	ReplicaSetConfigType                 = "replicaSetConfig"
	ReplicaSetConfigFieldMinReadySeconds = "minReadySeconds"
)

type ReplicaSetConfig struct {
	MinReadySeconds int64 `json:"minReadySeconds,omitempty" yaml:"min_ready_seconds,omitempty"`
}
