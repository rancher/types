package client

const (
	ObjectFieldSelectorType            = "objectFieldSelector"
	ObjectFieldSelectorFieldAPIVersion = "apiVersion"
	ObjectFieldSelectorFieldFieldPath  = "fieldPath"
)

type ObjectFieldSelector struct {
	APIVersion string `json:"apiVersion,omitempty" yaml:"api_version,omitempty"`
	FieldPath  string `json:"fieldPath,omitempty" yaml:"field_path,omitempty"`
}
