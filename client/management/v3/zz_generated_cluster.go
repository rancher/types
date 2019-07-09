package client

import (
	"github.com/rancher/norman/types"
)

const (
	ClusterType                                      = "cluster"
	ClusterFieldAPIEndpoint                          = "apiEndpoint"
	ClusterFieldAgentImage                           = "agentImage"
	ClusterFieldAllocatable                          = "allocatable"
	ClusterFieldAnnotations                          = "annotations"
	ClusterFieldAppliedEnableNetworkPolicy           = "appliedEnableNetworkPolicy"
	ClusterFieldAppliedPodSecurityPolicyTemplateName = "appliedPodSecurityPolicyTemplateId"
	ClusterFieldAppliedSpec                          = "appliedSpec"
	ClusterFieldAuthImage                            = "authImage"
	ClusterFieldCACert                               = "caCert"
	ClusterFieldCapabilities                         = "capabilities"
	ClusterFieldCapacity                             = "capacity"
	ClusterFieldCertificatesExpiration               = "certificatesExpiration"
	ClusterFieldClusterTemplateAnswers               = "answers"
	ClusterFieldClusterTemplateID                    = "clusterTemplateId"
	ClusterFieldClusterTemplateRevisionID            = "clusterTemplateRevisionId"
	ClusterFieldComponentStatuses                    = "componentStatuses"
	ClusterFieldConditions                           = "conditions"
	ClusterFieldCreated                              = "created"
	ClusterFieldCreatorID                            = "creatorId"
	ClusterFieldDefaultClusterRoleForProjectMembers  = "defaultClusterRoleForProjectMembers"
	ClusterFieldDefaultPodSecurityPolicyTemplateID   = "defaultPodSecurityPolicyTemplateId"
	ClusterFieldDescription                          = "description"
	ClusterFieldDesiredAgentImage                    = "desiredAgentImage"
	ClusterFieldDesiredAuthImage                     = "desiredAuthImage"
	ClusterFieldDockerRootDir                        = "dockerRootDir"
	ClusterFieldDriver                               = "driver"
	ClusterFieldEnableClusterAlerting                = "enableClusterAlerting"
	ClusterFieldEnableClusterMonitoring              = "enableClusterMonitoring"
	ClusterFieldEnableNetworkPolicy                  = "enableNetworkPolicy"
	ClusterFieldFailedSpec                           = "failedSpec"
	ClusterFieldImportedConfig                       = "importedConfig"
	ClusterFieldInternal                             = "internal"
	ClusterFieldIstioEnabled                         = "istioEnabled"
	ClusterFieldLabels                               = "labels"
	ClusterFieldLimits                               = "limits"
	ClusterFieldLocalClusterAuthEndpoint             = "localClusterAuthEndpoint"
	ClusterFieldMonitoringStatus                     = "monitoringStatus"
	ClusterFieldName                                 = "name"
	ClusterFieldOwnerReferences                      = "ownerReferences"
	ClusterFieldRancherKubernetesEngineConfig        = "rancherKubernetesEngineConfig"
	ClusterFieldRemoved                              = "removed"
	ClusterFieldRequested                            = "requested"
	ClusterFieldState                                = "state"
	ClusterFieldTransitioning                        = "transitioning"
	ClusterFieldTransitioningMessage                 = "transitioningMessage"
	ClusterFieldUUID                                 = "uuid"
	ClusterFieldVersion                              = "version"
)

