package client

import (
	"github.com/rancher/norman/types"
)

const (
	AlertmanagerType                      = "alertmanager"
	AlertmanagerFieldAdditionalPeers      = "additionalPeers"
	AlertmanagerFieldAffinity             = "affinity"
	AlertmanagerFieldAnnotations          = "annotations"
	AlertmanagerFieldBaseImage            = "baseImage"
	AlertmanagerFieldConfigMaps           = "configMaps"
	AlertmanagerFieldContainers           = "containers"
	AlertmanagerFieldCreated              = "created"
	AlertmanagerFieldCreatorID            = "creatorId"
	AlertmanagerFieldExternalURL          = "externalUrl"
	AlertmanagerFieldImagePullSecrets     = "imagePullSecrets"
	AlertmanagerFieldLabels               = "labels"
	AlertmanagerFieldListenLocal          = "listenLocal"
	AlertmanagerFieldLogLevel             = "logLevel"
	AlertmanagerFieldName                 = "name"
	AlertmanagerFieldNamespaceId          = "namespaceId"
	AlertmanagerFieldNodeSelector         = "nodeSelector"
	AlertmanagerFieldOwnerReferences      = "ownerReferences"
	AlertmanagerFieldPaused               = "paused"
	AlertmanagerFieldPodMetadata          = "podMetadata"
	AlertmanagerFieldPriorityClassName    = "priorityClassName"
	AlertmanagerFieldProjectID            = "projectId"
	AlertmanagerFieldRemoved              = "removed"
	AlertmanagerFieldReplicas             = "replicas"
	AlertmanagerFieldResources            = "resources"
	AlertmanagerFieldRetention            = "retention"
	AlertmanagerFieldRoutePrefix          = "routePrefix"
	AlertmanagerFieldSHA                  = "sha"
	AlertmanagerFieldSecrets              = "secrets"
	AlertmanagerFieldSecurityContext      = "securityContext"
	AlertmanagerFieldServiceAccountName   = "serviceAccountName"
	AlertmanagerFieldState                = "state"
	AlertmanagerFieldStorage              = "storage"
	AlertmanagerFieldTag                  = "tag"
	AlertmanagerFieldTolerations          = "tolerations"
	AlertmanagerFieldTransitioning        = "transitioning"
	AlertmanagerFieldTransitioningMessage = "transitioningMessage"
	AlertmanagerFieldUUID                 = "uuid"
	AlertmanagerFieldVersion              = "version"
)

