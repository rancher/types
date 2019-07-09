package client

const (
	StepType                      = "step"
	StepFieldApplyAppConfig       = "applyAppConfig"
	StepFieldApplyYamlConfig      = "applyYamlConfig"
	StepFieldCPULimit             = "cpuLimit"
	StepFieldCPURequest           = "cpuRequest"
	StepFieldEnv                  = "env"
	StepFieldEnvFrom              = "envFrom"
	StepFieldMemoryLimit          = "memoryLimit"
	StepFieldMemoryRequest        = "memoryRequest"
	StepFieldPrivileged           = "privileged"
	StepFieldPublishCatalogConfig = "publishCatalogConfig"
	StepFieldPublishImageConfig   = "publishImageConfig"
	StepFieldRunScriptConfig      = "runScriptConfig"
	StepFieldSourceCodeConfig     = "sourceCodeConfig"
	StepFieldWhen                 = "when"
)

type Step struct {
	ApplyAppConfig       *ApplyAppConfig       `json:"applyAppConfig,omitempty" yaml:"apply_app_config,omitempty"`
	ApplyYamlConfig      *ApplyYamlConfig      `json:"applyYamlConfig,omitempty" yaml:"apply_yaml_config,omitempty"`
	CPULimit             string                `json:"cpuLimit,omitempty" yaml:"cpu_limit,omitempty"`
	CPURequest           string                `json:"cpuRequest,omitempty" yaml:"cpu_request,omitempty"`
	Env                  map[string]string     `json:"env,omitempty" yaml:"env,omitempty"`
	EnvFrom              []EnvFrom             `json:"envFrom,omitempty" yaml:"env_from,omitempty"`
	MemoryLimit          string                `json:"memoryLimit,omitempty" yaml:"memory_limit,omitempty"`
	MemoryRequest        string                `json:"memoryRequest,omitempty" yaml:"memory_request,omitempty"`
	Privileged           bool                  `json:"privileged,omitempty" yaml:"privileged,omitempty"`
	PublishCatalogConfig *PublishCatalogConfig `json:"publishCatalogConfig,omitempty" yaml:"publish_catalog_config,omitempty"`
	PublishImageConfig   *PublishImageConfig   `json:"publishImageConfig,omitempty" yaml:"publish_image_config,omitempty"`
	RunScriptConfig      *RunScriptConfig      `json:"runScriptConfig,omitempty" yaml:"run_script_config,omitempty"`
	SourceCodeConfig     *SourceCodeConfig     `json:"sourceCodeConfig,omitempty" yaml:"source_code_config,omitempty"`
	When                 *Constraints          `json:"when,omitempty" yaml:"when,omitempty"`
}
