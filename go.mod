module github.com/rancher/types

go 1.13

replace (
	github.com/knative/pkg => github.com/rancher/pkg v0.0.0-20190514055449-b30ab9de040e
	github.com/matryer/moq => github.com/rancher/moq v0.0.0-20190404221404-ee5226d43009
	k8s.io/client-go => k8s.io/client-go v0.18.0
)

require (
	github.com/coreos/prometheus-operator v0.36.0
	github.com/knative/pkg v0.0.0-20190817231834-12ee58e32cc8
	github.com/pkg/errors v0.9.1
	github.com/rancher/lasso v0.0.0-20200515155337-a34e1e26ad91
	github.com/rancher/norman v0.0.0-20200609220258-00d350370ee8
	github.com/rancher/wrangler v0.6.2-0.20200515155908-1923f3f8ec3f
	github.com/rancher/wrangler-api v0.6.1-0.20200515193802-dcf70881b087
	github.com/sirupsen/logrus v1.4.2
	k8s.io/api v0.18.0
	k8s.io/apiextensions-apiserver v0.18.0
	k8s.io/apimachinery v0.18.0
	k8s.io/apiserver v0.18.0
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/gengo v0.0.0-20200114144118-36b2048a9120
	k8s.io/kube-aggregator v0.18.0
)
