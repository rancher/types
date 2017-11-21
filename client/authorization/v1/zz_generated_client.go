package client

import (
	"github.com/rancher/norman/clientbase"
)

type Client struct {
	clientbase.APIBaseClient

	Project                    ProjectOperations
	ProjectRoleTemplate        ProjectRoleTemplateOperations
	PodSecurityPolicyTemplate  PodSecurityPolicyTemplateOperations
	ProjectRoleTemplateBinding ProjectRoleTemplateBindingOperations
	ClusterRoleTemplate        ClusterRoleTemplateOperations
	ClusterRoleTemplateBinding ClusterRoleTemplateBindingOperations
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
	client.ProjectRoleTemplate = newProjectRoleTemplateClient(client)
	client.PodSecurityPolicyTemplate = newPodSecurityPolicyTemplateClient(client)
	client.ProjectRoleTemplateBinding = newProjectRoleTemplateBindingClient(client)
	client.ClusterRoleTemplate = newClusterRoleTemplateClient(client)
	client.ClusterRoleTemplateBinding = newClusterRoleTemplateBindingClient(client)

	return client, nil
}
