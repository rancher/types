module github.com/rancher/types

go 1.12

replace (
	github.com/knative/pkg => github.com/rancher/pkg v0.0.0-20190514055449-b30ab9de040e
	github.com/matryer/moq => github.com/rancher/moq v0.0.0-20190404221404-ee5226d43009
	k8s.io/client-go => k8s.io/client-go v0.17.2
)

require (
	github.com/coreos/prometheus-operator v0.33.0
	github.com/knative/pkg v0.0.0-20190817231834-12ee58e32cc8
	github.com/pkg/errors v0.8.1
	github.com/rancher/norman v0.0.0-20200211155126-fc45a55d4dfd
	github.com/rancher/wrangler v0.4.2-0.20200214231136-099089b8a398
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/tools v0.0.0-20191029041327-9cc4af7d6b2c // indirect
	k8s.io/api v0.17.2
	k8s.io/apiextensions-apiserver v0.17.2
	k8s.io/apimachinery v0.17.2
	k8s.io/apiserver v0.17.2
	k8s.io/client-go v2.0.0-alpha.0.0.20181121191925-a47917edff34+incompatible
	k8s.io/gengo v0.0.0-20191120174120-e74f70b9b27e
	k8s.io/kube-aggregator v0.17.2
)
