package client

import "k8s.io/apimachinery/pkg/util/intstr"

const (
	IngressBackendType             = "ingressBackend"
	IngressBackendFieldServiceID   = "serviceId"
	IngressBackendFieldTargetPort  = "targetPort"
	IngressBackendFieldWorkloadIDs = "workloadIds"
)

type IngressBackend struct {
	ServiceID   string             `json:"serviceId,omitempty" yaml:"service_id,omitempty"`
	TargetPort  intstr.IntOrString `json:"targetPort,omitempty" yaml:"target_port,omitempty"`
	WorkloadIDs []string           `json:"workloadIds,omitempty" yaml:"workload_ids,omitempty"`
}
