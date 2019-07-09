package client

import "k8s.io/apimachinery/pkg/util/intstr"

const (
	ProbeType                     = "probe"
	ProbeFieldCommand             = "command"
	ProbeFieldFailureThreshold    = "failureThreshold"
	ProbeFieldHTTPHeaders         = "httpHeaders"
	ProbeFieldHost                = "host"
	ProbeFieldInitialDelaySeconds = "initialDelaySeconds"
	ProbeFieldPath                = "path"
	ProbeFieldPeriodSeconds       = "periodSeconds"
	ProbeFieldPort                = "port"
	ProbeFieldScheme              = "scheme"
	ProbeFieldSuccessThreshold    = "successThreshold"
	ProbeFieldTCP                 = "tcp"
	ProbeFieldTimeoutSeconds      = "timeoutSeconds"
)

type Probe struct {
	Command             []string           `json:"command,omitempty" yaml:"command,omitempty"`
	FailureThreshold    int64              `json:"failureThreshold,omitempty" yaml:"failure_threshold,omitempty"`
	HTTPHeaders         []HTTPHeader       `json:"httpHeaders,omitempty" yaml:"http_headers,omitempty"`
	Host                string             `json:"host,omitempty" yaml:"host,omitempty"`
	InitialDelaySeconds int64              `json:"initialDelaySeconds,omitempty" yaml:"initial_delay_seconds,omitempty"`
	Path                string             `json:"path,omitempty" yaml:"path,omitempty"`
	PeriodSeconds       int64              `json:"periodSeconds,omitempty" yaml:"period_seconds,omitempty"`
	Port                intstr.IntOrString `json:"port,omitempty" yaml:"port,omitempty"`
	Scheme              string             `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	SuccessThreshold    int64              `json:"successThreshold,omitempty" yaml:"success_threshold,omitempty"`
	TCP                 bool               `json:"tcp,omitempty" yaml:"tcp,omitempty"`
	TimeoutSeconds      int64              `json:"timeoutSeconds,omitempty" yaml:"timeout_seconds,omitempty"`
}
