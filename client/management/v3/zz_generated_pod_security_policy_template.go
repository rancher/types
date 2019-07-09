package client

import (
	"github.com/rancher/norman/types"
)

const (
	PodSecurityPolicyTemplateType                                 = "podSecurityPolicyTemplate"
	PodSecurityPolicyTemplateFieldAllowPrivilegeEscalation        = "allowPrivilegeEscalation"
	PodSecurityPolicyTemplateFieldAllowedCapabilities             = "allowedCapabilities"
	PodSecurityPolicyTemplateFieldAllowedFlexVolumes              = "allowedFlexVolumes"
	PodSecurityPolicyTemplateFieldAllowedHostPaths                = "allowedHostPaths"
	PodSecurityPolicyTemplateFieldAllowedProcMountTypes           = "allowedProcMountTypes"
	PodSecurityPolicyTemplateFieldAllowedUnsafeSysctls            = "allowedUnsafeSysctls"
	PodSecurityPolicyTemplateFieldAnnotations                     = "annotations"
	PodSecurityPolicyTemplateFieldCreated                         = "created"
	PodSecurityPolicyTemplateFieldCreatorID                       = "creatorId"
	PodSecurityPolicyTemplateFieldDefaultAddCapabilities          = "defaultAddCapabilities"
	PodSecurityPolicyTemplateFieldDefaultAllowPrivilegeEscalation = "defaultAllowPrivilegeEscalation"
	PodSecurityPolicyTemplateFieldDescription                     = "description"
	PodSecurityPolicyTemplateFieldFSGroup                         = "fsGroup"
	PodSecurityPolicyTemplateFieldForbiddenSysctls                = "forbiddenSysctls"
	PodSecurityPolicyTemplateFieldHostIPC                         = "hostIPC"
	PodSecurityPolicyTemplateFieldHostNetwork                     = "hostNetwork"
	PodSecurityPolicyTemplateFieldHostPID                         = "hostPID"
	PodSecurityPolicyTemplateFieldHostPorts                       = "hostPorts"
	PodSecurityPolicyTemplateFieldLabels                          = "labels"
	PodSecurityPolicyTemplateFieldName                            = "name"
	PodSecurityPolicyTemplateFieldOwnerReferences                 = "ownerReferences"
	PodSecurityPolicyTemplateFieldPrivileged                      = "privileged"
	PodSecurityPolicyTemplateFieldReadOnlyRootFilesystem          = "readOnlyRootFilesystem"
	PodSecurityPolicyTemplateFieldRemoved                         = "removed"
	PodSecurityPolicyTemplateFieldRequiredDropCapabilities        = "requiredDropCapabilities"
	PodSecurityPolicyTemplateFieldRunAsUser                       = "runAsUser"
	PodSecurityPolicyTemplateFieldSELinux                         = "seLinux"
	PodSecurityPolicyTemplateFieldSupplementalGroups              = "supplementalGroups"
	PodSecurityPolicyTemplateFieldUUID                            = "uuid"
	PodSecurityPolicyTemplateFieldVolumes                         = "volumes"
)

type PodSecurityPolicyTemplate struct {
	types.Resource
	AllowPrivilegeEscalation        *bool                              `json:"allowPrivilegeEscalation,omitempty" yaml:"allow_privilege_escalation,omitempty"`
	AllowedCapabilities             []string                           `json:"allowedCapabilities,omitempty" yaml:"allowed_capabilities,omitempty"`
	AllowedFlexVolumes              []AllowedFlexVolume                `json:"allowedFlexVolumes,omitempty" yaml:"allowed_flex_volumes,omitempty"`
	AllowedHostPaths                []AllowedHostPath                  `json:"allowedHostPaths,omitempty" yaml:"allowed_host_paths,omitempty"`
	AllowedProcMountTypes           []string                           `json:"allowedProcMountTypes,omitempty" yaml:"allowed_proc_mount_types,omitempty"`
	AllowedUnsafeSysctls            []string                           `json:"allowedUnsafeSysctls,omitempty" yaml:"allowed_unsafe_sysctls,omitempty"`
	Annotations                     map[string]string                  `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created                         string                             `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                       string                             `json:"creatorId,omitempty" yaml:"creator_id,omitempty"`
	DefaultAddCapabilities          []string                           `json:"defaultAddCapabilities,omitempty" yaml:"default_add_capabilities,omitempty"`
	DefaultAllowPrivilegeEscalation *bool                              `json:"defaultAllowPrivilegeEscalation,omitempty" yaml:"default_allow_privilege_escalation,omitempty"`
	Description                     string                             `json:"description,omitempty" yaml:"description,omitempty"`
	FSGroup                         *FSGroupStrategyOptions            `json:"fsGroup,omitempty" yaml:"fs_group,omitempty"`
	ForbiddenSysctls                []string                           `json:"forbiddenSysctls,omitempty" yaml:"forbidden_sysctls,omitempty"`
	HostIPC                         bool                               `json:"hostIPC,omitempty" yaml:"host_ipc,omitempty"`
	HostNetwork                     bool                               `json:"hostNetwork,omitempty" yaml:"host_network,omitempty"`
	HostPID                         bool                               `json:"hostPID,omitempty" yaml:"host_pid,omitempty"`
	HostPorts                       []HostPortRange                    `json:"hostPorts,omitempty" yaml:"host_ports,omitempty"`
	Labels                          map[string]string                  `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                            string                             `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences                 []OwnerReference                   `json:"ownerReferences,omitempty" yaml:"owner_references,omitempty"`
	Privileged                      bool                               `json:"privileged,omitempty" yaml:"privileged,omitempty"`
	ReadOnlyRootFilesystem          bool                               `json:"readOnlyRootFilesystem,omitempty" yaml:"read_only_root_filesystem,omitempty"`
	Removed                         string                             `json:"removed,omitempty" yaml:"removed,omitempty"`
	RequiredDropCapabilities        []string                           `json:"requiredDropCapabilities,omitempty" yaml:"required_drop_capabilities,omitempty"`
	RunAsUser                       *RunAsUserStrategyOptions          `json:"runAsUser,omitempty" yaml:"run_as_user,omitempty"`
	SELinux                         *SELinuxStrategyOptions            `json:"seLinux,omitempty" yaml:"se_linux,omitempty"`
	SupplementalGroups              *SupplementalGroupsStrategyOptions `json:"supplementalGroups,omitempty" yaml:"supplemental_groups,omitempty"`
	UUID                            string                             `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Volumes                         []string                           `json:"volumes,omitempty" yaml:"volumes,omitempty"`
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
	Replace(existing *PodSecurityPolicyTemplate) (*PodSecurityPolicyTemplate, error)
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

func (c *PodSecurityPolicyTemplateClient) Replace(obj *PodSecurityPolicyTemplate) (*PodSecurityPolicyTemplate, error) {
	resp := &PodSecurityPolicyTemplate{}
	err := c.apiClient.Ops.DoReplace(PodSecurityPolicyTemplateType, &obj.Resource, obj, resp)
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
