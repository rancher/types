package generator

import (
	"fmt"
	"path"
	"reflect"
	"strings"

	"github.com/rancher/norman/generator"
	"github.com/rancher/norman/types"
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

func GenerateNativeTypes(objs ...interface{}) {
	pkgNamePaths := strings.Split(reflect.TypeOf(objs[0]).PkgPath(), "/")
	version := pkgNamePaths[len(pkgNamePaths)-1]
	group := pkgNamePaths[len(pkgNamePaths)-2]
	groupPath := group

	if group == "core" {
		group = ""
	}

	k8sOutputPackage := path.Join(basePackage, baseK8s, groupPath, version)

	if err := generator.GenerateControllerForTypes(&types.APIVersion{
		Version: version,
		Group:   group,
		Path:    fmt.Sprintf("/k8s/%s-%s", groupPath, version),
	}, k8sOutputPackage, objs...); err != nil {
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
