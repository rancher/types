package client

const (
	ObjectMetaType                            = "objectMeta"
	ObjectMetaFieldAnnotations                = "annotations"
	ObjectMetaFieldClusterName                = "clusterName"
	ObjectMetaFieldCreationTimestamp          = "creationTimestamp"
	ObjectMetaFieldDeletionGracePeriodSeconds = "deletionGracePeriodSeconds"
	ObjectMetaFieldDeletionTimestamp          = "deletionTimestamp"
	ObjectMetaFieldFinalizers                 = "finalizers"
	ObjectMetaFieldGenerateName               = "generateName"
	ObjectMetaFieldGeneration                 = "generation"
	ObjectMetaFieldInitializers               = "initializers"
	ObjectMetaFieldLabels                     = "labels"
	ObjectMetaFieldName                       = "name"
	ObjectMetaFieldNamespace                  = "namespace"
	ObjectMetaFieldOwnerReferences            = "ownerReferences"
	ObjectMetaFieldResourceVersion            = "resourceVersion"
	ObjectMetaFieldSelfLink                   = "selfLink"
	ObjectMetaFieldUID                        = "uid"
)

type ObjectMeta struct {
	Annotations                map[string]string `json:"annotations,omitempty"`
	ClusterName                string            `json:"clusterName,omitempty"`
	CreationTimestamp          string            `json:"creationTimestamp,omitempty"`
	DeletionGracePeriodSeconds *int64            `json:"deletionGracePeriodSeconds,omitempty"`
	DeletionTimestamp          string            `json:"deletionTimestamp,omitempty"`
	Finalizers                 []string          `json:"finalizers,omitempty"`
	GenerateName               string            `json:"generateName,omitempty"`
	Generation                 *int64            `json:"generation,omitempty"`
	Initializers               *Initializers     `json:"initializers,omitempty"`
	Labels                     map[string]string `json:"labels,omitempty"`
	Name                       string            `json:"name,omitempty"`
	Namespace                  string            `json:"namespace,omitempty"`
	OwnerReferences            []OwnerReference  `json:"ownerReferences,omitempty"`
	ResourceVersion            string            `json:"resourceVersion,omitempty"`
	SelfLink                   string            `json:"selfLink,omitempty"`
	UID                        string            `json:"uid,omitempty"`
}
