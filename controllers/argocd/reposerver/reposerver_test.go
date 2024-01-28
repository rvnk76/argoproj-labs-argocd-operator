package reposerver

import (
	"testing"

	argoproj "github.com/argoproj-labs/argocd-operator/api/v1beta1"
	"github.com/argoproj-labs/argocd-operator/common"
	"github.com/argoproj-labs/argocd-operator/pkg/monitoring"
	"github.com/argoproj-labs/argocd-operator/pkg/resource"
	"github.com/argoproj-labs/argocd-operator/pkg/util"
	"github.com/argoproj-labs/argocd-operator/pkg/workloads"
	"github.com/argoproj-labs/argocd-operator/tests/test"
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func makeTestReposerverReconciler(cr *argoproj.ArgoCD, objs ...client.Object) *RepoServerReconciler {
	schemeOpt := func(s *runtime.Scheme) {
		monitoringv1.AddToScheme(s)
		argoproj.AddToScheme(s)
	}
	sch := test.MakeTestReconcilerScheme(schemeOpt)

	client := test.MakeTestReconcilerClient(sch, objs, []client.Object{}, []runtime.Object{})

	return &RepoServerReconciler{
		Client:   client,
		Scheme:   sch,
		Instance: cr,
		Logger:   util.NewLogger(common.RepoServerController),
	}
}

func TestDeleteResources(t *testing.T) {
	tests := []struct {
		name                   string
		resources              []client.Object
		prometheusAPIAvailable bool
		expectedErrors         []string
	}{
		{
			name: "Prometheus API available",
			resources: []client.Object{
				test.MakeTestServiceAccount(
					func(sa *corev1.ServiceAccount) {
						sa.Name = resourceName
					},
				),
				test.MakeTestService(
					func(s *corev1.Service) {
						s.Name = resourceName
					},
				),
				test.MakeTestServiceMonitor(
					func(sm *monitoringv1.ServiceMonitor) {
						sm.Name = resourceMetricsName
					},
				),
				test.MakeTestSecret(
					func(sec *corev1.Secret) {
						sec.Name = common.ArgoCDRepoServerTLSSecretName
					},
				),
				test.MakeTestDeployment(
					func(d *appsv1.Deployment) {
						d.Name = resourceName
					},
				),
			},
			prometheusAPIAvailable: true,
			expectedErrors:         nil,
		},
		{
			name: "Prometheus API not available",
			resources: []client.Object{
				test.MakeTestServiceAccount(
					func(sa *corev1.ServiceAccount) {
						sa.Name = resourceName
					},
				),
				test.MakeTestService(
					func(s *corev1.Service) {
						s.Name = resourceName
					},
				),
				test.MakeTestSecret(
					func(sec *corev1.Secret) {
						sec.Name = common.ArgoCDRepoServerTLSSecretName
					},
				),
				test.MakeTestDeployment(
					func(d *appsv1.Deployment) {
						d.Name = resourceName
					},
				),
			},
			prometheusAPIAvailable: false,
			expectedErrors:         nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reconciler := makeTestReposerverReconciler(
				test.MakeTestArgoCD(),
				tt.resources...,
			)

			reconciler.varSetter()

			monitoring.SetPrometheusAPIFound(tt.prometheusAPIAvailable)
			defer monitoring.SetPrometheusAPIFound(false)

			err := reconciler.DeleteResources()

			if len(tt.expectedErrors) > 0 {
				assert.Error(t, err, "Expected an error but got none.")
				for _, expectedError := range tt.expectedErrors {
					assert.Contains(t, err.Error(), expectedError, "Error message did not contain the expected substring.")
				}
			} else {
				assert.NoError(t, err, "Expected no error but got one.")
			}

			for _, obj := range tt.resources {
				_, err := resource.GetObject(resourceName, test.TestNamespace, obj, reconciler.Client)
				assert.True(t, apierrors.IsNotFound(err))
			}
		})
	}
}

func TestTriggerRollout(t *testing.T) {
	tests := []struct {
		name             string
		reconciler       *RepoServerReconciler
		deploymentexists bool
		expectedError    bool
	}{
		{
			name: "Deployment exists",
			reconciler: makeTestReposerverReconciler(
				test.MakeTestArgoCD(),
				test.MakeTestDeployment(
					func(d *appsv1.Deployment) {
						d.Name = "test-argocd-repo-server"
					},
				),
			),
			deploymentexists: true,
			expectedError:    false,
		},
		{
			name: "Deployment does not exist",
			reconciler: makeTestReposerverReconciler(
				test.MakeTestArgoCD(),
			),
			deploymentexists: false,
			expectedError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.reconciler.varSetter()

			err := tt.reconciler.TriggerRollout(test.TestKey)

			if tt.expectedError {
				assert.Error(t, err, "Expected an error but got none.")
			} else {
				assert.NoError(t, err, "Expected no error but got one.")
			}

			if !tt.expectedError {
				dep, err := workloads.GetDeployment(resourceName, test.TestNamespace, tt.reconciler.Client)
				assert.NoError(t, err)

				_, ok := dep.Spec.Template.ObjectMeta.Labels[test.TestKey]
				assert.True(t, ok)
			}

		})
	}
}
