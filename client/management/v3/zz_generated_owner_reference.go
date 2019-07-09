package client

const (
	OwnerReferenceType                    = "ownerReference"
	OwnerReferenceFieldAPIVersion         = "apiVersion"
	OwnerReferenceFieldBlockOwnerDeletion = "blockOwnerDeletion"
	OwnerReferenceFieldController         = "controller"
	OwnerReferenceFieldKind               = "kind"
	OwnerReferenceFieldName               = "name"
	OwnerReferenceFieldUID                = "uid"
)

type OwnerReference struct {
	APIVersion         string `json:"apiVersion,omitempty" yaml:"api_version,omitempty"`
	BlockOwnerDeletion *bool  `json:"blockOwnerDeletion,omitempty" yaml:"block_owner_deletion,omitempty"`
	Controller         *bool  `json:"controller,omitempty" yaml:"controller,omitempty"`
	Kind               string `json:"kind,omitempty" yaml:"kind,omitempty"`
	Name               string `json:"name,omitempty" yaml:"name,omitempty"`
	UID                string `json:"uid,omitempty" yaml:"uid,omitempty"`
}
