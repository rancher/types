//go:generate go run generator/cleanup/main.go
//go:generate go run main.go

package main

import (
	clusterSchema "github.com/rancher/types/apis/cluster.cattle.io/v3/schema"
	managementSchema "github.com/rancher/types/apis/management.cattle.io/v3/schema"
	projectSchema "github.com/rancher/types/apis/project.cattle.io/v3/schema"
	"github.com/rancher/types/generator"
	"k8s.io/api/apps/v1beta2"
	"k8s.io/api/core/v1"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
)

func main() {
	generator.Generate(managementSchema.Schemas)
	generator.Generate(clusterSchema.Schemas)
	generator.Generate(projectSchema.Schemas)
	// Group by API group
	generator.GenerateNativeTypes(v1.SchemeGroupVersion, v1.Pod{}, v1.Node{}, v1.ComponentStatus{}, v1.Namespace{}, v1.Event{})
	generator.GenerateNativeTypes(v1beta2.SchemeGroupVersion, v1beta2.Deployment{})
	generator.GenerateNativeTypes(rbacv1.SchemeGroupVersion, rbacv1.RoleBinding{}, rbacv1.ClusterRoleBinding{}, rbacv1.ClusterRole{},
		rbacv1.Role{})
	generator.GenerateNativeTypes(extv1beta1.SchemeGroupVersion, extv1beta1.PodSecurityPolicy{})
}
