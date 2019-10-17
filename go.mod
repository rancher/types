module github.com/rancher/types

go 1.12

replace (
	github.com/knative/pkg => github.com/rancher/pkg v0.0.0-20190514055449-b30ab9de040e
	github.com/matryer/moq => github.com/rancher/moq v0.0.0-20190404221404-ee5226d43009
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90
)

require (
	github.com/coreos/prometheus-operator v0.33.0
	github.com/knative/pkg v0.0.0-20190817231834-12ee58e32cc8
	github.com/pkg/errors v0.8.1
	github.com/rancher/norman v0.0.0-20191003174345-0ac7dd6ccb36
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/tools v0.0.0-20190926165942-a8d5d34286bd // indirect
	k8s.io/api v0.0.0-20190918155943-95b840bb6a1f
	k8s.io/apiextensions-apiserver v0.0.0-20190918161926-8f644eb6e783
	k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655
	k8s.io/apiserver v0.0.0-20190918160949-bfa5e2e684ad
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/gengo v0.0.0-20190822140433-26a664648505
	k8s.io/kube-aggregator v0.0.0-20190805183716-8439689952da
)
