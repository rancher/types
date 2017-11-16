//go:generate go run main.go

package main

import (
	authzSchema "github.com/rancher/types/apis/authorization.cattle.io/v1/schema"
	authnSchema "github.com/rancher/types/apis/authentication.cattle.io/v1/schema"
	clusterSchema "github.com/rancher/types/apis/cluster.cattle.io/v1/schema"
	workloadSchema "github.com/rancher/types/apis/workload.cattle.io/v1/schema"
	"github.com/rancher/types/generator"
)

func main() {
	generator.Generate(clusterSchema.Schemas)
	generator.Generate(workloadSchema.Schemas)
	generator.Generate(authzSchema.Schemas)
	generator.Generate(authnSchema.Schemas)
}
