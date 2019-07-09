package client

const (
	SetPodSecurityPolicyTemplateInputType                               = "setPodSecurityPolicyTemplateInput"
	SetPodSecurityPolicyTemplateInputFieldPodSecurityPolicyTemplateName = "podSecurityPolicyTemplateId"
)

type SetPodSecurityPolicyTemplateInput struct {
	PodSecurityPolicyTemplateName string `json:"podSecurityPolicyTemplateId,omitempty" yaml:"pod_security_policy_template_id,omitempty"`
}
