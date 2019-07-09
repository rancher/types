package client

const (
	AlertmanagerSpecType                    = "alertmanagerSpec"
	AlertmanagerSpecFieldAdditionalPeers    = "additionalPeers"
	AlertmanagerSpecFieldAffinity           = "affinity"
	AlertmanagerSpecFieldBaseImage          = "baseImage"
	AlertmanagerSpecFieldConfigMaps         = "configMaps"
	AlertmanagerSpecFieldContainers         = "containers"
	AlertmanagerSpecFieldExternalURL        = "externalUrl"
	AlertmanagerSpecFieldImagePullSecrets   = "imagePullSecrets"
	AlertmanagerSpecFieldListenLocal        = "listenLocal"
	AlertmanagerSpecFieldLogLevel           = "logLevel"
	AlertmanagerSpecFieldNodeSelector       = "nodeSelector"
	AlertmanagerSpecFieldPaused             = "paused"
	AlertmanagerSpecFieldPodMetadata        = "podMetadata"
	AlertmanagerSpecFieldPriorityClassName  = "priorityClassName"
	AlertmanagerSpecFieldReplicas           = "replicas"
	AlertmanagerSpecFieldResources          = "resources"
	AlertmanagerSpecFieldRetention          = "retention"
	AlertmanagerSpecFieldRoutePrefix        = "routePrefix"
	AlertmanagerSpecFieldSHA                = "sha"
	AlertmanagerSpecFieldSecrets            = "secrets"
	AlertmanagerSpecFieldSecurityContext    = "securityContext"
	AlertmanagerSpecFieldServiceAccountName = "serviceAccountName"
	AlertmanagerSpecFieldStorage            = "storage"
	AlertmanagerSpecFieldTag                = "tag"
	AlertmanagerSpecFieldTolerations        = "tolerations"
	AlertmanagerSpecFieldVersion            = "version"
)

type AlertmanagerSpec struct {
	AdditionalPeers    []string               `json:"additionalPeers,omitempty" yaml:"additional_peers,omitempty"`
	Affinity           *Affinity              `json:"affinity,omitempty" yaml:"affinity,omitempty"`
	BaseImage          string                 `json:"baseImage,omitempty" yaml:"base_image,omitempty"`
	ConfigMaps         []string               `json:"configMaps,omitempty" yaml:"config_maps,omitempty"`
	Containers         []Container            `json:"containers,omitempty" yaml:"containers,omitempty"`
	ExternalURL        string                 `json:"externalUrl,omitempty" yaml:"external_url,omitempty"`
	ImagePullSecrets   []LocalObjectReference `json:"imagePullSecrets,omitempty" yaml:"image_pull_secrets,omitempty"`
	ListenLocal        bool                   `json:"listenLocal,omitempty" yaml:"listen_local,omitempty"`
	LogLevel           string                 `json:"logLevel,omitempty" yaml:"log_level,omitempty"`
	NodeSelector       map[string]string      `json:"nodeSelector,omitempty" yaml:"node_selector,omitempty"`
	Paused             bool                   `json:"paused,omitempty" yaml:"paused,omitempty"`
	PodMetadata        *ObjectMeta            `json:"podMetadata,omitempty" yaml:"pod_metadata,omitempty"`
	PriorityClassName  string                 `json:"priorityClassName,omitempty" yaml:"priority_class_name,omitempty"`
	Replicas           *int64                 `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	Resources          *ResourceRequirements  `json:"resources,omitempty" yaml:"resources,omitempty"`
	Retention          string                 `json:"retention,omitempty" yaml:"retention,omitempty"`
	RoutePrefix        string                 `json:"routePrefix,omitempty" yaml:"route_prefix,omitempty"`
	SHA                string                 `json:"sha,omitempty" yaml:"sha,omitempty"`
	Secrets            []string               `json:"secrets,omitempty" yaml:"secrets,omitempty"`
	SecurityContext    *PodSecurityContext    `json:"securityContext,omitempty" yaml:"security_context,omitempty"`
	ServiceAccountName string                 `json:"serviceAccountName,omitempty" yaml:"service_account_name,omitempty"`
	Storage            *StorageSpec           `json:"storage,omitempty" yaml:"storage,omitempty"`
	Tag                string                 `json:"tag,omitempty" yaml:"tag,omitempty"`
	Tolerations        []Toleration           `json:"tolerations,omitempty" yaml:"tolerations,omitempty"`
	Version            string                 `json:"version,omitempty" yaml:"version,omitempty"`
}
