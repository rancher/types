package client

const (
	NodePoolSpecType                         = "nodePoolSpec"
	NodePoolSpecFieldClusterID               = "clusterId"
	NodePoolSpecFieldControlPlane            = "controlPlane"
	NodePoolSpecFieldDeleteNotReadyAfterSecs = "deleteNotReadyAfterSecs"
	NodePoolSpecFieldDisplayName             = "displayName"
	NodePoolSpecFieldEtcd                    = "etcd"
	NodePoolSpecFieldHostnamePrefix          = "hostnamePrefix"
	NodePoolSpecFieldNodeAnnotations         = "nodeAnnotations"
	NodePoolSpecFieldNodeLabels              = "nodeLabels"
	NodePoolSpecFieldNodeTemplateID          = "nodeTemplateId"
	NodePoolSpecFieldQuantity                = "quantity"
	NodePoolSpecFieldWorker                  = "worker"
)

type NodePoolSpec struct {
	ClusterID               string            `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`
	ControlPlane            bool              `json:"controlPlane,omitempty" yaml:"control_plane,omitempty"`
	DeleteNotReadyAfterSecs int64             `json:"deleteNotReadyAfterSecs,omitempty" yaml:"delete_not_ready_after_secs,omitempty"`
	DisplayName             string            `json:"displayName,omitempty" yaml:"display_name,omitempty"`
	Etcd                    bool              `json:"etcd,omitempty" yaml:"etcd,omitempty"`
	HostnamePrefix          string            `json:"hostnamePrefix,omitempty" yaml:"hostname_prefix,omitempty"`
	NodeAnnotations         map[string]string `json:"nodeAnnotations,omitempty" yaml:"node_annotations,omitempty"`
	NodeLabels              map[string]string `json:"nodeLabels,omitempty" yaml:"node_labels,omitempty"`
	NodeTemplateID          string            `json:"nodeTemplateId,omitempty" yaml:"node_template_id,omitempty"`
	Quantity                int64             `json:"quantity,omitempty" yaml:"quantity,omitempty"`
	Worker                  bool              `json:"worker,omitempty" yaml:"worker,omitempty"`
}
