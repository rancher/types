package client

import (
	"github.com/rancher/norman/clientbase"
)

type Client struct {
	clientbase.APIBaseClient

	Project                   ProjectOperations
	RoleTemplate              RoleTemplateOperations
	PodSecurityPolicyTemplate PodSecurityPolicyTemplateOperations
	ProjectRoleBinding        ProjectRoleBindingOperations
}

func NewClient(opts *clientbase.ClientOpts) (*Client, error) {
	baseClient, err := clientbase.NewAPIClient(opts)
	if err != nil {
		return nil, err
	}

	client := &Client{
		APIBaseClient: baseClient,
	}

	client.Project = newProjectClient(client)
	client.RoleTemplate = newRoleTemplateClient(client)
	client.PodSecurityPolicyTemplate = newPodSecurityPolicyTemplateClient(client)
	client.ProjectRoleBinding = newProjectRoleBindingClient(client)

	return client, nil
}
