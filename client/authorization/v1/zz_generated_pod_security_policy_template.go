package client

import (
	"github.com/rancher/norman/types"
)

const (
	PodSecurityPolicyTemplateType            = "podSecurityPolicyTemplate"
	PodSecurityPolicyTemplateFieldAPIVersion = "apiVersion"
	PodSecurityPolicyTemplateFieldKind       = "kind"
	PodSecurityPolicyTemplateFieldObjectMeta = "objectMeta"
	PodSecurityPolicyTemplateFieldSpec       = "spec"
)

type PodSecurityPolicyTemplate struct {
	types.Resource
	APIVersion string                `json:"apiVersion,omitempty"`
	Kind       string                `json:"kind,omitempty"`
	ObjectMeta ObjectMeta            `json:"objectMeta,omitempty"`
	Spec       PodSecurityPolicySpec `json:"spec,omitempty"`
}
type PodSecurityPolicyTemplateCollection struct {
	types.Collection
	Data   []PodSecurityPolicyTemplate `json:"data,omitempty"`
	client *PodSecurityPolicyTemplateClient
}

type PodSecurityPolicyTemplateClient struct {
	apiClient *Client
}

type PodSecurityPolicyTemplateOperations interface {
	List(opts *types.ListOpts) (*PodSecurityPolicyTemplateCollection, error)
	Create(opts *PodSecurityPolicyTemplate) (*PodSecurityPolicyTemplate, error)
	Update(existing *PodSecurityPolicyTemplate, updates interface{}) (*PodSecurityPolicyTemplate, error)
	ByID(id string) (*PodSecurityPolicyTemplate, error)
	Delete(container *PodSecurityPolicyTemplate) error
}

func newPodSecurityPolicyTemplateClient(apiClient *Client) *PodSecurityPolicyTemplateClient {
	return &PodSecurityPolicyTemplateClient{
		apiClient: apiClient,
	}
}

func (c *PodSecurityPolicyTemplateClient) Create(container *PodSecurityPolicyTemplate) (*PodSecurityPolicyTemplate, error) {
	resp := &PodSecurityPolicyTemplate{}
	err := c.apiClient.Ops.DoCreate(PodSecurityPolicyTemplateType, container, resp)
	return resp, err
}

func (c *PodSecurityPolicyTemplateClient) Update(existing *PodSecurityPolicyTemplate, updates interface{}) (*PodSecurityPolicyTemplate, error) {
	resp := &PodSecurityPolicyTemplate{}
	err := c.apiClient.Ops.DoUpdate(PodSecurityPolicyTemplateType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *PodSecurityPolicyTemplateClient) List(opts *types.ListOpts) (*PodSecurityPolicyTemplateCollection, error) {
	resp := &PodSecurityPolicyTemplateCollection{}
	err := c.apiClient.Ops.DoList(PodSecurityPolicyTemplateType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *PodSecurityPolicyTemplateCollection) Next() (*PodSecurityPolicyTemplateCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &PodSecurityPolicyTemplateCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *PodSecurityPolicyTemplateClient) ByID(id string) (*PodSecurityPolicyTemplate, error) {
	resp := &PodSecurityPolicyTemplate{}
	err := c.apiClient.Ops.DoByID(PodSecurityPolicyTemplateType, id, resp)
	return resp, err
}

func (c *PodSecurityPolicyTemplateClient) Delete(container *PodSecurityPolicyTemplate) error {
	return c.apiClient.Ops.DoResourceDelete(PodSecurityPolicyTemplateType, &container.Resource)
}
