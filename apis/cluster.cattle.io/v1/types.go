package v1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterConditionType string

const (
	// ClusterConditionReady Cluster ready to serve API (healthy when true, unehalthy when false)
	ClusterConditionReady = "Ready"
	// ClusterConditionProvisioned Cluster is provisioned
	ClusterConditionProvisioned = "Provisioned"
	// ClusterConditionUpdating Cluster is being updating (upgrading, scaling up)
	ClusterConditionUpdating = "Updating"
	// ClusterConditionNoDiskPressure true when all cluster nodes have sufficient disk
	ClusterConditionNoDiskPressure = "NoDiskPressure"
	// ClusterConditionNoMemoryPressure true when all cluster nodes have sufficient memory
	ClusterConditionNoMemoryPressure = "NoMemoryPressure"
	// More conditions can be added if unredlying controllers request it
)

type Cluster struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec ClusterSpec `json:"spec"`
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status *ClusterStatus `json:"status"`
}

type ClusterSpec struct {
	GoogleKubernetesEngineConfig *GoogleKubernetesEngineConfig `json:"googleKubernetesEngineConfig,omitempty"`
	AzureKubernetesServiceConfig *AzureKubernetesServiceConfig `json:"azureKubernetesServiceConfig,omitempty"`
	RKEConfig                    *RKEConfig                    `json:"rkeConfig,omitempty"`
}

type ClusterStatus struct {
	//Conditions represent the latest available observations of an object's current state:
	//More info: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#typical-status-properties
	Conditions []ClusterCondition `json:"conditions,omitempty"`
	//Component statuses will represent cluster's components (etcd/controller/scheduler) health
	// https://kubernetes.io/docs/api-reference/v1.8/#componentstatus-v1-core
	ComponentStatuses   []ClusterComponentStatus
	APIEndpoint         string          `json:"apiEndpoint,omitempty"`
	ServiceAccountToken string          `json:"serviceAccountToken,omitempty"`
	CACert              string          `json:"caCert,omitempty"`
	Capacity            v1.ResourceList `json:"capacity,omitempty"`
	Allocatable         v1.ResourceList `json:"allocatable,omitempty"`
	AppliedSpec         ClusterSpec     `json:"clusterSpec,omitempty"`
}

type ClusterComponentStatus struct {
	Name       string
	Conditions []v1.ComponentCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,2,rep,name=conditions"`
}

type ClusterCondition struct {
	// Type of cluster condition.
	Type ClusterConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// The last time this condition was updated.
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
}

type GoogleKubernetesEngineConfig struct {
	// ProjectID is the ID of your project to use when creating a cluster
	ProjectID string `json:"projectId,omitempty"`
	// The zone to launch the cluster
	Zone string `json:"zone,omitempty"`
	// The IP address range of the container pods
	ClusterIpv4Cidr string `json:"clusterIpv4Cidr,omitempty"`
	// An optional description of this cluster
	Description string `json:"description,omitempty"`
	// The number of nodes in this cluster
	NodeCount int64 `json:"nodeCount,omitempty"`
	// Size of the disk attached to each node
	DiskSizeGb int64 `json:"diskSizeGb,omitempty"`
	// The name of a Google Compute Engine
	MachineType string `json:"machineType,omitempty"`
	// Node kubernetes version
	NodeVersion string `json:"nodeVersion,omitempty"`
	// the master kubernetes version
	MasterVersion string `json:"masterVersion,omitempty"`
	// The map of Kubernetes labels (key/value pairs) to be applied
	// to each node.
	Labels map[string]string `json:"labels,omitempty"`
	// The path to the credential file(key.json)
	CredentialPath string `json:"credentialPath,omitempty"`
	// Enable alpha feature
	EnableAlphaFeature bool `json:"enableAlphaFeature,omitempty"`
}

type AzureKubernetesServiceConfig struct {
	//TBD
}

type RKEConfig struct {
	// Kubernetes nodes
	Hosts []RKEConfigHost `yaml:"hosts"`
	// Kubernetes components
	Services RKEConfigServices `yaml:"services"`
	// Network plugin used in the kubernetes cluster (flannel, calico)
	NetworkPlugin string `yaml:"network_plugin"`
	// Authentication type used in the cluster (default: x509)
	AuthType string `yaml:"auth_type"`
}

type RKEConfigHost struct {
	// SSH IP address of the host
	IP string `yaml:"ip"`
	// Advertised address that will be used for components communication
	AdvertiseAddress string `yaml:"advertise_address"`
	// Host role in kubernetes cluster (controlplane, worker, or etcd)
	Role []string `yaml:"role"`
	// Hostname of the host
	Hostname string `yaml:"hostname"`
	// SSH usesr that will be used by RKE
	User string `yaml:"user"`
	// Docker socket on the host that will be used in tunneling
	DockerSocket string `yaml:"docker_socket"`
}

type RKEConfigServices struct {
	// Etcd Service
	Etcd ETCDService `yaml:"etcd"`
	// KubeAPI Service
	KubeAPI KubeAPIService `yaml:"kube-api"`
	// KubeController Service
	KubeController KubeControllerService `yaml:"kube-controller"`
	// Scheduler Service
	Scheduler SchedulerService `yaml:"scheduler"`
	// Kubelet Service
	Kubelet KubeletService `yaml:"kubelet"`
	// KubeProxy Service
	Kubeproxy KubeproxyService `yaml:"kubeproxy"`
}

type ETCDService struct {
	// Base service properties
	BaseService `yaml:",inline"`
}

type KubeAPIService struct {
	// Base service properties
	BaseService `yaml:",inline"`
	// Virtual IP range that will be used by Kubernetes services
	ServiceClusterIPRange string `yaml:"service_cluster_ip_range"`
}

type KubeControllerService struct {
	// Base service properties
	BaseService `yaml:",inline"`
	// CIDR Range for Pods in cluster
	ClusterCIDR string `yaml:"cluster_cidr"`
	// Virtual IP range that will be used by Kubernetes services
	ServiceClusterIPRange string `yaml:"service_cluster_ip_range"`
}

type KubeletService struct {
	// Base service properties
	BaseService `yaml:",inline"`
	// Domain of the cluster (default: "cluster.local")
	ClusterDomain string `yaml:"cluster_domain"`
	// The image whose network/ipc namespaces containers in each pod will use
	InfraContainerImage string `yaml:"infra_container_image"`
	// Cluster DNS service ip
	ClusterDNSServer string `yaml:"cluster_dns_server"`
}

type KubeproxyService struct {
	// Base service properties
	BaseService `yaml:",inline"`
}

type SchedulerService struct {
	// Base service properties
	BaseService `yaml:",inline"`
}

type BaseService struct {
	// Docker image of the service
	Image string `yaml:"image"`
	// Extra arguments that are added to the services
	ExtraArgs []string `yaml:"extra_args"`
}

type ClusterNode struct {
	v1.Node
}
