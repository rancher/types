package client

const (
	ResourceQuotaLimitType                        = "resourceQuotaLimit"
	ResourceQuotaLimitFieldConfigMaps             = "configMaps"
	ResourceQuotaLimitFieldLimitsCPU              = "limitsCpu"
	ResourceQuotaLimitFieldLimitsMemory           = "limitsMemory"
	ResourceQuotaLimitFieldPersistentVolumeClaims = "persistentVolumeClaims"
	ResourceQuotaLimitFieldPods                   = "pods"
	ResourceQuotaLimitFieldReplicationControllers = "replicationControllers"
	ResourceQuotaLimitFieldRequestsCPU            = "requestsCpu"
	ResourceQuotaLimitFieldRequestsMemory         = "requestsMemory"
	ResourceQuotaLimitFieldRequestsStorage        = "requestsStorage"
	ResourceQuotaLimitFieldSecrets                = "secrets"
	ResourceQuotaLimitFieldServices               = "services"
	ResourceQuotaLimitFieldServicesLoadBalancers  = "servicesLoadBalancers"
	ResourceQuotaLimitFieldServicesNodePorts      = "servicesNodePorts"
)

type ResourceQuotaLimit struct {
	ConfigMaps             string `json:"configMaps,omitempty" yaml:"config_maps,omitempty"`
	LimitsCPU              string `json:"limitsCpu,omitempty" yaml:"limits_cpu,omitempty"`
	LimitsMemory           string `json:"limitsMemory,omitempty" yaml:"limits_memory,omitempty"`
	PersistentVolumeClaims string `json:"persistentVolumeClaims,omitempty" yaml:"persistent_volume_claims,omitempty"`
	Pods                   string `json:"pods,omitempty" yaml:"pods,omitempty"`
	ReplicationControllers string `json:"replicationControllers,omitempty" yaml:"replication_controllers,omitempty"`
	RequestsCPU            string `json:"requestsCpu,omitempty" yaml:"requests_cpu,omitempty"`
	RequestsMemory         string `json:"requestsMemory,omitempty" yaml:"requests_memory,omitempty"`
	RequestsStorage        string `json:"requestsStorage,omitempty" yaml:"requests_storage,omitempty"`
	Secrets                string `json:"secrets,omitempty" yaml:"secrets,omitempty"`
	Services               string `json:"services,omitempty" yaml:"services,omitempty"`
	ServicesLoadBalancers  string `json:"servicesLoadBalancers,omitempty" yaml:"services_load_balancers,omitempty"`
	ServicesNodePorts      string `json:"servicesNodePorts,omitempty" yaml:"services_node_ports,omitempty"`
}
