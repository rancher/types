package v1

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
	IsDerived        bool   `json:"IsDerived,omitempty"`
}
