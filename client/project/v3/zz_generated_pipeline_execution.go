package client

import (
	"github.com/rancher/norman/types"
)

const (
	PipelineExecutionType                      = "pipelineExecution"
	PipelineExecutionFieldAnnotations          = "annotations"
	PipelineExecutionFieldAuthor               = "author"
	PipelineExecutionFieldAvatarURL            = "avatarUrl"
	PipelineExecutionFieldBranch               = "branch"
	PipelineExecutionFieldCommit               = "commit"
	PipelineExecutionFieldConditions           = "conditions"
	PipelineExecutionFieldCreated              = "created"
	PipelineExecutionFieldCreatorID            = "creatorId"
	PipelineExecutionFieldEmail                = "email"
	PipelineExecutionFieldEnded                = "ended"
	PipelineExecutionFieldEvent                = "event"
	PipelineExecutionFieldExecutionState       = "executionState"
	PipelineExecutionFieldHTMLLink             = "htmlLink"
	PipelineExecutionFieldLabels               = "labels"
	PipelineExecutionFieldMessage              = "message"
	PipelineExecutionFieldName                 = "name"
	PipelineExecutionFieldNamespaceId          = "namespaceId"
	PipelineExecutionFieldOwnerReferences      = "ownerReferences"
	PipelineExecutionFieldPipelineConfig       = "pipelineConfig"
	PipelineExecutionFieldPipelineID           = "pipelineId"
	PipelineExecutionFieldProjectID            = "projectId"
	PipelineExecutionFieldRef                  = "ref"
	PipelineExecutionFieldRemoved              = "removed"
	PipelineExecutionFieldRepositoryURL        = "repositoryUrl"
	PipelineExecutionFieldRun                  = "run"
	PipelineExecutionFieldStages               = "stages"
	PipelineExecutionFieldStarted              = "started"
	PipelineExecutionFieldState                = "state"
	PipelineExecutionFieldTitle                = "title"
	PipelineExecutionFieldTransitioning        = "transitioning"
	PipelineExecutionFieldTransitioningMessage = "transitioningMessage"
	PipelineExecutionFieldTriggerUserID        = "triggerUserId"
	PipelineExecutionFieldTriggeredBy          = "triggeredBy"
	PipelineExecutionFieldUUID                 = "uuid"
)

