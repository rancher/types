package v3

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type GlobalDNS struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	//ObjectMeta.Name is the FQDN
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalDNSSpec   `json:"spec,omitempty"`
	Status GlobalDNSStatus `json:"status,omitempty"`
}

type GlobalDNSSpec struct {
	RootDomain   string   `json:"rootDomain" norman:"required"`
	ProviderName string   `json:"providerName,omitempty" norman:"required"`
	ClusterNames []string `json:"clusterNames,omitempty" norman:"type=array[reference[cluster]]"`
}

type GlobalDNSStatus struct {
	Endpoints []string `json:"endpoints,omitempty"`
}
