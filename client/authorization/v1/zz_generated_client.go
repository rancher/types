package client

import (
	"github.com/rancher/norman/clientbase"
)

type Client struct {
	clientbase.APIBaseClient

	Project                    ProjectOperations
	ProjectRoleTemplate        ProjectRoleTemplateOperations
	PodSecurityPolicyTemplate  PodSecurityPolicyTemplateOperations
	ClusterRoleTemplate        ClusterRoleTemplateOperations
	ClusterRoleTemplateBinding ClusterRoleTemplateBindingOperations
	ProjectRoleTemplateBinding ProjectRoleTemplateBindingOperations
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
	client.ClusterRoleTemplate = newClusterRoleTemplateClient(client)
	client.ClusterRoleTemplateBinding = newClusterRoleTemplateBindingClient(client)
	client.ProjectRoleTemplateBinding = newProjectRoleTemplateBindingClient(client)

	return client, nil
}
