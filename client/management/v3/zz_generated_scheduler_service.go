package client

const (
	SchedulerServiceType            = "schedulerService"
	SchedulerServiceFieldExtraArgs  = "extraArgs"
	SchedulerServiceFieldExtraBinds = "extraBinds"
	SchedulerServiceFieldExtraEnv   = "extraEnv"
	SchedulerServiceFieldImage      = "image"
)

type SchedulerService struct {
	ExtraArgs  map[string]string `json:"extraArgs,omitempty" yaml:"extra_args,omitempty"`
	ExtraBinds []string          `json:"extraBinds,omitempty" yaml:"extra_binds,omitempty"`
	ExtraEnv   []string          `json:"extraEnv,omitempty" yaml:"extra_env,omitempty"`
	Image      string            `json:"image,omitempty" yaml:"image,omitempty"`
}
