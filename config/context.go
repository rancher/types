package config

import (
	"context"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/signal"
	appsv1beta2 "github.com/rancher/types/apis/apps/v1beta2"
	authzv1 "github.com/rancher/types/apis/authorization.cattle.io/v1"
	clusterv1 "github.com/rancher/types/apis/cluster.cattle.io/v1"
	corev1 "github.com/rancher/types/apis/core/v1"
	workloadv1 "github.com/rancher/types/apis/workload.cattle.io/v1"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

type ClusterContext struct {
	RESTConfig        rest.Config
	UnversionedClient rest.Interface
	Cluster           clusterv1.Interface
	Authorization     authzv1.Interface
}

type WorkloadContext struct {
	Cluster           *ClusterContext
	RESTConfig        rest.Config
	UnversionedClient rest.Interface
	Apps              appsv1beta2.Interface
	Workload          workloadv1.Interface
	Core              corev1.Interface
}

func NewClusterContext(config rest.Config) (*ClusterContext, error) {
	var err error

	context := &ClusterContext{
		RESTConfig: config,
	}

	context.Cluster, err = clusterv1.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	context.Authorization, err = authzv1.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	dynamicConfig := config
	if dynamicConfig.NegotiatedSerializer == nil {
		configConfig := dynamic.ContentConfig()
		dynamicConfig.NegotiatedSerializer = configConfig.NegotiatedSerializer
	}

	context.UnversionedClient, err = rest.UnversionedRESTClientFor(&dynamicConfig)
	if err != nil {
		return nil, err
	}

	return context, err
}

func (c *ClusterContext) Start(ctx context.Context) error {
	logrus.Info("Syncing cluster controllers")
	err := controller.Sync(ctx,
		c.Cluster,
		c.Authorization)
	if err != nil {
		return err
	}

	logrus.Info("Starting cluster controllers")
	if err := c.Cluster.Start(ctx, 5); err != nil {
		return err
	}

	if err := c.Authorization.Start(ctx, 5); err != nil {
		return err
	}

	logrus.Info("Cluster context started")
	return nil
}

func (c *ClusterContext) StartAndWait() error {
	ctx := signal.SigTermCancelContext(context.Background())
	c.Start(ctx)
	<-ctx.Done()
	return ctx.Err()
}

func NewWorkloadContext(clusterConfig, config rest.Config) (*WorkloadContext, error) {
	var err error
	context := &WorkloadContext{
		RESTConfig: config,
	}

	context.Cluster, err = NewClusterContext(clusterConfig)
	if err != nil {
		return nil, err
	}

	context.Apps, err = appsv1beta2.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	context.Workload, err = workloadv1.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	context.Core, err = corev1.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	dynamicConfig := config
	if dynamicConfig.NegotiatedSerializer == nil {
		configConfig := dynamic.ContentConfig()
		dynamicConfig.NegotiatedSerializer = configConfig.NegotiatedSerializer
	}

	context.UnversionedClient, err = rest.UnversionedRESTClientFor(&dynamicConfig)
	if err != nil {
		return nil, err
	}

	return context, err
}
