package client
import (
	"github.com/rancher/norman/types"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
    SourceCodeRepositoryType = "sourceCodeRepository"
	SourceCodeRepositoryFieldAnnotations = "annotations"
	SourceCodeRepositoryFieldCreated = "created"
	SourceCodeRepositoryFieldCreatorID = "creatorId"
	SourceCodeRepositoryFieldDefaultBranch = "defaultBranch"
	SourceCodeRepositoryFieldLabels = "labels"
	SourceCodeRepositoryFieldLanguage = "language"
	SourceCodeRepositoryFieldName = "name"
	SourceCodeRepositoryFieldOwnerReferences = "ownerReferences"
	SourceCodeRepositoryFieldPermissions = "permissions"
	SourceCodeRepositoryFieldProjectID = "projectId"
	SourceCodeRepositoryFieldRemoved = "removed"
	SourceCodeRepositoryFieldSourceCodeCredentialID = "sourceCodeCredentialId"
	SourceCodeRepositoryFieldSourceCodeType = "sourceCodeType"
	SourceCodeRepositoryFieldState = "state"
	SourceCodeRepositoryFieldStatus = "status"
	SourceCodeRepositoryFieldTransitioning = "transitioning"
	SourceCodeRepositoryFieldTransitioningMessage = "transitioningMessage"
	SourceCodeRepositoryFieldURL = "url"
	SourceCodeRepositoryFieldUUID = "uuid"
	SourceCodeRepositoryFieldUserID = "userId"
)

type SourceCodeRepository struct {
    types.Resource
        Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
        Created string `json:"created,omitempty" yaml:"created,omitempty"`
        CreatorID string `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
        DefaultBranch string `json:"defaultBranch,omitempty" yaml:"defaultBranch,omitempty"`
        Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
        Language string `json:"language,omitempty" yaml:"language,omitempty"`
        Name string `json:"name,omitempty" yaml:"name,omitempty"`
        OwnerReferences []OwnerReference `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
        Permissions *RepoPerm `json:"permissions,omitempty" yaml:"permissions,omitempty"`
        ProjectID string `json:"projectId,omitempty" yaml:"projectId,omitempty"`
        Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`
        SourceCodeCredentialID string `json:"sourceCodeCredentialId,omitempty" yaml:"sourceCodeCredentialId,omitempty"`
        SourceCodeType string `json:"sourceCodeType,omitempty" yaml:"sourceCodeType,omitempty"`
        State string `json:"state,omitempty" yaml:"state,omitempty"`
        Status *SourceCodeRepositoryStatus `json:"status,omitempty" yaml:"status,omitempty"`
        Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
        TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
        URL string `json:"url,omitempty" yaml:"url,omitempty"`
        UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
        UserID string `json:"userId,omitempty" yaml:"userId,omitempty"`
}


type SourceCodeRepositoryCollection struct {
    types.Collection
    Data []SourceCodeRepository `json:"data,omitempty"`
    client *SourceCodeRepositoryClient
}

type SourceCodeRepositoryClient struct {
    apiClient *Client
}

type SourceCodeRepositoryOperations interface {
    List(opts *types.ListOpts) (*SourceCodeRepositoryCollection, error)
    Create(opts *SourceCodeRepository) (*SourceCodeRepository, error)
    Update(existing *SourceCodeRepository, updates interface{}) (*SourceCodeRepository, error)
    Replace(existing *SourceCodeRepository) (*SourceCodeRepository, error)
    ByID(id string) (*SourceCodeRepository, error)
    Delete(container *SourceCodeRepository) error
    
    
}

func newSourceCodeRepositoryClient(apiClient *Client) *SourceCodeRepositoryClient {
    return &SourceCodeRepositoryClient{
        apiClient: apiClient,
    }
}

func (c *SourceCodeRepositoryClient) Create(container *SourceCodeRepository) (*SourceCodeRepository, error) {
    resp := &SourceCodeRepository{}
    err := c.apiClient.Ops.DoCreate(SourceCodeRepositoryType, container, resp)
    return resp, err
}

func (c *SourceCodeRepositoryClient) Update(existing *SourceCodeRepository, updates interface{}) (*SourceCodeRepository, error) {
    resp := &SourceCodeRepository{}
    err := c.apiClient.Ops.DoUpdate(SourceCodeRepositoryType, &existing.Resource, updates, resp)
    return resp, err
}

func (c *SourceCodeRepositoryClient) Replace(obj *SourceCodeRepository) (*SourceCodeRepository, error) {
	resp := &SourceCodeRepository{}
	err := c.apiClient.Ops.DoReplace(SourceCodeRepositoryType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *SourceCodeRepositoryClient) List(opts *types.ListOpts) (*SourceCodeRepositoryCollection, error) {
    resp := &SourceCodeRepositoryCollection{}
    err := c.apiClient.Ops.DoList(SourceCodeRepositoryType, opts, resp)
    resp.client = c
    return resp, err
}

func (cc *SourceCodeRepositoryCollection) Next() (*SourceCodeRepositoryCollection, error) {
    if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
        resp := &SourceCodeRepositoryCollection{}
        err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
        resp.client = cc.client
        return resp, err
    }
    return nil, nil
}

func (c *SourceCodeRepositoryClient) ByID(id string) (*SourceCodeRepository, error) {
    resp := &SourceCodeRepository{}
    err := c.apiClient.Ops.DoByID(SourceCodeRepositoryType, id, resp)
    return resp, err
}

func (c *SourceCodeRepositoryClient) Delete(container *SourceCodeRepository) error {
    return c.apiClient.Ops.DoResourceDelete(SourceCodeRepositoryType, &container.Resource)
}




