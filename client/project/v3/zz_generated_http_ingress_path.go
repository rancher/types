package client

import "k8s.io/apimachinery/pkg/util/intstr"

const (
	HTTPIngressPathType             = "httpIngressPath"
	HTTPIngressPathFieldPath        = "path"
	HTTPIngressPathFieldServiceID   = "serviceId"
	HTTPIngressPathFieldTargetPort  = "targetPort"
	HTTPIngressPathFieldWorkloadIDs = "workloadIds"
)

type HTTPIngressPath struct {
	Path        string             `json:"path,omitempty" yaml:"path,omitempty"`
	ServiceID   string             `json:"serviceId,omitempty" yaml:"service_id,omitempty"`
	TargetPort  intstr.IntOrString `json:"targetPort,omitempty" yaml:"target_port,omitempty"`
	WorkloadIDs []string           `json:"workloadIds,omitempty" yaml:"workload_ids,omitempty"`
}