type PipelineExecution struct {
	types.Resource
	Annotations          map[string]string   `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Author               string              `json:"author,omitempty" yaml:"author,omitempty"`
	AvatarURL            string              `json:"avatarUrl,omitempty" yaml:"avatar_url,omitempty"`
	Branch               string              `json:"branch,omitempty" yaml:"branch,omitempty"`
	Commit               string              `json:"commit,omitempty" yaml:"commit,omitempty"`
	Conditions           []PipelineCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	Created              string              `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string              `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	Email                string              `json:"email,omitempty" yaml:"email,omitempty"`
	Ended                string              `json:"ended,omitempty" yaml:"ended,omitempty"`
	Event                string              `json:"event,omitempty" yaml:"event,omitempty"`
	ExecutionState       string              `json:"executionState,omitempty" yaml:"execution_state,omitempty"`
	HTMLLink             string              `json:"htmlLink,omitempty" yaml:"html_link,omitempty"`
	Labels               map[string]string   `json:"labels,omitempty" yaml:"labels,omitempty"`
	Message              string              `json:"message,omitempty" yaml:"message,omitempty"`
	Name                 string              `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId          string              `json:"namespaceId,omitempty" yaml:"namespace_id,omitempty"`
	OwnerReferences      []OwnerReference    `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	PipelineConfig       *PipelineConfig     `json:"pipelineConfig,omitempty" yaml:"pipeline_config,omitempty"`
	PipelineID           string              `json:"pipelineId,omitempty" yaml:"pipeline_id,omitempty"`
	ProjectID            string              `json:"projectId,omitempty" yaml:"project_id,omitempty"`
	Ref                  string              `json:"ref,omitempty" yaml:"ref,omitempty"`
	Removed              string              `json:"removed,omitempty" yaml:"removed,omitempty"`
	RepositoryURL        string              `json:"repositoryUrl,omitempty" yaml:"repository_url,omitempty"`
	Run                  int64               `json:"run,omitempty" yaml:"run,omitempty"`
	Stages               []StageStatus       `json:"stages,omitempty" yaml:"stages,omitempty"`
	Started              string              `json:"started,omitempty" yaml:"started,omitempty"`
	State                string              `json:"state,omitempty" yaml:"state,omitempty"`
	Title                string              `json:"title,omitempty" yaml:"title,omitempty"`
	Transitioning        string              `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage string              `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`
	TriggerUserID        string              `json:"triggerUserId,omitempty" yaml:"trigger_user_id,omitempty"`
	TriggeredBy          string              `json:"triggeredBy,omitempty" yaml:"triggered_by,omitempty"`
	UUID                 string              `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type PipelineExecutionCollection struct {
	types.Collection
	Data   []PipelineExecution `json:"data,omitempty"`
	client *PipelineExecutionClient
}

type PipelineExecutionClient struct {
	apiClient *Client
}

type PipelineExecutionOperations interface {
	List(opts *types.ListOpts) (*PipelineExecutionCollection, error)
	Create(opts *PipelineExecution) (*PipelineExecution, error)
	Update(existing *PipelineExecution, updates interface{}) (*PipelineExecution, error)
	Replace(existing *PipelineExecution) (*PipelineExecution, error)
	ByID(id string) (*PipelineExecution, error)
	Delete(container *PipelineExecution) error

	ActionRerun(resource *PipelineExecution) error

	ActionStop(resource *PipelineExecution) error
}

func newPipelineExecutionClient(apiClient *Client) *PipelineExecutionClient {
	return &PipelineExecutionClient{
		apiClient: apiClient,
	}
}

func (c *PipelineExecutionClient) Create(container *PipelineExecution) (*PipelineExecution, error) {
	resp := &PipelineExecution{}
	err := c.apiClient.Ops.DoCreate(PipelineExecutionType, container, resp)
	return resp, err
}

func (c *PipelineExecutionClient) Update(existing *PipelineExecution, updates interface{}) (*PipelineExecution, error) {
	resp := &PipelineExecution{}
	err := c.apiClient.Ops.DoUpdate(PipelineExecutionType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *PipelineExecutionClient) Replace(obj *PipelineExecution) (*PipelineExecution, error) {
	resp := &PipelineExecution{}
	err := c.apiClient.Ops.DoReplace(PipelineExecutionType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *PipelineExecutionClient) List(opts *types.ListOpts) (*PipelineExecutionCollection, error) {
	resp := &PipelineExecutionCollection{}
	err := c.apiClient.Ops.DoList(PipelineExecutionType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *PipelineExecutionCollection) Next() (*PipelineExecutionCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &PipelineExecutionCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *PipelineExecutionClient) ByID(id string) (*PipelineExecution, error) {
	resp := &PipelineExecution{}
	err := c.apiClient.Ops.DoByID(PipelineExecutionType, id, resp)
	return resp, err
}

func (c *PipelineExecutionClient) Delete(container *PipelineExecution) error {
	return c.apiClient.Ops.DoResourceDelete(PipelineExecutionType, &container.Resource)
}

func (c *PipelineExecutionClient) ActionRerun(resource *PipelineExecution) error {
	err := c.apiClient.Ops.DoAction(PipelineExecutionType, "rerun", &resource.Resource, nil, nil)
	return err
}

func (c *PipelineExecutionClient) ActionStop(resource *PipelineExecution) error {
	err := c.apiClient.Ops.DoAction(PipelineExecutionType, "stop", &resource.Resource, nil, nil)
	return err
}