type Alertmanager struct {
	types.Resource
	AdditionalPeers      []string               `json:"additionalPeers,omitempty" yaml:"additional_peers,omitempty"`
	Affinity             *Affinity              `json:"affinity,omitempty" yaml:"affinity,omitempty"`
	Annotations          map[string]string      `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	BaseImage            string                 `json:"baseImage,omitempty" yaml:"base_image,omitempty"`
	ConfigMaps           []string               `json:"configMaps,omitempty" yaml:"config_maps,omitempty"`
	Containers           []Container            `json:"containers,omitempty" yaml:"containers,omitempty"`
	Created              string                 `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string                 `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	ExternalURL          string                 `json:"externalUrl,omitempty" yaml:"external_url,omitempty"`
	ImagePullSecrets     []LocalObjectReference `json:"imagePullSecrets,omitempty" yaml:"image_pull_secrets,omitempty"`
	Labels               map[string]string      `json:"labels,omitempty" yaml:"labels,omitempty"`
	ListenLocal          bool                   `json:"listenLocal,omitempty" yaml:"listen_local,omitempty"`
	LogLevel             string                 `json:"logLevel,omitempty" yaml:"log_level,omitempty"`
	Name                 string                 `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId          string                 `json:"namespaceId,omitempty" yaml:"namespace_id,omitempty"`
	NodeSelector         map[string]string      `json:"nodeSelector,omitempty" yaml:"node_selector,omitempty"`
	OwnerReferences      []OwnerReference       `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	Paused               bool                   `json:"paused,omitempty" yaml:"paused,omitempty"`
	PodMetadata          *ObjectMeta            `json:"podMetadata,omitempty" yaml:"pod_metadata,omitempty"`
	PriorityClassName    string                 `json:"priorityClassName,omitempty" yaml:"priority_class_name,omitempty"`
	ProjectID            string                 `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	Removed              string                 `json:"removed,omitempty" yaml:"removed,omitempty"`
	Replicas             *int64                 `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	Resources            *ResourceRequirements  `json:"resources,omitempty" yaml:"resources,omitempty"`
	Retention            string                 `json:"retention,omitempty" yaml:"retention,omitempty"`
	RoutePrefix          string                 `json:"routePrefix,omitempty" yaml:"route_prefix,omitempty"`
	SHA                  string                 `json:"sha,omitempty" yaml:"sha,omitempty"`
	Secrets              []string               `json:"secrets,omitempty" yaml:"secrets,omitempty"`
	SecurityContext      *PodSecurityContext    `json:"securityContext,omitempty" yaml:"security_context,omitempty"`
	ServiceAccountName   string                 `json:"serviceAccountName,omitempty" yaml:"service_account_name,omitempty"`
	State                string                 `json:"state,omitempty" yaml:"state,omitempty"`
	Storage              *StorageSpec           `json:"storage,omitempty" yaml:"storage,omitempty"`
	Tag                  string                 `json:"tag,omitempty" yaml:"tag,omitempty"`
	Tolerations          []Toleration           `json:"tolerations,omitempty" yaml:"tolerations,omitempty"`
	Transitioning        string                 `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage string                 `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`
	UUID                 string                 `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Version              string                 `json:"version,omitempty" yaml:"version,omitempty"`
}

type AlertmanagerCollection struct {
	types.Collection
	Data   []Alertmanager `json:"data,omitempty"`
	client *AlertmanagerClient
}

type AlertmanagerClient struct {
	apiClient *Client
}

type AlertmanagerOperations interface {
	List(opts *types.ListOpts) (*AlertmanagerCollection, error)
	Create(opts *Alertmanager) (*Alertmanager, error)
	Update(existing *Alertmanager, updates interface{}) (*Alertmanager, error)
	Replace(existing *Alertmanager) (*Alertmanager, error)
	ByID(id string) (*Alertmanager, error)
	Delete(container *Alertmanager) error
}

func newAlertmanagerClient(apiClient *Client) *AlertmanagerClient {
	return &AlertmanagerClient{
		apiClient: apiClient,
	}
}

func (c *AlertmanagerClient) Create(container *Alertmanager) (*Alertmanager, error) {
	resp := &Alertmanager{}
	err := c.apiClient.Ops.DoCreate(AlertmanagerType, container, resp)
	return resp, err
}

func (c *AlertmanagerClient) Update(existing *Alertmanager, updates interface{}) (*Alertmanager, error) {
	resp := &Alertmanager{}
	err := c.apiClient.Ops.DoUpdate(AlertmanagerType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *AlertmanagerClient) Replace(obj *Alertmanager) (*Alertmanager, error) {
	resp := &Alertmanager{}
	err := c.apiClient.Ops.DoReplace(AlertmanagerType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *AlertmanagerClient) List(opts *types.ListOpts) (*AlertmanagerCollection, error) {
	resp := &AlertmanagerCollection{}
	err := c.apiClient.Ops.DoList(AlertmanagerType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *AlertmanagerCollection) Next() (*AlertmanagerCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &AlertmanagerCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *AlertmanagerClient) ByID(id string) (*Alertmanager, error) {
	resp := &Alertmanager{}
	err := c.apiClient.Ops.DoByID(AlertmanagerType, id, resp)
	return resp, err
}

func (c *AlertmanagerClient) Delete(container *Alertmanager) error {
	return c.apiClient.Ops.DoResourceDelete(AlertmanagerType, &container.Resource)
}
