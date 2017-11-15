package client

const (
	ObjectMetaType                   = "objectMeta"
	ObjectMetaFieldAnnotations       = "annotations"
	ObjectMetaFieldCreationTimestamp = "creationTimestamp"
	ObjectMetaFieldDeletionTimestamp = "deletionTimestamp"
	ObjectMetaFieldLabels            = "labels"
	ObjectMetaFieldName              = "name"
	ObjectMetaFieldNamespace         = "namespace"
	ObjectMetaFieldUID               = "uid"
)

type ObjectMeta struct {
	Annotations       map[string]string `json:"annotations,omitempty"`
	CreationTimestamp string            `json:"creationTimestamp,omitempty"`
	DeletionTimestamp string            `json:"deletionTimestamp,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	Name              string            `json:"name,omitempty"`
	Namespace         string            `json:"namespace,omitempty"`
	UID               string            `json:"uid,omitempty"`
}
