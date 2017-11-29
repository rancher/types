//go:generate go run generator/cleanup/main.go
//go:generate go run main.go

package main

import (
	authzSchema "github.com/rancher/types/apis/authorization.cattle.io/v1/schema"
	clusterSchema "github.com/rancher/types/apis/cluster.cattle.io/v1/schema"
	workloadSchema "github.com/rancher/types/apis/workload.cattle.io/v1/schema"
	"github.com/rancher/types/generator"
	"k8s.io/api/apps/v1beta2"
	"k8s.io/api/core/v1"
)

func main() {
	generator.Generate(clusterSchema.Schemas)
	generator.Generate(workloadSchema.Schemas)
	generator.Generate(authzSchema.Schemas)
	// Group by API group
	generator.GenerateNativeTypes(v1.Pod{}, v1.Node{})
	generator.GenerateNativeTypes(v1beta2.Deployment{})
}
