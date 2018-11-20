package client

const (
	AlertmanagerSpecType                    = "alertmanagerSpec"
	AlertmanagerSpecFieldAffinity           = "affinity"
	AlertmanagerSpecFieldBaseImage          = "baseImage"
	AlertmanagerSpecFieldContainers         = "containers"
	AlertmanagerSpecFieldExternalURL        = "externalUrl"
	AlertmanagerSpecFieldImagePullSecrets   = "imagePullSecrets"
	AlertmanagerSpecFieldListenLocal        = "listenLocal"
	AlertmanagerSpecFieldLogLevel           = "logLevel"
	AlertmanagerSpecFieldNodeSelector       = "nodeSelector"
	AlertmanagerSpecFieldPaused             = "paused"
	AlertmanagerSpecFieldPodMetadata        = "podMetadata"
	AlertmanagerSpecFieldReplicas           = "replicas"
	AlertmanagerSpecFieldResources          = "resources"
	AlertmanagerSpecFieldRetention          = "retention"
	AlertmanagerSpecFieldRoutePrefix        = "routePrefix"
	AlertmanagerSpecFieldSecrets            = "secrets"
	AlertmanagerSpecFieldSecurityContext    = "securityContext"
	AlertmanagerSpecFieldServiceAccountName = "serviceAccountName"
	AlertmanagerSpecFieldStorage            = "storage"
	AlertmanagerSpecFieldTag                = "tag"
	AlertmanagerSpecFieldTolerations        = "tolerations"
	AlertmanagerSpecFieldVersion            = "version"
)

type AlertmanagerSpec struct {
	Affinity           *Affinity              `json:"affinity,omitempty" yaml:"affinity,omitempty"`
	BaseImage          string                 `json:"baseImage,omitempty" yaml:"baseImage,omitempty"`
	Containers         []Container            `json:"containers,omitempty" yaml:"containers,omitempty"`
	ExternalURL        string                 `json:"externalUrl,omitempty" yaml:"externalUrl,omitempty"`
	ImagePullSecrets   []LocalObjectReference `json:"imagePullSecrets,omitempty" yaml:"imagePullSecrets,omitempty"`
	ListenLocal        bool                   `json:"listenLocal,omitempty" yaml:"listenLocal,omitempty"`
	LogLevel           string                 `json:"logLevel,omitempty" yaml:"logLevel,omitempty"`
	NodeSelector       map[string]string      `json:"nodeSelector,omitempty" yaml:"nodeSelector,omitempty"`
	Paused             bool                   `json:"paused,omitempty" yaml:"paused,omitempty"`
	PodMetadata        *ObjectMeta            `json:"podMetadata,omitempty" yaml:"podMetadata,omitempty"`
	Replicas           *int64                 `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	Resources          *ResourceRequirements  `json:"resources,omitempty" yaml:"resources,omitempty"`
	Retention          string                 `json:"retention,omitempty" yaml:"retention,omitempty"`
	RoutePrefix        string                 `json:"routePrefix,omitempty" yaml:"routePrefix,omitempty"`
	Secrets            []string               `json:"secrets,omitempty" yaml:"secrets,omitempty"`
	SecurityContext    *PodSecurityContext    `json:"securityContext,omitempty" yaml:"securityContext,omitempty"`
	ServiceAccountName string                 `json:"serviceAccountName,omitempty" yaml:"serviceAccountName,omitempty"`
	Storage            *StorageSpec           `json:"storage,omitempty" yaml:"storage,omitempty"`
	Tag                string                 `json:"tag,omitempty" yaml:"tag,omitempty"`
	Tolerations        []Toleration           `json:"tolerations,omitempty" yaml:"tolerations,omitempty"`
	Version            string                 `json:"version,omitempty" yaml:"version,omitempty"`
}
