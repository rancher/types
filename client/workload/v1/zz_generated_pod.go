package client

import (
	"github.com/rancher/norman/types"
)

const (
	PodType                               = "pod"
	PodFieldAPIVersion                    = "apiVersion"
	PodFieldActiveDeadlineSeconds         = "activeDeadlineSeconds"
	PodFieldAnnotations                   = "annotations"
	PodFieldAutomountServiceAccountToken  = "automountServiceAccountToken"
	PodFieldContainers                    = "containers"
	PodFieldCreationTimestamp             = "creationTimestamp"
	PodFieldDNSPolicy                     = "dnsPolicy"
	PodFieldDeletionTimestamp             = "deletionTimestamp"
	PodFieldDeprecatedServiceAccount      = "deprecatedServiceAccount"
	PodFieldFSGroup                       = "fsGroup"
	PodFieldHostAliases                   = "hostAliases"
	PodFieldHostname                      = "hostname"
	PodFieldIPC                           = "ipc"
	PodFieldImagePullSecrets              = "imagePullSecrets"
	PodFieldKind                          = "kind"
	PodFieldLabels                        = "labels"
	PodFieldName                          = "name"
	PodFieldNamespace                     = "namespace"
	PodFieldNet                           = "net"
	PodFieldNodeName                      = "nodeName"
	PodFieldPID                           = "pid"
	PodFieldPriority                      = "priority"
	PodFieldPriorityClassName             = "priorityClassName"
	PodFieldRestartPolicy                 = "restartPolicy"
	PodFieldRunAsNonRoot                  = "runAsNonRoot"
	PodFieldRunAsUser                     = "runAsUser"
	PodFieldSchedulerName                 = "schedulerName"
	PodFieldServiceAccountName            = "serviceAccountName"
	PodFieldSubdomain                     = "subdomain"
	PodFieldSupplementalGroups            = "supplementalGroups"
	PodFieldTerminationGracePeriodSeconds = "terminationGracePeriodSeconds"
	PodFieldTolerations                   = "tolerations"
	PodFieldUID                           = "uid"
	PodFieldVolumes                       = "volumes"
)

type Pod struct {
	types.Resource
	APIVersion                    string                 `json:"apiVersion,omitempty"`
	ActiveDeadlineSeconds         *int64                 `json:"activeDeadlineSeconds,omitempty"`
	Annotations                   map[string]string      `json:"annotations,omitempty"`
	AutomountServiceAccountToken  *bool                  `json:"automountServiceAccountToken,omitempty"`
	Containers                    map[string]Container   `json:"containers,omitempty"`
	CreationTimestamp             string                 `json:"creationTimestamp,omitempty"`
	DNSPolicy                     string                 `json:"dnsPolicy,omitempty"`
	DeletionTimestamp             string                 `json:"deletionTimestamp,omitempty"`
	DeprecatedServiceAccount      string                 `json:"deprecatedServiceAccount,omitempty"`
	FSGroup                       *int64                 `json:"fsGroup,omitempty"`
	HostAliases                   map[string]HostAlias   `json:"hostAliases,omitempty"`
	Hostname                      string                 `json:"hostname,omitempty"`
	IPC                           string                 `json:"ipc,omitempty"`
	ImagePullSecrets              []LocalObjectReference `json:"imagePullSecrets,omitempty"`
	Kind                          string                 `json:"kind,omitempty"`
	Labels                        map[string]string      `json:"labels,omitempty"`
	Name                          string                 `json:"name,omitempty"`
	Namespace                     string                 `json:"namespace,omitempty"`
	Net                           string                 `json:"net,omitempty"`
	NodeName                      string                 `json:"nodeName,omitempty"`
	PID                           string                 `json:"pid,omitempty"`
	Priority                      *int64                 `json:"priority,omitempty"`
	PriorityClassName             string                 `json:"priorityClassName,omitempty"`
	RestartPolicy                 string                 `json:"restartPolicy,omitempty"`
	RunAsNonRoot                  *bool                  `json:"runAsNonRoot,omitempty"`
	RunAsUser                     *int64                 `json:"runAsUser,omitempty"`
	SchedulerName                 string                 `json:"schedulerName,omitempty"`
	ServiceAccountName            string                 `json:"serviceAccountName,omitempty"`
	Subdomain                     string                 `json:"subdomain,omitempty"`
	SupplementalGroups            []int64                `json:"supplementalGroups,omitempty"`
	TerminationGracePeriodSeconds *int64                 `json:"terminationGracePeriodSeconds,omitempty"`
	Tolerations                   []Toleration           `json:"tolerations,omitempty"`
	UID                           string                 `json:"uid,omitempty"`
	Volumes                       map[string]Volume      `json:"volumes,omitempty"`
}
type PodCollection struct {
	types.Collection
	Data   []Pod `json:"data,omitempty"`
	client *PodClient
}

type PodClient struct {
	apiClient *Client
}

type PodOperations interface {
	List(opts *types.ListOpts) (*PodCollection, error)
	Create(opts *Pod) (*Pod, error)
	Update(existing *Pod, updates interface{}) (*Pod, error)
	ByID(id string) (*Pod, error)
	Delete(container *Pod) error
}

func newPodClient(apiClient *Client) *PodClient {
	return &PodClient{
		apiClient: apiClient,
	}
}

func (c *PodClient) Create(container *Pod) (*Pod, error) {
	resp := &Pod{}
	err := c.apiClient.Ops.DoCreate(PodType, container, resp)
	return resp, err
}

func (c *PodClient) Update(existing *Pod, updates interface{}) (*Pod, error) {
	resp := &Pod{}
	err := c.apiClient.Ops.DoUpdate(PodType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *PodClient) List(opts *types.ListOpts) (*PodCollection, error) {
	resp := &PodCollection{}
	err := c.apiClient.Ops.DoList(PodType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *PodCollection) Next() (*PodCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &PodCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *PodClient) ByID(id string) (*Pod, error) {
	resp := &Pod{}
	err := c.apiClient.Ops.DoByID(PodType, id, resp)
	return resp, err
}

func (c *PodClient) Delete(container *Pod) error {
	return c.apiClient.Ops.DoResourceDelete(PodType, &container.Resource)
}
