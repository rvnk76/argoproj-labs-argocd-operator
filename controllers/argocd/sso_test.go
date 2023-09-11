// Copyright 2021 ArgoCD Operator Developers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package argocd

import (
	"context"
	"errors"
	"testing"

	oappsv1 "github.com/openshift/api/apps/v1"
	routev1 "github.com/openshift/api/route/v1"
	templatev1 "github.com/openshift/api/template/v1"
	"github.com/stretchr/testify/assert"
	k8sappsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/argoproj-labs/argocd-operator/api/v1alpha1"
	argov1alpha1 "github.com/argoproj-labs/argocd-operator/api/v1alpha1"
	"github.com/argoproj-labs/argocd-operator/pkg/workloads"
)

func makeFakeReconciler(t *testing.T, acd *argov1alpha1.ArgoCD, objs ...client.Object) *ArgoCDReconciler {
	t.Helper()
	s := scheme.Scheme
	// Register template scheme
	s.AddKnownTypes(templatev1.SchemeGroupVersion)
	s.AddKnownTypes(oappsv1.SchemeGroupVersion)
	assert.NoError(t, argov1alpha1.AddToScheme(s))
	templatev1.Install(s)
	oappsv1.Install(s)
	routev1.Install(s)

	cl := fake.NewClientBuilder().WithObjects(objs...).Build()
	return &ArgoCDReconciler{
		Client: cl,
		Scheme: s,
	}
}

func TestReconcile_testKeycloakTemplateInstance(t *testing.T) {
	logf.SetLogger(ZapLogger(true))
	a := makeTestArgoCDForKeycloak()

	workloads.SetTemplateAPIFound(true)
	r := makeFakeReconciler(t, a)
	assert.NoError(t, createNamespace(r, a.Namespace, ""))

	assert.NoError(t, r.reconcileSSO(a))

	templateInstance := &templatev1.TemplateInstance{}
	assert.NoError(t, r.Client.Get(
		context.TODO(),
		types.NamespacedName{
			Name:      "rhsso",
			Namespace: a.Namespace,
		},
		templateInstance))
}

func TestReconcile_noTemplateInstance(t *testing.T) {
	logf.SetLogger(ZapLogger(true))
	a := makeTestArgoCDForKeycloak()
	r := makeFakeReconciler(t, a)
	assert.NoError(t, createNamespace(r, a.Namespace, ""))

	assert.NoError(t, r.reconcileSSO(a))
}

