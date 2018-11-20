package client

const (
	PrometheusSpecType                               = "prometheusSpec"
	PrometheusSpecFieldAdditionalAlertManagerConfigs = "additionalAlertManagerConfigs"
	PrometheusSpecFieldAdditionalScrapeConfigs       = "additionalScrapeConfigs"
	PrometheusSpecFieldAffinity                      = "affinity"
	PrometheusSpecFieldAlerting                      = "alerting"
	PrometheusSpecFieldBaseImage                     = "baseImage"
	PrometheusSpecFieldContainers                    = "containers"
	PrometheusSpecFieldEvaluationInterval            = "evaluationInterval"
	PrometheusSpecFieldExternalLabels                = "externalLabels"
	PrometheusSpecFieldExternalURL                   = "externalUrl"
	PrometheusSpecFieldImagePullSecrets              = "imagePullSecrets"
	PrometheusSpecFieldListenLocal                   = "listenLocal"
	PrometheusSpecFieldLogLevel                      = "logLevel"
	PrometheusSpecFieldNodeSelector                  = "nodeSelector"
	PrometheusSpecFieldPodMetadata                   = "podMetadata"
	PrometheusSpecFieldRemoteRead                    = "remoteRead"
	PrometheusSpecFieldRemoteWrite                   = "remoteWrite"
	PrometheusSpecFieldReplicas                      = "replicas"
	PrometheusSpecFieldResources                     = "resources"
	PrometheusSpecFieldRetention                     = "retention"
	PrometheusSpecFieldRoutePrefix                   = "routePrefix"
	PrometheusSpecFieldRuleSelector                  = "ruleSelector"
	PrometheusSpecFieldScrapeInterval                = "scrapeInterval"
	PrometheusSpecFieldSecrets                       = "secrets"
	PrometheusSpecFieldSecurityContext               = "securityContext"
	PrometheusSpecFieldServiceAccountName            = "serviceAccountName"
	PrometheusSpecFieldServiceMonitorSelector        = "serviceMonitorSelector"
	PrometheusSpecFieldStorage                       = "storage"
	PrometheusSpecFieldTag                           = "tag"
	PrometheusSpecFieldTolerations                   = "tolerations"
	PrometheusSpecFieldVersion                       = "version"
)

type PrometheusSpec struct {
	AdditionalAlertManagerConfigs *SecretKeySelector     `json:"additionalAlertManagerConfigs,omitempty" yaml:"additionalAlertManagerConfigs,omitempty"`
	AdditionalScrapeConfigs       *SecretKeySelector     `json:"additionalScrapeConfigs,omitempty" yaml:"additionalScrapeConfigs,omitempty"`
	Affinity                      *Affinity              `json:"affinity,omitempty" yaml:"affinity,omitempty"`
	Alerting                      *AlertingSpec          `json:"alerting,omitempty" yaml:"alerting,omitempty"`
	BaseImage                     string                 `json:"baseImage,omitempty" yaml:"baseImage,omitempty"`
	Containers                    []Container            `json:"containers,omitempty" yaml:"containers,omitempty"`
	EvaluationInterval            string                 `json:"evaluationInterval,omitempty" yaml:"evaluationInterval,omitempty"`
	ExternalLabels                map[string]string      `json:"externalLabels,omitempty" yaml:"externalLabels,omitempty"`
	ExternalURL                   string                 `json:"externalUrl,omitempty" yaml:"externalUrl,omitempty"`
	ImagePullSecrets              []LocalObjectReference `json:"imagePullSecrets,omitempty" yaml:"imagePullSecrets,omitempty"`
	ListenLocal                   bool                   `json:"listenLocal,omitempty" yaml:"listenLocal,omitempty"`
	LogLevel                      string                 `json:"logLevel,omitempty" yaml:"logLevel,omitempty"`
	NodeSelector                  map[string]string      `json:"nodeSelector,omitempty" yaml:"nodeSelector,omitempty"`
	PodMetadata                   *ObjectMeta            `json:"podMetadata,omitempty" yaml:"podMetadata,omitempty"`
	RemoteRead                    []RemoteReadSpec       `json:"remoteRead,omitempty" yaml:"remoteRead,omitempty"`
	RemoteWrite                   []RemoteWriteSpec      `json:"remoteWrite,omitempty" yaml:"remoteWrite,omitempty"`
	Replicas                      *int64                 `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	Resources                     *ResourceRequirements  `json:"resources,omitempty" yaml:"resources,omitempty"`
	Retention                     string                 `json:"retention,omitempty" yaml:"retention,omitempty"`
	RoutePrefix                   string                 `json:"routePrefix,omitempty" yaml:"routePrefix,omitempty"`
	RuleSelector                  *LabelSelector         `json:"ruleSelector,omitempty" yaml:"ruleSelector,omitempty"`
	ScrapeInterval                string                 `json:"scrapeInterval,omitempty" yaml:"scrapeInterval,omitempty"`
	Secrets                       []string               `json:"secrets,omitempty" yaml:"secrets,omitempty"`
	SecurityContext               *PodSecurityContext    `json:"securityContext,omitempty" yaml:"securityContext,omitempty"`
	ServiceAccountName            string                 `json:"serviceAccountName,omitempty" yaml:"serviceAccountName,omitempty"`
	ServiceMonitorSelector        *LabelSelector         `json:"serviceMonitorSelector,omitempty" yaml:"serviceMonitorSelector,omitempty"`
	Storage                       *StorageSpec           `json:"storage,omitempty" yaml:"storage,omitempty"`
	Tag                           string                 `json:"tag,omitempty" yaml:"tag,omitempty"`
	Tolerations                   []Toleration           `json:"tolerations,omitempty" yaml:"tolerations,omitempty"`
	Version                       string                 `json:"version,omitempty" yaml:"version,omitempty"`
}
