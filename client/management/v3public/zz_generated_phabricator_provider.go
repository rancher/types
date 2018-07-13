package client

const (
	PhabricatorProviderType                 = "phabricatorProvider"
	PhabricatorProviderFieldAnnotations     = "annotations"
	PhabricatorProviderFieldCreated         = "created"
	PhabricatorProviderFieldCreatorID       = "creatorId"
	PhabricatorProviderFieldLabels          = "labels"
	PhabricatorProviderFieldName            = "name"
	PhabricatorProviderFieldOwnerReferences = "ownerReferences"
	PhabricatorProviderFieldRedirectURL     = "redirectUrl"
	PhabricatorProviderFieldRemoved         = "removed"
	PhabricatorProviderFieldType            = "type"
	PhabricatorProviderFieldUuid            = "uuid"
)

type PhabricatorProvider struct {
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	RedirectURL     string            `json:"redirectUrl,omitempty" yaml:"redirectUrl,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	Type            string            `json:"type,omitempty" yaml:"type,omitempty"`
	Uuid            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
