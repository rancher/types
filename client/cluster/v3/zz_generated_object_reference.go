package client

const (
	ObjectReferenceType                 = "objectReference"
	ObjectReferenceFieldAPIVersion      = "apiVersion"
	ObjectReferenceFieldFieldPath       = "fieldPath"
	ObjectReferenceFieldKind            = "kind"
	ObjectReferenceFieldName            = "name"
	ObjectReferenceFieldNamespace       = "namespace"
	ObjectReferenceFieldResourceVersion = "resourceVersion"
	ObjectReferenceFieldUID             = "uid"
)

type ObjectReference struct {
	APIVersion      string `json:"apiVersion,omitempty" yaml:"api_version,omitempty"`
	FieldPath       string `json:"fieldPath,omitempty" yaml:"field_path,omitempty"`
	Kind            string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name            string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace       string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	ResourceVersion string `json:"resourceVersion,omitempty" yaml:"resource_version,omitempty"`
	UID             string `json:"uid,omitempty" yaml:"uid,omitempty"`
}
