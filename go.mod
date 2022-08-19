module github.com/openfaas/faas-netes

go 1.15

require (
	github.com/google/go-cmp v0.5.5
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/onsi/gomega v1.8.1 // indirect
	github.com/openfaas/faas-provider v0.19.1
	github.com/openfaas/faas/gateway v0.0.0-20210311210633-a6dbb4cd0285
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.2
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v0.18.2
	k8s.io/code-generator v0.18.2
	k8s.io/klog v1.0.0
)
