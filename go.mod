module github.com/argoproj-labs/argocd-operator

go 1.14

require (
	github.com/argoproj/argo-cd v1.5.8
	github.com/coreos/prometheus-operator v0.40.0
	github.com/go-openapi/spec v0.19.7
	github.com/google/go-cmp v0.4.0
	github.com/json-iterator/go v1.1.10
	github.com/keycloak/keycloak-operator v0.0.0-20210824124316-64b497530099
	github.com/openshift/api v3.9.1-0.20190916204813-cdbe64fb0c91+incompatible
	github.com/openshift/client-go v0.0.0-20200325131901-f7baeb993edb
	github.com/operator-framework/operator-sdk v0.18.2
	github.com/pkg/errors v0.9.1
	github.com/sethvargo/go-password v0.2.0
	github.com/spf13/pflag v1.0.5
	gopkg.in/yaml.v2 v2.3.0
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.20.6
	k8s.io/apimachinery v0.20.6
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20201113171705-d219536bb9fd
	sigs.k8s.io/controller-runtime v0.6.0
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible // Required by OLM
	k8s.io/client-go => k8s.io/client-go v0.19.2 // Required by prometheus-operator
)

replace k8s.io/api => k8s.io/api v0.19.2

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.2

replace k8s.io/apimachinery => k8s.io/apimachinery v0.19.5-rc.0

replace k8s.io/apiserver => k8s.io/apiserver v0.19.2

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.19.2

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.19.2

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.19.2

replace k8s.io/code-generator => k8s.io/code-generator v0.19.5-rc.0

replace k8s.io/component-base => k8s.io/component-base v0.19.2

replace k8s.io/controller-manager => k8s.io/controller-manager v0.19.15-rc.0

replace k8s.io/cri-api => k8s.io/cri-api v0.19.5-rc.0

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.19.2

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.19.2

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.19.2

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.19.2

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.19.2

replace k8s.io/kubectl => k8s.io/kubectl v0.19.2

replace k8s.io/kubelet => k8s.io/kubelet v0.19.2

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.19.2

replace k8s.io/metrics => k8s.io/metrics v0.19.2

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.19.2

replace k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.19.2

replace k8s.io/sample-controller => k8s.io/sample-controller v0.19.2
