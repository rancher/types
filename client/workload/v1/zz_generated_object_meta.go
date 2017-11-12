package client

const (
	ObjectMetaType                            = "objectMeta"
	ObjectMetaFieldAnnotations                = "annotations"
	ObjectMetaFieldCreated                    = "created"
	ObjectMetaFieldDeletionGracePeriodSeconds = "deletionGracePeriodSeconds"
	ObjectMetaFieldLabels                     = "labels"
	ObjectMetaFieldName                       = "name"
	ObjectMetaFieldNamespace                  = "namespace"
	ObjectMetaFieldRemoved                    = "removed"
	ObjectMetaFieldUuid                       = "uuid"
)

type ObjectMeta struct {
	Annotations                map[string]string `json:"annotations,omitempty"`
	Created                    string            `json:"created,omitempty"`
	DeletionGracePeriodSeconds *int64            `json:"deletionGracePeriodSeconds,omitempty"`
	Labels                     map[string]string `json:"labels,omitempty"`
	Name                       string            `json:"name,omitempty"`
	Namespace                  string            `json:"namespace,omitempty"`
	Removed                    string            `json:"removed,omitempty"`
	Uuid                       string            `json:"uuid,omitempty"`
}
