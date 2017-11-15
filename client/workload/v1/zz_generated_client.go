package client

import (
	"github.com/rancher/norman/clientbase"
)

type Client struct {
	clientbase.APIBaseClient

	Pod                   PodOperations
	Namespace             NamespaceOperations
	Node                  NodeOperations
	ReplicaSet            ReplicaSetOperations
	Deployment            DeploymentOperations
	PersistentVolumeClaim PersistentVolumeClaimOperations
	StatefulSet           StatefulSetOperations
}

func NewClient(opts *clientbase.ClientOpts) (*Client, error) {
	baseClient, err := clientbase.NewAPIClient(opts)
	if err != nil {
		return nil, err
	}

	client := &Client{
		APIBaseClient: baseClient,
	}

	client.Pod = newPodClient(client)
	client.Namespace = newNamespaceClient(client)
	client.Node = newNodeClient(client)
	client.ReplicaSet = newReplicaSetClient(client)
	client.Deployment = newDeploymentClient(client)
	client.PersistentVolumeClaim = newPersistentVolumeClaimClient(client)
	client.StatefulSet = newStatefulSetClient(client)

	return client, nil
}
