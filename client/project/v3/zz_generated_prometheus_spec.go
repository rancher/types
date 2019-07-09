package client

const (
	PrometheusSpecType                               = "prometheusSpec"
	PrometheusSpecFieldAdditionalAlertManagerConfigs = "additionalAlertManagerConfigs"
	PrometheusSpecFieldAdditionalAlertRelabelConfigs = "additionalAlertRelabelConfigs"
	PrometheusSpecFieldAdditionalScrapeConfigs       = "additionalScrapeConfigs"
	PrometheusSpecFieldAffinity                      = "affinity"
	PrometheusSpecFieldAlerting                      = "alerting"
	PrometheusSpecFieldBaseImage                     = "baseImage"
	PrometheusSpecFieldConfigMaps                    = "configMaps"
	PrometheusSpecFieldContainers                    = "containers"
	PrometheusSpecFieldEvaluationInterval            = "evaluationInterval"
	PrometheusSpecFieldExternalLabels                = "externalLabels"
	PrometheusSpecFieldExternalURL                   = "externalUrl"
	PrometheusSpecFieldImagePullSecrets              = "imagePullSecrets"
	PrometheusSpecFieldListenLocal                   = "listenLocal"
	PrometheusSpecFieldLogLevel                      = "logLevel"
	PrometheusSpecFieldNodeSelector                  = "nodeSelector"
	PrometheusSpecFieldPodMetadata                   = "podMetadata"
	PrometheusSpecFieldPriorityClassName             = "priorityClassName"
	PrometheusSpecFieldRemoteRead                    = "remoteRead"
	PrometheusSpecFieldRemoteWrite                   = "remoteWrite"
	PrometheusSpecFieldReplicas                      = "replicas"
	PrometheusSpecFieldResources                     = "resources"
	PrometheusSpecFieldRetention                     = "retention"
	PrometheusSpecFieldRoutePrefix                   = "routePrefix"
	PrometheusSpecFieldRuleSelector                  = "ruleSelector"
	PrometheusSpecFieldSHA                           = "sha"
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
	AdditionalAlertManagerConfigs *SecretKeySelector     `json:"additionalAlertManagerConfigs,omitempty" yaml:"additional_alert_manager_configs,omitempty"`
	AdditionalAlertRelabelConfigs *SecretKeySelector     `json:"additionalAlertRelabelConfigs,omitempty" yaml:"additional_alert_relabel_configs,omitempty"`
	AdditionalScrapeConfigs       *SecretKeySelector     `json:"additionalScrapeConfigs,omitempty" yaml:"additional_scrape_configs,omitempty"`
	Affinity                      *Affinity              `json:"affinity,omitempty" yaml:"affinity,omitempty"`
	Alerting                      *AlertingSpec          `json:"alerting,omitempty" yaml:"alerting,omitempty"`
	BaseImage                     string                 `json:"baseImage,omitempty" yaml:"base_image,omitempty"`
	ConfigMaps                    []string               `json:"configMaps,omitempty" yaml:"config_maps,omitempty"`
	Containers                    []Container            `json:"containers,omitempty" yaml:"containers,omitempty"`
	EvaluationInterval            string                 `json:"evaluationInterval,omitempty" yaml:"evaluation_interval,omitempty"`
	ExternalLabels                map[string]string      `json:"externalLabels,omitempty" yaml:"external_labels,omitempty"`
	ExternalURL                   string                 `json:"externalUrl,omitempty" yaml:"external_url,omitempty"`
	ImagePullSecrets              []LocalObjectReference `json:"imagePullSecrets,omitempty" yaml:"image_pull_secrets,omitempty"`
	ListenLocal                   bool                   `json:"listenLocal,omitempty" yaml:"listen_local,omitempty"`
	LogLevel                      string                 `json:"logLevel,omitempty" yaml:"log_level,omitempty"`
	NodeSelector                  map[string]string      `json:"nodeSelector,omitempty" yaml:"node_selector,omitempty"`
	PodMetadata                   *ObjectMeta            `json:"podMetadata,omitempty" yaml:"pod_metadata,omitempty"`
	PriorityClassName             string                 `json:"priorityClassName,omitempty" yaml:"priority_class_name,omitempty"`
	RemoteRead                    []RemoteReadSpec       `json:"remoteRead,omitempty" yaml:"remote_read,omitempty"`
	RemoteWrite                   []RemoteWriteSpec      `json:"remoteWrite,omitempty" yaml:"remote_write,omitempty"`
	Replicas                      *int64                 `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	Resources                     *ResourceRequirements  `json:"resources,omitempty" yaml:"resources,omitempty"`
	Retention                     string                 `json:"retention,omitempty" yaml:"retention,omitempty"`
	RoutePrefix                   string                 `json:"routePrefix,omitempty" yaml:"route_prefix,omitempty"`
	RuleSelector                  *LabelSelector         `json:"ruleSelector,omitempty" yaml:"rule_selector,omitempty"`
	SHA                           string                 `json:"sha,omitempty" yaml:"sha,omitempty"`
	ScrapeInterval                string                 `json:"scrapeInterval,omitempty" yaml:"scrape_interval,omitempty"`
	Secrets                       []string               `json:"secrets,omitempty" yaml:"secrets,omitempty"`
	SecurityContext               *PodSecurityContext    `json:"securityContext,omitempty" yaml:"security_context,omitempty"`
	ServiceAccountName            string                 `json:"serviceAccountName,omitempty" yaml:"service_account_name,omitempty"`
	ServiceMonitorSelector        *LabelSelector         `json:"serviceMonitorSelector,omitempty" yaml:"service_monitor_selector,omitempty"`
	Storage                       *StorageSpec           `json:"storage,omitempty" yaml:"storage,omitempty"`
	Tag                           string                 `json:"tag,omitempty" yaml:"tag,omitempty"`
	Tolerations                   []Toleration           `json:"tolerations,omitempty" yaml:"tolerations,omitempty"`
	Version                       string                 `json:"version,omitempty" yaml:"version,omitempty"`
}
