package client

import "k8s.io/apimachinery/pkg/util/intstr"

const (
	HTTPIngressPathType            = "httpIngressPath"
	HTTPIngressPathFieldServiceId  = "serviceId"
	HTTPIngressPathFieldTargetPort = "targetPort"
	HTTPIngressPathFieldWorkloadId = "workloadId"
)

type HTTPIngressPath struct {
	ServiceId  string             `json:"serviceId,omitempty"`
	TargetPort intstr.IntOrString `json:"targetPort,omitempty"`
	WorkloadId string             `json:"workloadId,omitempty"`
}
