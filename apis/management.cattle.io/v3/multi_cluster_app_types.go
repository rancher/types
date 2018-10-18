package v3

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type GlobalDNS struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	DNSName      string `json:"dnsName,omitempty" norman:"required"`
	RootDomain   string `json:"rootDomain" norman:"required"`
	TTLSeconds   int64  `json:"ttl" norman:"default=300"`
	ProviderName string `json:"providerName,omitempty" norman:"required"`
}
