//go:generate go run main.go

package main

import (
	"path"

	"strings"

	"github.com/rancher/norman/generator"
	"github.com/rancher/norman/types"
	clusterSchema "github.com/rancher/types/io.cattle.cluster/v1/schema"
)

var (
	basePackage = "github.com/rancher/types"
	baseCattle  = "client"
)

func main() {
	generate(clusterSchema.Schemas)
}

func generate(schemas *types.Schemas) {
	version := getVersion(schemas)
	groupParts := strings.Split(version.Group, ".")

	cattleOutputPackage := path.Join(basePackage, baseCattle, groupParts[len(groupParts)-1], version.Version)
	k8sOutputPackage := path.Join(basePackage, version.Group, version.Version)

	if err := generator.Generate(schemas, cattleOutputPackage, k8sOutputPackage); err != nil {
		panic(err)
	}
}

func getVersion(schemas *types.Schemas) *types.APIVersion {
	var version types.APIVersion
	for _, schema := range schemas.Schemas() {
		if version.Group == "" {
			version = schema.Version
			continue
		}
		if version.Group != schema.Version.Group ||
			version.Version != schema.Version.Version {
			panic("schema set contains two APIVersions")
		}
	}

	return &version
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
