package generator

import (
	"fmt"
	"path"
	"strings"

	"github.com/rancher/norman/generator"
	"github.com/rancher/norman/types"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	basePackage = "github.com/rancher/types"
	baseCattle  = "client"
	baseK8s     = "apis"
)

func Generate(schemas *types.Schemas) {
	version := getVersion(schemas)
	group := strings.Split(version.Group, ".")[0]

	cattleOutputPackage := path.Join(basePackage, baseCattle, group, version.Version)
	k8sOutputPackage := path.Join(basePackage, baseK8s, version.Group, version.Version)

	if err := generator.Generate(schemas, cattleOutputPackage, k8sOutputPackage); err != nil {
		panic(err)
	}
}

func GenerateNativeTypes(gv schema.GroupVersion, nsObjs []interface{}, objs []interface{}) {
	version := gv.Version
	group := gv.Group
	groupPath := group

	if groupPath == "" {
		groupPath = "core"
	}

	k8sOutputPackage := path.Join(basePackage, baseK8s, groupPath, version)

	if err := generator.GenerateControllerForTypes(&types.APIVersion{
		Version: version,
		Group:   group,
		Path:    fmt.Sprintf("/k8s/%s-%s", groupPath, version),
	}, k8sOutputPackage, nsObjs, objs); err != nil {
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
