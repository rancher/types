package image

import "strings"

var Mirrors = map[string]string{}

var (
	ManifestList = []string{
		"quay.io/coreos/etcd:v3.2.24",
		"gcr.io/google_containers/k8s-dns-kube-dns:1.15.0",
		"gcr.io/google_containers/k8s-dns-dnsmasq-nanny:1.15.0",
		"gcr.io/google_containers/k8s-dns-sidecar:1.15.0",
		"gcr.io/google_containers/k8s-dns-kube-dns:1.14.13",
		"gcr.io/google_containers/k8s-dns-dnsmasq-nanny:1.14.13",
		"gcr.io/google_containers/k8s-dns-sidecar:1.14.13",
		"gcr.io/google_containers/cluster-proportional-autoscaler:1.0.0",
		"quay.io/coreos/flannel:v0.10.0",
		"gcr.io/google_containers/pause:3.1",
		"gcr.io/google_containers/metrics-server:v0.3.1",
		"k8s.gcr.io/defaultbackend:1.4",
	}
	ArchitectureList = []string{
		"amd64",
		"arm64",
	}
)

func Mirror(image string) string {
	orig := image
	if strings.HasPrefix(image, "weaveworks") {
		return image
	}

	image = strings.Replace(image, "gcr.io/google_containers", "rancher", 1)
	image = strings.Replace(image, "quay.io/coreos/", "rancher/coreos-", 1)
	image = strings.Replace(image, "quay.io/calico/", "rancher/calico-", 1)
	image = strings.Replace(image, "k8s.gcr.io/", "rancher/nginx-ingress-controller-", 1)
	image = strings.Replace(image, "plugins/docker", "rancher/jenkins-plugins-docker", 1)
	image = strings.Replace(image, "kibana", "rancher/kibana", 1)
	image = strings.Replace(image, "jenkins/", "rancher/jenkins-", 1)
	image = strings.Replace(image, "alpine/git", "rancher/alpine-git", 1)
	image = strings.Replace(image, "prom/", "rancher/prom-", 1)
	image = strings.Replace(image, "quay.io/pires", "rancher", 1)

	Mirrors[image] = orig
	return image
}