func TestReconcile_illegalSSOConfiguration(t *testing.T) {
	logf.SetLogger(ZapLogger(true))

	tests := []struct {
		name                     string
		argoCD                   *argov1alpha1.ArgoCD
		wantErr                  bool
		Err                      error
		wantSSOConfigLegalStatus string
	}{
		{
			name:                     "no conflicts - no sso configured",
			argoCD:                   makeTestArgoCD(func(ac *argov1alpha1.ArgoCD) {}),
			wantErr:                  false,
			wantSSOConfigLegalStatus: "Unknown",
		},
		{
			name: "no conflict - case insensitive sso provider value",
			argoCD: makeTestArgoCD(func(ac *argov1alpha1.ArgoCD) {
				ac.Spec.SSO = &v1alpha1.ArgoCDSSOSpec{
					Provider: "DEX",
					Dex: &v1alpha1.ArgoCDDexSpec{
						Config: "test-config",
					},
				}
			}),
			wantErr:                  false,
			wantSSOConfigLegalStatus: "Success",
		},
		{
			name: "no conflict - valid dex sso configurations",
			argoCD: makeTestArgoCD(func(ac *argov1alpha1.ArgoCD) {
				ac.Spec.SSO = &v1alpha1.ArgoCDSSOSpec{
					Provider: "dex",
					Dex: &v1alpha1.ArgoCDDexSpec{
						Config:         "test-config",
						OpenShiftOAuth: false,
					},
				}
			}),
			wantErr:                  false,
			wantSSOConfigLegalStatus: "Success",
		},
		{
			name: "no conflict - valid keycloak sso configurations",
			argoCD: makeTestArgoCD(func(ac *argov1alpha1.ArgoCD) {
				ac.Spec.SSO = &v1alpha1.ArgoCDSSOSpec{
					Provider: "keycloak",
				}
			}),
			wantErr:                  false,
			wantSSOConfigLegalStatus: "Success",
		},
		{
			name: "sso provider dex but no .spec.sso.dex provided",
			argoCD: makeTestArgoCD(func(ac *argov1alpha1.ArgoCD) {
				ac.Spec.SSO = &v1alpha1.ArgoCDSSOSpec{
					Provider: v1alpha1.SSOProviderTypeDex,
				}
			}),
			wantErr:                  true,
			Err:                      errors.New("illegal SSO configuration: must supply valid dex configuration when requested SSO provider is dex"),
			wantSSOConfigLegalStatus: "Failed",
		},
		{
			name: "sso provider dex + `.spec.sso.keycloak`",
			argoCD: makeTestArgoCD(func(ac *argov1alpha1.ArgoCD) {
				ac.Spec.SSO = &v1alpha1.ArgoCDSSOSpec{
					Keycloak: &v1alpha1.ArgoCDKeycloakSpec{
						Image:   "test-image",
						Version: "test-image-version",
					},
					Provider: argov1alpha1.SSOProviderTypeDex,
					Dex: &v1alpha1.ArgoCDDexSpec{
						Config: "test",
					},
				}
			}),
			wantErr:                  true,
			Err:                      errors.New("illegal SSO configuration: cannot supply keycloak configuration in .spec.sso.keycloak when requested SSO provider is dex"),
			wantSSOConfigLegalStatus: "Failed",
		},
		{
			name: "sso provider keycloak + `.spec.sso.dex`",
			argoCD: makeTestArgoCD(func(ac *argov1alpha1.ArgoCD) {
				ac.Spec.SSO = &v1alpha1.ArgoCDSSOSpec{
					Provider: argov1alpha1.SSOProviderTypeKeycloak,
					Dex: &v1alpha1.ArgoCDDexSpec{
						Config:         "test-config",
						OpenShiftOAuth: true,
					},
				}
			}),
			wantErr:                  true,
			Err:                      errors.New("illegal SSO configuration: cannot supply dex configuration when requested SSO provider is keycloak"),
			wantSSOConfigLegalStatus: "Failed",
		},
		{
			name: "sso provider missing but sso.dex/keycloak supplied",
			argoCD: makeTestArgoCD(func(ac *argov1alpha1.ArgoCD) {
				ac.Spec.SSO = &v1alpha1.ArgoCDSSOSpec{
					Dex: &v1alpha1.ArgoCDDexSpec{
						Config:         "test-config",
						OpenShiftOAuth: true,
					},
					Keycloak: &v1alpha1.ArgoCDKeycloakSpec{
						Image: "test-image",
					},
				}
			}),
			wantErr:                  true,
			Err:                      errors.New("illegal SSO configuration: Cannot specify SSO provider spec without specifying SSO provider type"),
			wantSSOConfigLegalStatus: "Failed",
		},
		{
			name: "unsupported sso provider but sso.dex/keycloak supplied",
			argoCD: makeTestArgoCD(func(ac *argov1alpha1.ArgoCD) {
				ac.Spec.SSO = &v1alpha1.ArgoCDSSOSpec{
					Provider: "Unsupported",
					Dex: &v1alpha1.ArgoCDDexSpec{
						Config:         "test-config",
						OpenShiftOAuth: true,
					},
				}
			}),
			wantErr:                  true,
			Err:                      errors.New("illegal SSO configuration: Unsupported SSO provider type. Supported providers are dex and keycloak"),
			wantSSOConfigLegalStatus: "Failed",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := makeTestReconciler(t, test.argoCD)
			assert.NoError(t, createNamespace(r, test.argoCD.Namespace, ""))

			err := r.reconcileSSO(test.argoCD)
			assert.Equal(t, test.wantSSOConfigLegalStatus, ssoConfigLegalStatus)
			if err != nil {
				if !test.wantErr {
					// ignore unexpected errors for legal sso configurations.
					// keycloak reconciliation code expects a live cluster &
					// therefore throws unexpected errors during unit testing
					if ssoConfigLegalStatus != ssoLegalSuccess {
						t.Errorf("Got unexpected error")
					}
				} else {
					assert.Equal(t, test.Err, err)
				}
			} else {
				if test.wantErr {
					t.Errorf("expected error but didn't get one")
				}
			}
		})
	}

}

