//go:generate go run main.go

package main

import (
	"github.com/rancher/types/generator"
	authzSchema "github.com/rancher/types/io.cattle.authorization/v1/schema"
	clusterSchema "github.com/rancher/types/io.cattle.cluster/v1/schema"
	workloadSchema "github.com/rancher/types/io.cattle.workload/v1/schema"
)

func main() {
	generator.Generate(clusterSchema.Schemas)
	generator.Generate(workloadSchema.Schemas)
	generator.Generate(authzSchema.Schemas)
}
