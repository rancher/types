package client

const (
	ReplicationControllerSpecType                       = "replicationControllerSpec"
	ReplicationControllerSpecFieldReplicationController = "replicationController"
	ReplicationControllerSpecFieldSelector              = "selector"
	ReplicationControllerSpecFieldTemplate              = "template"
)

type ReplicationControllerSpec struct {
	ReplicationController *ReplicationControllerConfig `json:"replicationController,omitempty"`
	Selector              map[string]string            `json:"selector,omitempty"`
	Template              *PodTemplateSpec             `json:"template,omitempty"`
}
