package client

import (
	"github.com/rancher/norman/clientbase"
)

type Client struct {
	clientbase.APIBaseClient

	Cluster         ClusterOperations
	ClusterNode     ClusterNodeOperations
	Machine         MachineOperations
	MachineDriver   MachineDriverOperations
	MachineTemplate MachineTemplateOperations
	ClusterEvent    ClusterEventOperations
}

func NewClient(opts *clientbase.ClientOpts) (*Client, error) {
	baseClient, err := clientbase.NewAPIClient(opts)
	if err != nil {
		return nil, err
	}

	client := &Client{
		APIBaseClient: baseClient,
	}

	client.Cluster = newClusterClient(client)
	client.ClusterNode = newClusterNodeClient(client)
	client.Machine = newMachineClient(client)
	client.MachineDriver = newMachineDriverClient(client)
	client.MachineTemplate = newMachineTemplateClient(client)
	client.ClusterEvent = newClusterEventClient(client)

	return client, nil
}
