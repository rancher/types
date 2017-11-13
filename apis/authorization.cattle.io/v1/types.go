package v1

import (
	extv1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	ClusterName string `json:"clusterName,omitempty"`
}

type RoleTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Rules []rbacv1.PolicyRule `json:"rules,omitempty"`

	RoleTemplates []string `json:"roles,omitempty"`
}

type PodSecurityPolicyTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec extv1.PodSecurityPolicySpec `json:"spec,omitempty"`
}

type ProjectRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Subjects []rbacv1.Subject `json:"subjects,omitempty"`

	ProjectName string `json:"projectRef,omitempty"`

	RoleTemplateName string `json:"roleTemplateName,omitempty"`
}
