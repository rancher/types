package client

import (
	"github.com/rancher/norman/types"
)

const (
	CatalogTemplateType                          = "catalogTemplate"
	CatalogTemplateFieldAnnotations              = "annotations"
	CatalogTemplateFieldCatalogID                = "catalogId"
	CatalogTemplateFieldCategories               = "categories"
	CatalogTemplateFieldCategory                 = "category"
	CatalogTemplateFieldClusterCatalogID         = "clusterCatalogId"
	CatalogTemplateFieldClusterID                = "clusterId"
	CatalogTemplateFieldCreated                  = "created"
	CatalogTemplateFieldCreatorID                = "creatorId"
	CatalogTemplateFieldDefaultTemplateVersionID = "defaultTemplateVersionId"
	CatalogTemplateFieldDefaultVersion           = "defaultVersion"
	CatalogTemplateFieldDescription              = "description"
	CatalogTemplateFieldFolderName               = "folderName"
	CatalogTemplateFieldIcon                     = "icon"
	CatalogTemplateFieldIconFilename             = "iconFilename"
	CatalogTemplateFieldLabels                   = "labels"
	CatalogTemplateFieldMaintainer               = "maintainer"
	CatalogTemplateFieldName                     = "name"
	CatalogTemplateFieldOwnerReferences          = "ownerReferences"
	CatalogTemplateFieldPath                     = "path"
	CatalogTemplateFieldProjectCatalogID         = "projectCatalogId"
	CatalogTemplateFieldProjectID                = "projectId"
	CatalogTemplateFieldProjectURL               = "projectURL"
	CatalogTemplateFieldReadme                   = "readme"
	CatalogTemplateFieldRemoved                  = "removed"
	CatalogTemplateFieldState                    = "state"
	CatalogTemplateFieldStatus                   = "status"
	CatalogTemplateFieldTransitioning            = "transitioning"
	CatalogTemplateFieldTransitioningMessage     = "transitioningMessage"
	CatalogTemplateFieldUUID                     = "uuid"
	CatalogTemplateFieldUpgradeFrom              = "upgradeFrom"
	CatalogTemplateFieldVersionLinks             = "versionLinks"
	CatalogTemplateFieldVersions                 = "versions"
)

type CatalogTemplate struct {
	types.Resource
	Annotations              map[string]string     `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	CatalogID                string                `json:"catalogId,omitempty" yaml:"catalog_id,omitempty"`
	Categories               []string              `json:"categories,omitempty" yaml:"categories,omitempty"`
	Category                 string                `json:"category,omitempty" yaml:"category,omitempty"`
	ClusterCatalogID         string                `json:"clusterCatalogId,omitempty" yaml:"cluster_catalog_id,omitempty"`
	ClusterID                string                `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`
	Created                  string                `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                string                `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	DefaultTemplateVersionID string                `json:"defaultTemplateVersionId,omitempty" yaml:"default_template_version_id,omitempty"`
	DefaultVersion           string                `json:"defaultVersion,omitempty" yaml:"default_version,omitempty"`
	Description              string                `json:"description,omitempty" yaml:"description,omitempty"`
	FolderName               string                `json:"folderName,omitempty" yaml:"folder_name,omitempty"`
	Icon                     string                `json:"icon,omitempty" yaml:"icon,omitempty"`
	IconFilename             string                `json:"iconFilename,omitempty" yaml:"icon_filename,omitempty"`
	Labels                   map[string]string     `json:"labels,omitempty" yaml:"labels,omitempty"`
	Maintainer               string                `json:"maintainer,omitempty" yaml:"maintainer,omitempty"`
	Name                     string                `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences          []OwnerReference      `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	Path                     string                `json:"path,omitempty" yaml:"path,omitempty"`
	ProjectCatalogID         string                `json:"projectCatalogId,omitempty" yaml:"project_catalog_id,omitempty"`
	ProjectID                string                `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	ProjectURL               string                `json:"projectURL,omitempty" yaml:"project_url,omitempty"`
	Readme                   string                `json:"readme,omitempty" yaml:"readme,omitempty"`
	Removed                  string                `json:"removed,omitempty" yaml:"removed,omitempty"`
	State                    string                `json:"state,omitempty" yaml:"state,omitempty"`
	Status                   *TemplateStatus       `json:"status,omitempty" yaml:"status,omitempty"`
	Transitioning            string                `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage     string                `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`
	UUID                     string                `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	UpgradeFrom              string                `json:"upgradeFrom,omitempty" yaml:"upgrade_from,omitempty"`
	VersionLinks             map[string]string     `json:"versionLinks,omitempty" yaml:"version_links,omitempty"`
	Versions                 []TemplateVersionSpec `json:"versions,omitempty" yaml:"versions,omitempty"`
}

type CatalogTemplateCollection struct {
	types.Collection
	Data   []CatalogTemplate `json:"data,omitempty"`
	client *CatalogTemplateClient
}

type CatalogTemplateClient struct {
	apiClient *Client
}

type CatalogTemplateOperations interface {
	List(opts *types.ListOpts) (*CatalogTemplateCollection, error)
	Create(opts *CatalogTemplate) (*CatalogTemplate, error)
	Update(existing *CatalogTemplate, updates interface{}) (*CatalogTemplate, error)
	Replace(existing *CatalogTemplate) (*CatalogTemplate, error)
	ByID(id string) (*CatalogTemplate, error)
	Delete(container *CatalogTemplate) error
}

func newCatalogTemplateClient(apiClient *Client) *CatalogTemplateClient {
	return &CatalogTemplateClient{
		apiClient: apiClient,
	}
}

func (c *CatalogTemplateClient) Create(container *CatalogTemplate) (*CatalogTemplate, error) {
	resp := &CatalogTemplate{}
	err := c.apiClient.Ops.DoCreate(CatalogTemplateType, container, resp)
	return resp, err
}

func (c *CatalogTemplateClient) Update(existing *CatalogTemplate, updates interface{}) (*CatalogTemplate, error) {
	resp := &CatalogTemplate{}
	err := c.apiClient.Ops.DoUpdate(CatalogTemplateType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *CatalogTemplateClient) Replace(obj *CatalogTemplate) (*CatalogTemplate, error) {
	resp := &CatalogTemplate{}
	err := c.apiClient.Ops.DoReplace(CatalogTemplateType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *CatalogTemplateClient) List(opts *types.ListOpts) (*CatalogTemplateCollection, error) {
	resp := &CatalogTemplateCollection{}
	err := c.apiClient.Ops.DoList(CatalogTemplateType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *CatalogTemplateCollection) Next() (*CatalogTemplateCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &CatalogTemplateCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *CatalogTemplateClient) ByID(id string) (*CatalogTemplate, error) {
	resp := &CatalogTemplate{}
	err := c.apiClient.Ops.DoByID(CatalogTemplateType, id, resp)
	return resp, err
}

func (c *CatalogTemplateClient) Delete(container *CatalogTemplate) error {
	return c.apiClient.Ops.DoResourceDelete(CatalogTemplateType, &container.Resource)
}