func TestReconcile_testKeycloakK8sInstance(t *testing.T) {
	logf.SetLogger(ZapLogger(true))
	a := makeTestArgoCDForKeycloak()

	// Cluster does not have a template instance
	workloads.SetTemplateAPIFound(false)
	r := makeReconciler(t, a)
	assert.NoError(t, createNamespace(r, a.Namespace, ""))

	assert.NoError(t, r.reconcileSSO(a))
}

func TestReconcile_testKeycloakInstanceResources(t *testing.T) {
	logf.SetLogger(ZapLogger(true))
	a := makeTestArgoCDForKeycloak()

	// Cluster does not have a template instance
	workloads.SetTemplateAPIFound(false)
	r := makeReconciler(t, a)
	assert.NoError(t, createNamespace(r, a.Namespace, ""))

	assert.NoError(t, r.reconcileSSO(a))

	// Keycloak Deployment
	deployment := &k8sappsv1.Deployment{}
	err := r.Client.Get(context.TODO(), types.NamespacedName{Name: defaultKeycloakIdentifier, Namespace: a.Namespace}, deployment)
	assert.NoError(t, err)

	assert.Equal(t, deployment.Name, defaultKeycloakIdentifier)
	assert.Equal(t, deployment.Namespace, a.Namespace)

	testLabels := map[string]string{
		"app": defaultKeycloakIdentifier,
	}
	assert.Equal(t, deployment.Labels, testLabels)

	testSelector := &v1.LabelSelector{
		MatchLabels: map[string]string{
			"app": defaultKeycloakIdentifier,
		},
	}
	assert.Equal(t, deployment.Spec.Selector, testSelector)

	assert.Equal(t, deployment.Spec.Template.ObjectMeta.Labels, testLabels)
	assert.Equal(t, deployment.Spec.Template.Spec.Containers[0].Name,
		defaultKeycloakIdentifier)
	assert.Equal(t, deployment.Spec.Template.Spec.Containers[0].Image,
		getKeycloakContainerImage(a))

	testEnv := []corev1.EnvVar{
		{Name: "KEYCLOAK_USER", Value: defaultKeycloakAdminUser},
		{Name: "KEYCLOAK_PASSWORD", Value: defaultKeycloakAdminPassword},
		{Name: "PROXY_ADDRESS_FORWARDING", Value: "true"},
	}
	assert.Equal(t, deployment.Spec.Template.Spec.Containers[0].Env,
		testEnv)

	// Keycloak Service
	svc := &corev1.Service{}
	err = r.Client.Get(context.TODO(), types.NamespacedName{Name: defaultKeycloakIdentifier, Namespace: a.Namespace}, svc)
	assert.NoError(t, err)

	assert.Equal(t, svc.Name, defaultKeycloakIdentifier)
	assert.Equal(t, svc.Namespace, a.Namespace)
	assert.Equal(t, svc.Labels, testLabels)

	assert.Equal(t, svc.Spec.Selector, testLabels)
	assert.Equal(t, svc.Spec.Type, corev1.ServiceType("LoadBalancer"))

	// Keycloak Ingress
	ing := &networkingv1.Ingress{}
	testPathType := networkingv1.PathTypeImplementationSpecific
	err = r.Client.Get(context.TODO(), types.NamespacedName{Name: defaultKeycloakIdentifier, Namespace: a.Namespace}, ing)
	assert.NoError(t, err)

	assert.Equal(t, ing.Name, defaultKeycloakIdentifier)
	assert.Equal(t, ing.Namespace, a.Namespace)

	testTLS := []networkingv1.IngressTLS{
		{
			Hosts: []string{keycloakIngressHost},
		},
	}
	assert.Equal(t, ing.Spec.TLS, testTLS)

	testRules := []networkingv1.IngressRule{
		{
			Host: keycloakIngressHost,
			IngressRuleValue: networkingv1.IngressRuleValue{
				HTTP: &networkingv1.HTTPIngressRuleValue{
					Paths: []networkingv1.HTTPIngressPath{
						{
							Path: "/",
							Backend: networkingv1.IngressBackend{
								Service: &networkingv1.IngressServiceBackend{
									Name: defaultKeycloakIdentifier,
									Port: networkingv1.ServiceBackendPort{
										Name: "http",
									},
								},
							},
							PathType: &testPathType,
						},
					},
				},
			},
		},
	}

	assert.Equal(t, ing.Spec.Rules, testRules)
}
