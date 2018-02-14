package client

import "k8s.io/apimachinery/pkg/util/intstr"

const (
	IngressBackendType            = "ingressBackend"
	IngressBackendFieldServiceId  = "serviceId"
	IngressBackendFieldTargetPort = "targetPort"
	IngressBackendFieldWorkloadId = "workloadId"
)

type IngressBackend struct {
	ServiceId  string             `json:"serviceId,omitempty"`
	TargetPort intstr.IntOrString `json:"targetPort,omitempty"`
	WorkloadId string             `json:"workloadId,omitempty"`
}
