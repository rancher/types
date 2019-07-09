package client

const (
	RecipientType              = "recipient"
	RecipientFieldNotifierID   = "notifierId"
	RecipientFieldNotifierType = "notifierType"
	RecipientFieldRecipient    = "recipient"
)

type Recipient struct {
	NotifierID   string `json:"notifierId,omitempty" yaml:"notifier_id,omitempty"`
	NotifierType string `json:"notifierType,omitempty" yaml:"notifier_type,omitempty"`
	Recipient    string `json:"recipient,omitempty" yaml:"recipient,omitempty"`
}
