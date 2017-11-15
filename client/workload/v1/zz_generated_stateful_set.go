package client

import (
	"github.com/rancher/norman/types"
)

const (
	StatefulSetType                      = "statefulSet"
	StatefulSetFieldAPIVersion           = "apiVersion"
	StatefulSetFieldAnnotations          = "annotations"
	StatefulSetFieldCreationTimestamp    = "creationTimestamp"
	StatefulSetFieldDeletionTimestamp    = "deletionTimestamp"
	StatefulSetFieldDeploy               = "deploy"
	StatefulSetFieldKind                 = "kind"
	StatefulSetFieldLabels               = "labels"
	StatefulSetFieldName                 = "name"
	StatefulSetFieldNamespace            = "namespace"
	StatefulSetFieldServiceName          = "serviceName"
	StatefulSetFieldStatus               = "status"
	StatefulSetFieldTemplate             = "template"
	StatefulSetFieldUID                  = "uid"
	StatefulSetFieldUpdateStrategy       = "updateStrategy"
	StatefulSetFieldVolumeClaimTemplates = "volumeClaimTemplates"
)

type StatefulSet struct {
	types.Resource
	APIVersion           string                           `json:"apiVersion,omitempty"`
	Annotations          map[string]string                `json:"annotations,omitempty"`
	CreationTimestamp    string                           `json:"creationTimestamp,omitempty"`
	DeletionTimestamp    string                           `json:"deletionTimestamp,omitempty"`
	Deploy               *DeployParams                    `json:"deploy,omitempty"`
	Kind                 string                           `json:"kind,omitempty"`
	Labels               map[string]string                `json:"labels,omitempty"`
	Name                 string                           `json:"name,omitempty"`
	Namespace            string                           `json:"namespace,omitempty"`
	ServiceName          string                           `json:"serviceName,omitempty"`
	Status               StatefulSetStatus                `json:"status,omitempty"`
	Template             PodTemplateSpec                  `json:"template,omitempty"`
	UID                  string                           `json:"uid,omitempty"`
	UpdateStrategy       StatefulSetUpdateStrategy        `json:"updateStrategy,omitempty"`
	VolumeClaimTemplates map[string]PersistentVolumeClaim `json:"volumeClaimTemplates,omitempty"`
}
type StatefulSetCollection struct {
	types.Collection
	Data   []StatefulSet `json:"data,omitempty"`
	client *StatefulSetClient
}

type StatefulSetClient struct {
	apiClient *Client
}

type StatefulSetOperations interface {
	List(opts *types.ListOpts) (*StatefulSetCollection, error)
	Create(opts *StatefulSet) (*StatefulSet, error)
	Update(existing *StatefulSet, updates interface{}) (*StatefulSet, error)
	ByID(id string) (*StatefulSet, error)
	Delete(container *StatefulSet) error
}

func newStatefulSetClient(apiClient *Client) *StatefulSetClient {
	return &StatefulSetClient{
		apiClient: apiClient,
	}
}

func (c *StatefulSetClient) Create(container *StatefulSet) (*StatefulSet, error) {
	resp := &StatefulSet{}
	err := c.apiClient.Ops.DoCreate(StatefulSetType, container, resp)
	return resp, err
}

func (c *StatefulSetClient) Update(existing *StatefulSet, updates interface{}) (*StatefulSet, error) {
	resp := &StatefulSet{}
	err := c.apiClient.Ops.DoUpdate(StatefulSetType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *StatefulSetClient) List(opts *types.ListOpts) (*StatefulSetCollection, error) {
	resp := &StatefulSetCollection{}
	err := c.apiClient.Ops.DoList(StatefulSetType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *StatefulSetCollection) Next() (*StatefulSetCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &StatefulSetCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *StatefulSetClient) ByID(id string) (*StatefulSet, error) {
	resp := &StatefulSet{}
	err := c.apiClient.Ops.DoByID(StatefulSetType, id, resp)
	return resp, err
}

func (c *StatefulSetClient) Delete(container *StatefulSet) error {
	return c.apiClient.Ops.DoResourceDelete(StatefulSetType, &container.Resource)
}
