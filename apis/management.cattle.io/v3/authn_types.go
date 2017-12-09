package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Token struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	TokenID          string `json:"tokenID,omitempty"`
	TokenValue       string `json:"tokenValue,omitempty"`
	User             string `json:"user,omitempty"`
	ExternalID       string `json:"externalID,omitempty"`
	AuthProvider     string `json:"authProvider,omitempty"`
	TTLMillis        string `json:"ttl,omitempty"`
	RefreshTTLMillis string `json:"refreshTTL,omitempty"`
	LastUpdateTime   string `json:"lastUpdateTime,omitempty"`
	IsDerived        bool   `json:"isDerived,omitempty"`
	Description      string `json:"description,omitempty"`
}

type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	UserName    string `json:"userName,omitempty"`
	Secret      string `json:"secret,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	ExternalID  string `json:"externalID,omitempty"`
}

type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	GroupName   string `json:"groupName,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	ExternalID  string `json:"externalID,omitempty"`
}

type GroupMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	GroupExternalID  string `json:"groupExternalID,omitempty"`
	MemberExternalID string `json:"memberExternalID,omitempty"`
}

type Identity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	ExternalID     string            `json:"externalID,omitempty"`
	DisplayName    string            `json:"displayName,omitempty"`
	IdentityName   string            `json:"identityName,omitempty"`
	ProfilePicture string            `json:"profilePicture,omitempty"`
	ProfileURL     string            `json:"profileURL,omitempty"`
	Kind           string            `json:"kind,omitempty"`
	Me             bool              `json:"me,omitempty"`
	MemberOf       bool              `json:"memberOf,omitempty"`
	ExtraInfo      map[string]string `json:"extraInfo,omitempty"`
}
