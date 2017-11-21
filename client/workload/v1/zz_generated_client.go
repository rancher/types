package client

import (
	"github.com/rancher/norman/clientbase"
)

type Client struct {
	clientbase.APIBaseClient

	Pod                   PodOperations
	Namespace             NamespaceOperations
	Node                  NodeOperations
	Deployment            DeploymentOperations
	PersistentVolumeClaim PersistentVolumeClaimOperations
	StatefulSet           StatefulSetOperations
	ReplicaSet            ReplicaSetOperations
	ReplicationController ReplicationControllerOperations
	DaemonSet             DaemonSetOperations
	Workload              WorkloadOperations
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
	client.Deployment = newDeploymentClient(client)
	client.PersistentVolumeClaim = newPersistentVolumeClaimClient(client)
	client.StatefulSet = newStatefulSetClient(client)
	client.ReplicaSet = newReplicaSetClient(client)
	client.ReplicationController = newReplicationControllerClient(client)
	client.DaemonSet = newDaemonSetClient(client)
	client.Workload = newWorkloadClient(client)

	return client, nil
}
