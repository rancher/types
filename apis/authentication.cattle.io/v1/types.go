package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Token struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	TokenKey   string `json:"tokenKey,omitempty"`
	TokenValue string `json:"tokenValue,omitempty"`
	User       string `json:"user,omitempty"`
	IsCLI      bool   `json:"isCLI,omitempty"`
}