type Cluster struct {
	types.Resource
	APIEndpoint                          string                         `json:"apiEndpoint,omitempty" yaml:"api_endpoint,omitempty"`
	AgentImage                           string                         `json:"agentImage,omitempty" yaml:"agent_image,omitempty"`
	Allocatable                          map[string]string              `json:"allocatable,omitempty" yaml:"allocatable,omitempty"`
	Annotations                          map[string]string              `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	AppliedEnableNetworkPolicy           bool                           `json:"appliedEnableNetworkPolicy,omitempty" yaml:"applied_enable_network_policy,omitempty"`
	AppliedPodSecurityPolicyTemplateName string                         `json:"appliedPodSecurityPolicyTemplateId,omitempty" yaml:"applied_pod_security_policy_template_id,omitempty"`
	AppliedSpec                          *ClusterSpec                   `json:"appliedSpec,omitempty" yaml:"applied_spec,omitempty"`
	AuthImage                            string                         `json:"authImage,omitempty" yaml:"auth_image,omitempty"`
	CACert                               string                         `json:"caCert,omitempty" yaml:"ca_cert,omitempty"`
	Capabilities                         *Capabilities                  `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
	Capacity                             map[string]string              `json:"capacity,omitempty" yaml:"capacity,omitempty"`
	CertificatesExpiration               map[string]CertExpiration      `json:"certificatesExpiration,omitempty" yaml:"certificates_expiration,omitempty"`
	ClusterTemplateAnswers               *Answer                        `json:"answers,omitempty" yaml:"answers,omitempty"`
	ClusterTemplateID                    string                         `json:"clusterTemplateId,omitempty" yaml:"cluster_template_id,omitempty"`
	ClusterTemplateRevisionID            string                         `json:"clusterTemplateRevisionId,omitempty" yaml:"cluster_template_revision_id,omitempty"`
	ComponentStatuses                    []ClusterComponentStatus       `json:"componentStatuses,omitempty" yaml:"component_statuses,omitempty"`
	Conditions                           []ClusterCondition             `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	Created                              string                         `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                            string                         `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	DefaultClusterRoleForProjectMembers  string                         `json:"defaultClusterRoleForProjectMembers,omitempty" yaml:"default_cluster_role_for_project_members,omitempty"`
	DefaultPodSecurityPolicyTemplateID   string                         `json:"defaultPodSecurityPolicyTemplateId,omitempty" yaml:"default_pod_security_policy_template_id,omitempty"`
	Description                          string                         `json:"description,omitempty" yaml:"description,omitempty"`
	DesiredAgentImage                    string                         `json:"desiredAgentImage,omitempty" yaml:"desired_agent_image,omitempty"`
	DesiredAuthImage                     string                         `json:"desiredAuthImage,omitempty" yaml:"desired_auth_image,omitempty"`
	DockerRootDir                        string                         `json:"dockerRootDir,omitempty" yaml:"docker_root_dir,omitempty"`
	Driver                               string                         `json:"driver,omitempty" yaml:"driver,omitempty"`
	EnableClusterAlerting                bool                           `json:"enableClusterAlerting,omitempty" yaml:"enable_cluster_alerting,omitempty"`
	EnableClusterMonitoring              bool                           `json:"enableClusterMonitoring,omitempty" yaml:"enable_cluster_monitoring,omitempty"`
	EnableNetworkPolicy                  *bool                          `json:"enableNetworkPolicy,omitempty" yaml:"enable_network_policy,omitempty"`
	FailedSpec                           *ClusterSpec                   `json:"failedSpec,omitempty" yaml:"failed_spec,omitempty"`
	ImportedConfig                       *ImportedConfig                `json:"importedConfig,omitempty" yaml:"imported_config,omitempty"`
	Internal                             bool                           `json:"internal,omitempty" yaml:"internal,omitempty"`
	IstioEnabled                         bool                           `json:"istioEnabled,omitempty" yaml:"istio_enabled,omitempty"`
	Labels                               map[string]string              `json:"labels,omitempty" yaml:"labels,omitempty"`
	Limits                               map[string]string              `json:"limits,omitempty" yaml:"limits,omitempty"`
	LocalClusterAuthEndpoint             *LocalClusterAuthEndpoint      `json:"localClusterAuthEndpoint,omitempty" yaml:"local_cluster_auth_endpoint,omitempty"`
	MonitoringStatus                     *MonitoringStatus              `json:"monitoringStatus,omitempty" yaml:"monitoring_status,omitempty"`
	Name                                 string                         `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences                      []OwnerReference               `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	RancherKubernetesEngineConfig        *RancherKubernetesEngineConfig `json:"rancherKubernetesEngineConfig,omitempty" yaml:"rancher_kubernetes_engine_config,omitempty"`
	Removed                              string                         `json:"removed,omitempty" yaml:"removed,omitempty"`
	Requested                            map[string]string              `json:"requested,omitempty" yaml:"requested,omitempty"`
	State                                string                         `json:"state,omitempty" yaml:"state,omitempty"`
	Transitioning                        string                         `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage                 string                         `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`
	UUID                                 string                         `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Version                              *Info                          `json:"version,omitempty" yaml:"version,omitempty"`
}

type ClusterCollection struct {
	types.Collection
	Data   []Cluster `json:"data,omitempty"`
	client *ClusterClient
}

type ClusterClient struct {
	apiClient *Client
}

type ClusterOperations interface {
	List(opts *types.ListOpts) (*ClusterCollection, error)
	Create(opts *Cluster) (*Cluster, error)
	Update(existing *Cluster, updates interface{}) (*Cluster, error)
	Replace(existing *Cluster) (*Cluster, error)
	ByID(id string) (*Cluster, error)
	Delete(container *Cluster) error

	ActionBackupEtcd(resource *Cluster) error

	ActionDisableMonitoring(resource *Cluster) error

	ActionEditMonitoring(resource *Cluster, input *MonitoringInput) error

	ActionEnableMonitoring(resource *Cluster, input *MonitoringInput) error

	ActionExportYaml(resource *Cluster) (*ExportOutput, error)

	ActionGenerateKubeconfig(resource *Cluster) (*GenerateKubeConfigOutput, error)

	ActionImportYaml(resource *Cluster, input *ImportClusterYamlInput) (*ImportYamlOutput, error)

	ActionRestoreFromEtcdBackup(resource *Cluster, input *RestoreFromEtcdBackupInput) error

	ActionRotateCertificates(resource *Cluster, input *RotateCertificateInput) (*RotateCertificateOutput, error)

	ActionRunSecurityScan(resource *Cluster) error

	ActionViewMonitoring(resource *Cluster) (*MonitoringOutput, error)
}

func newClusterClient(apiClient *Client) *ClusterClient {
	return &ClusterClient{
		apiClient: apiClient,
	}
}

func (c *ClusterClient) Create(container *Cluster) (*Cluster, error) {
	resp := &Cluster{}
	err := c.apiClient.Ops.DoCreate(ClusterType, container, resp)
	return resp, err
}

func (c *ClusterClient) Update(existing *Cluster, updates interface{}) (*Cluster, error) {
	resp := &Cluster{}
	err := c.apiClient.Ops.DoUpdate(ClusterType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ClusterClient) Replace(obj *Cluster) (*Cluster, error) {
	resp := &Cluster{}
	err := c.apiClient.Ops.DoReplace(ClusterType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ClusterClient) List(opts *types.ListOpts) (*ClusterCollection, error) {
	resp := &ClusterCollection{}
	err := c.apiClient.Ops.DoList(ClusterType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ClusterCollection) Next() (*ClusterCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ClusterCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ClusterClient) ByID(id string) (*Cluster, error) {
	resp := &Cluster{}
	err := c.apiClient.Ops.DoByID(ClusterType, id, resp)
	return resp, err
}

func (c *ClusterClient) Delete(container *Cluster) error {
	return c.apiClient.Ops.DoResourceDelete(ClusterType, &container.Resource)
}

func (c *ClusterClient) ActionBackupEtcd(resource *Cluster) error {
	err := c.apiClient.Ops.DoAction(ClusterType, "backupEtcd", &resource.Resource, nil, nil)
	return err
}

func (c *ClusterClient) ActionDisableMonitoring(resource *Cluster) error {
	err := c.apiClient.Ops.DoAction(ClusterType, "disableMonitoring", &resource.Resource, nil, nil)
	return err
}

func (c *ClusterClient) ActionEditMonitoring(resource *Cluster, input *MonitoringInput) error {
	err := c.apiClient.Ops.DoAction(ClusterType, "editMonitoring", &resource.Resource, input, nil)
	return err
}

func (c *ClusterClient) ActionEnableMonitoring(resource *Cluster, input *MonitoringInput) error {
	err := c.apiClient.Ops.DoAction(ClusterType, "enableMonitoring", &resource.Resource, input, nil)
	return err
}

func (c *ClusterClient) ActionExportYaml(resource *Cluster) (*ExportOutput, error) {
	resp := &ExportOutput{}
	err := c.apiClient.Ops.DoAction(ClusterType, "exportYaml", &resource.Resource, nil, resp)
	return resp, err
}

func (c *ClusterClient) ActionGenerateKubeconfig(resource *Cluster) (*GenerateKubeConfigOutput, error) {
	resp := &GenerateKubeConfigOutput{}
	err := c.apiClient.Ops.DoAction(ClusterType, "generateKubeconfig", &resource.Resource, nil, resp)
	return resp, err
}

func (c *ClusterClient) ActionImportYaml(resource *Cluster, input *ImportClusterYamlInput) (*ImportYamlOutput, error) {
	resp := &ImportYamlOutput{}
	err := c.apiClient.Ops.DoAction(ClusterType, "importYaml", &resource.Resource, input, resp)
	return resp, err
}

func (c *ClusterClient) ActionRestoreFromEtcdBackup(resource *Cluster, input *RestoreFromEtcdBackupInput) error {
	err := c.apiClient.Ops.DoAction(ClusterType, "restoreFromEtcdBackup", &resource.Resource, input, nil)
	return err
}

func (c *ClusterClient) ActionRotateCertificates(resource *Cluster, input *RotateCertificateInput) (*RotateCertificateOutput, error) {
	resp := &RotateCertificateOutput{}
	err := c.apiClient.Ops.DoAction(ClusterType, "rotateCertificates", &resource.Resource, input, resp)
	return resp, err
}

func (c *ClusterClient) ActionRunSecurityScan(resource *Cluster) error {
	err := c.apiClient.Ops.DoAction(ClusterType, "runSecurityScan", &resource.Resource, nil, nil)
	return err
}

func (c *ClusterClient) ActionViewMonitoring(resource *Cluster) (*MonitoringOutput, error) {
	resp := &MonitoringOutput{}
	err := c.apiClient.Ops.DoAction(ClusterType, "viewMonitoring", &resource.Resource, nil, resp)
	return resp, err
}
