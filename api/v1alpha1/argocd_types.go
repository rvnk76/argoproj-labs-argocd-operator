/*
Copyright 2019, 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	routev1 "github.com/openshift/api/route/v1"

	"github.com/argoproj-labs/argocd-operator/common"

	autoscaling "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	SchemeBuilder.Register(&ArgoCD{}, &ArgoCDList{})
}

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// Important: Run "make" to regenerate code after modifying this file

//+kubebuilder:object:root=true

// ArgoCD is the Schema for the argocds API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
//+operator-sdk:csv:customresourcedefinitions:resources={{ArgoCD,v1alpha1,""}}
//+operator-sdk:csv:customresourcedefinitions:resources={{ArgoCDExport,v1alpha1,""}}
//+operator-sdk:csv:customresourcedefinitions:resources={{ConfigMap,v1,""}}
//+operator-sdk:csv:customresourcedefinitions:resources={{CronJob,v1beta1,""}}
//+operator-sdk:csv:customresourcedefinitions:resources={{Deployment,v1,""}}
//+operator-sdk:csv:customresourcedefinitions:resources={{Ingress,v1beta1,""}}
//+operator-sdk:csv:customresourcedefinitions:resources={{Job,v1,""}}
//+operator-sdk:csv:customresourcedefinitions:resources={{PersistentVolumeClaim,v1,""}}
//+operator-sdk:csv:customresourcedefinitions:resources={{Pod,v1,""}}
//+operator-sdk:csv:customresourcedefinitions:resources={{Prometheus,v1,""}}
//+operator-sdk:csv:customresourcedefinitions:resources={{ReplicaSet,v1,""}}
//+operator-sdk:csv:customresourcedefinitions:resources={{Route,v1,""}}
//+operator-sdk:csv:customresourcedefinitions:resources={{Secret,v1,""}}
//+operator-sdk:csv:customresourcedefinitions:resources={{Service,v1,""}}
//+operator-sdk:csv:customresourcedefinitions:resources={{ServiceMonitor,v1,""}}
//+operator-sdk:csv:customresourcedefinitions:resources={{StatefulSet,v1,""}}
type ArgoCD struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArgoCDSpec   `json:"spec,omitempty"`
	Status ArgoCDStatus `json:"status,omitempty"`
}

// ArgoCDApplicationControllerProcessorsSpec defines the options for the ArgoCD Application Controller processors.
type ArgoCDApplicationControllerProcessorsSpec struct {
	// Operation is the number of application operation processors.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Operation Processor Count'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Controller","urn:alm:descriptor:com.tectonic.ui:number"}
	Operation int32 `json:"operation,omitempty"`

	// Status is the number of application status processors.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Status Processor Count'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Controller","urn:alm:descriptor:com.tectonic.ui:number"}
	Status int32 `json:"status,omitempty"`
}

// ArgoCDApplicationControllerSpec defines the options for the ArgoCD Application Controller component.
type ArgoCDApplicationControllerSpec struct {
	// Processors contains the options for the Application Controller processors.
	Processors ArgoCDApplicationControllerProcessorsSpec `json:"processors,omitempty"`

	// LogLevel refers to the log level used by the Application Controller component. Defaults to ArgoCDDefaultLogLevel if not configured. Valid options are debug, info, error, and warn.
	LogLevel string `json:"logLevel,omitempty"`

	// Resources defines the Compute Resources required by the container for the Application Controller.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Resource Requirements'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Controller","urn:alm:descriptor:com.tectonic.ui:resourceRequirements"}
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// AppSync is used to control the sync frequency, by default the ArgoCD
	// controller polls Git every 3m by default.
	//
	// Set this to a duration, e.g. 10m or 600s to control the synchronisation
	// frequency.
	// +optional
	AppSync *metav1.Duration `json:"appSync,omitempty"`

	// Sharding contains the options for the Application Controller sharding configuration.
	Sharding ArgoCDApplicationControllerShardSpec `json:"sharding,omitempty"`
}

// ArgoCDApplicationControllerShardSpec defines the options available for enabling sharding for the Application Controller component.
type ArgoCDApplicationControllerShardSpec struct {

	// Enabled defines whether sharding should be enabled on the Application Controller component.
	Enabled bool `json:"enabled,omitempty"`

	// Replicas defines the number of replicas to run in the Application controller shard.
	Replicas int32 `json:"replicas,omitempty"`
}

// ArgoCDApplicationSet defines whether the Argo CD ApplicationSet controller should be installed.
type ArgoCDApplicationSet struct {

	// Image is the Argo CD ApplicationSet image (optional)
	Image string `json:"image,omitempty"`

	// Version is the Argo CD ApplicationSet image tag. (optional)
	Version string `json:"version,omitempty"`

	// Resources defines the Compute Resources required by the container for ApplicationSet.
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// LogLevel describes the log level that should be used by the ApplicationSet controller. Defaults to ArgoCDDefaultLogLevel if not set.  Valid options are debug,info, error, and warn.
	LogLevel string `json:"logLevel,omitempty"`
}

// ArgoCDCASpec defines the CA options for ArgCD.
type ArgoCDCASpec struct {
	// ConfigMapName is the name of the ConfigMap containing the CA Certificate.
	ConfigMapName string `json:"configMapName,omitempty"`

	// SecretName is the name of the Secret containing the CA Certificate and Key.
	SecretName string `json:"secretName,omitempty"`
}

// ArgoCDCertificateSpec defines the options for the ArgoCD certificates.
type ArgoCDCertificateSpec struct {
	// SecretName is the name of the Secret containing the Certificate and Key.
	SecretName string `json:"secretName"`
}

// ArgoCDDexSpec defines the desired state for the Dex server component.
type ArgoCDDexSpec struct {
	//Config is the dex connector configuration.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Configuration",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Dex","urn:alm:descriptor:com.tectonic.ui:text"}
	Config string `json:"config,omitempty"`

	// Image is the Dex container image.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Image",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Dex","urn:alm:descriptor:com.tectonic.ui:text"}
	Image string `json:"image,omitempty"`

	// OpenShiftOAuth enables OpenShift OAuth authentication for the Dex server.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="OpenShift OAuth Enabled'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Dex","urn:alm:descriptor:com.tectonic.ui:booleanSwitch"}
	OpenShiftOAuth bool `json:"openShiftOAuth,omitempty"`

	// Resources defines the Compute Resources required by the container for Dex.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Resource Requirements'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Dex","urn:alm:descriptor:com.tectonic.ui:resourceRequirements"}
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// Version is the Dex container image tag.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Version",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Dex","urn:alm:descriptor:com.tectonic.ui:text"}
	Version string `json:"version,omitempty"`
}

// ArgoCDDexOAuthSpec defines the desired state for the Dex OAuth configuration.
type ArgoCDDexOAuthSpec struct {
	// Enabled will toggle OAuth support for the Dex server.
	Enabled bool `json:"enabled"`
}

// ArgoCDGrafanaSpec defines the desired state for the Grafana component.
type ArgoCDGrafanaSpec struct {
	// Enabled will toggle Grafana support globally for ArgoCD.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Grafana","urn:alm:descriptor:com.tectonic.ui:booleanSwitch"}
	Enabled bool `json:"enabled"`

	// Host is the hostname to use for Ingress/Route resources.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Host",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Grafana","urn:alm:descriptor:com.tectonic.ui:text"}
	Host string `json:"host,omitempty"`

	// Image is the Grafana container image.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Image",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Grafana","urn:alm:descriptor:com.tectonic.ui:text"}
	Image string `json:"image,omitempty"`

	// Ingress defines the desired state for an Ingress for the Grafana component.
	Ingress ArgoCDIngressSpec `json:"ingress,omitempty"`

	// Resources defines the Compute Resources required by the container for Grafana.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Resource Requirements'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Grafana","urn:alm:descriptor:com.tectonic.ui:resourceRequirements"}
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// Route defines the desired state for an OpenShift Route for the Grafana component.
	Route ArgoCDRouteSpec `json:"route,omitempty"`

	// Size is the replica count for the Grafana Deployment.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Size",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Grafana","urn:alm:descriptor:com.tectonic.ui:podCount"}
	Size *int32 `json:"size,omitempty"`

	// Version is the Grafana container image tag.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Version",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Grafana","urn:alm:descriptor:com.tectonic.ui:text"}
	Version string `json:"version,omitempty"`
}

// ArgoCDHASpec defines the desired state for High Availability support for Argo CD.
type ArgoCDHASpec struct {
	// Enabled will toggle HA support globally for Argo CD.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:HA","urn:alm:descriptor:com.tectonic.ui:booleanSwitch"}
	Enabled bool `json:"enabled"`

	// RedisProxyImage is the Redis HAProxy container image.
	RedisProxyImage string `json:"redisProxyImage,omitempty"`

	// RedisProxyVersion is the Redis HAProxy container image tag.
	RedisProxyVersion string `json:"redisProxyVersion,omitempty"`

	// Resources defines the Compute Resources required by the container for HA.
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}

// ArgoCDImportSpec defines the desired state for the ArgoCD import/restore process.
type ArgoCDImportSpec struct {
	// Name of an ArgoCDExport from which to import data.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Name",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Import","urn:alm:descriptor:com.tectonic.ui:text"}
	Name string `json:"name"`

	// Namespace for the ArgoCDExport, defaults to the same namespace as the ArgoCD.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Import","urn:alm:descriptor:com.tectonic.ui:text"}
	Namespace *string `json:"namespace,omitempty"`
}

// ArgoCDIngressSpec defines the desired state for the Ingress resources.
type ArgoCDIngressSpec struct {
	// Annotations is the map of annotations to apply to the Ingress.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Enabled will toggle the creation of the Ingress.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Ingress Enabled'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Grafana","urn:alm:descriptor:com.tectonic.ui:fieldGroup:Prometheus","urn:alm:descriptor:com.tectonic.ui:fieldGroup:Server","urn:alm:descriptor:com.tectonic.ui:booleanSwitch"}
	Enabled bool `json:"enabled"`

	// Path used for the Ingress resource.
	Path string `json:"path,omitempty"`

	// TLS configuration. Currently the Ingress only supports a single TLS
	// port, 443. If multiple members of this list specify different hosts, they
	// will be multiplexed on the same port according to the hostname specified
	// through the SNI TLS extension, if the ingress controller fulfilling the
	// ingress supports SNI.
	// +optional
	TLS []extv1beta1.IngressTLS `json:"tls,omitempty"`
}

//+kubebuilder:object:root=true

// ArgoCDList contains a list of ArgoCD
type ArgoCDList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArgoCD `json:"items"`
}

// ArgoCDPrometheusSpec defines the desired state for the Prometheus component.
type ArgoCDPrometheusSpec struct {
	// Enabled will toggle Prometheus support globally for ArgoCD.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enabled",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Prometheus","urn:alm:descriptor:com.tectonic.ui:booleanSwitch"}
	Enabled bool `json:"enabled"`

	// Host is the hostname to use for Ingress/Route resources.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Host",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Prometheus","urn:alm:descriptor:com.tectonic.ui:text"}
	Host string `json:"host,omitempty"`

	// Ingress defines the desired state for an Ingress for the Prometheus component.
	Ingress ArgoCDIngressSpec `json:"ingress,omitempty"`

	// Route defines the desired state for an OpenShift Route for the Prometheus component.
	Route ArgoCDRouteSpec `json:"route,omitempty"`

	// Size is the replica count for the Prometheus StatefulSet.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Size",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Prometheus","urn:alm:descriptor:com.tectonic.ui:podCount"}
	Size *int32 `json:"size,omitempty"`
}

// ArgoCDRBACSpec defines the desired state for the Argo CD RBAC configuration.
type ArgoCDRBACSpec struct {
	// DefaultPolicy is the name of the default role which Argo CD will falls back to, when
	// authorizing API requests (optional). If omitted or empty, users may be still be able to login,
	// but will see no apps, projects, etc...
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Default Policy'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:RBAC","urn:alm:descriptor:com.tectonic.ui:text"}
	DefaultPolicy *string `json:"defaultPolicy,omitempty"`

	// Policy is CSV containing user-defined RBAC policies and role definitions.
	// Policy rules are in the form:
	//   p, subject, resource, action, object, effect
	// Role definitions and bindings are in the form:
	//   g, subject, inherited-subject
	// See https://github.com/argoproj/argo-cd/blob/master/docs/operator-manual/rbac.md for additional information.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Policy",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:RBAC","urn:alm:descriptor:com.tectonic.ui:text"}
	Policy *string `json:"policy,omitempty"`

	// Scopes controls which OIDC scopes to examine during rbac enforcement (in addition to `sub` scope).
	// If omitted, defaults to: '[groups]'.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Scopes",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:RBAC","urn:alm:descriptor:com.tectonic.ui:text"}
	Scopes *string `json:"scopes,omitempty"`
}

// ArgoCDRedisSpec defines the desired state for the Redis server component.
type ArgoCDRedisSpec struct {
	// Image is the Redis container image.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Image",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Redis","urn:alm:descriptor:com.tectonic.ui:text"}
	Image string `json:"image,omitempty"`

	// Resources defines the Compute Resources required by the container for Redis.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Resource Requirements'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Redis","urn:alm:descriptor:com.tectonic.ui:resourceRequirements"}
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// Version is the Redis container image tag.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Version",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Redis","urn:alm:descriptor:com.tectonic.ui:text"}
	Version string `json:"version,omitempty"`
}

// ArgoCDRepoSpec defines the desired state for the Argo CD repo server component.
type ArgoCDRepoSpec struct {

	// LogLevel describes the log level that should be used by the Repo Server. Defaults to ArgoCDDefaultLogLevel if not set.  Valid options are debug, info, error, and warn.
	LogLevel string `json:"logLevel,omitempty"`

	// MountSAToken describes whether you would like to have the Repo server mount the service account token
	MountSAToken bool `json:"mountsatoken,omitempty"`

	// Resources defines the Compute Resources required by the container for Redis.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Resource Requirements'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Repo","urn:alm:descriptor:com.tectonic.ui:resourceRequirements"}
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// ServiceAccount defines the ServiceAccount user that you would like the Repo server to use
	ServiceAccount string `json:"serviceaccount,omitempty"`

	// VerifyTLS defines whether repo server API should be accessed using strict TLS validation
	VerifyTLS bool `json:"verifytls,omitempty"`

	// AutoTLS specifies the method to use for automatic TLS configuration for the repo server
	// The value specified here can currently be:
	// - openshift - Use the OpenShift service CA to request TLS config
	AutoTLS string `json:"autotls,omitempty"`

	// Image is the ArgoCD Repo Server container image.
	Image string `json:"image,omitempty"`

	// Version is the ArgoCD Repo Server container image tag.
	Version string `json:"version,omitempty"`
}

// ArgoCDRouteSpec defines the desired state for an OpenShift Route.
type ArgoCDRouteSpec struct {
	// Annotations is the map of annotations to use for the Route resource.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Labels is the map of labels to use for the Route resource
	Labels map[string]string `json:"labels,omitempty"`

	// Enabled will toggle the creation of the OpenShift Route.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Route Enabled'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Grafana","urn:alm:descriptor:com.tectonic.ui:fieldGroup:Prometheus","urn:alm:descriptor:com.tectonic.ui:fieldGroup:Server","urn:alm:descriptor:com.tectonic.ui:booleanSwitch"}
	Enabled bool `json:"enabled"`

	// Path the router watches for, to route traffic for to the service.
	Path string `json:"path,omitempty"`

	// TLS provides the ability to configure certificates and termination for the Route.
	TLS *routev1.TLSConfig `json:"tls,omitempty"`

	// WildcardPolicy if any for the route. Currently only 'Subdomain' or 'None' is allowed.
	WildcardPolicy *routev1.WildcardPolicyType `json:"wildcardPolicy,omitempty"`
}

// ArgoCDServerAutoscaleSpec defines the desired state for autoscaling the Argo CD Server component.
type ArgoCDServerAutoscaleSpec struct {
	// Enabled will toggle autoscaling support for the Argo CD Server component.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Autoscale Enabled'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Server","urn:alm:descriptor:com.tectonic.ui:booleanSwitch"}
	Enabled bool `json:"enabled"`

	// HPA defines the HorizontalPodAutoscaler options for the Argo CD Server component.
	HPA *autoscaling.HorizontalPodAutoscalerSpec `json:"hpa,omitempty"`
}

// ArgoCDServerGRPCSpec defines the desired state for the Argo CD Server GRPC options.
type ArgoCDServerGRPCSpec struct {
	// Host is the hostname to use for Ingress/Route resources.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="GRPC Host",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Server","urn:alm:descriptor:com.tectonic.ui:text"}
	Host string `json:"host,omitempty"`

	// Ingress defines the desired state for the Argo CD Server GRPC Ingress.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="GRPC Ingress Enabled'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Server","urn:alm:descriptor:com.tectonic.ui:booleanSwitch"}
	Ingress ArgoCDIngressSpec `json:"ingress,omitempty"`
}

// ArgoCDServerSpec defines the options for the ArgoCD Server component.
type ArgoCDServerSpec struct {
	// Autoscale defines the autoscale options for the Argo CD Server component.
	Autoscale ArgoCDServerAutoscaleSpec `json:"autoscale,omitempty"`

	// GRPC defines the state for the Argo CD Server GRPC options.
	GRPC ArgoCDServerGRPCSpec `json:"grpc,omitempty"`

	// Host is the hostname to use for Ingress/Route resources.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Host",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Server","urn:alm:descriptor:com.tectonic.ui:text"}
	Host string `json:"host,omitempty"`

	// Ingress defines the desired state for an Ingress for the Argo CD Server component.
	Ingress ArgoCDIngressSpec `json:"ingress,omitempty"`

	// Insecure toggles the insecure flag.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Insecure",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Server","urn:alm:descriptor:com.tectonic.ui:booleanSwitch"}
	Insecure bool `json:"insecure,omitempty"`

	// LogLevel refers to the log level to be used by the ArgoCD Server component. Defaults to ArgoCDDefaultLogLevel if not set.  Valid options are debug, info, error, and warn.
	LogLevel string `json:"logLevel,omitempty"`

	// Resources defines the Compute Resources required by the container for the Argo CD server component.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Resource Requirements'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Server","urn:alm:descriptor:com.tectonic.ui:resourceRequirements"}
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`

	// Route defines the desired state for an OpenShift Route for the Argo CD Server component.
	Route ArgoCDRouteSpec `json:"route,omitempty"`

	// Service defines the options for the Service backing the ArgoCD Server component.
	Service ArgoCDServerServiceSpec `json:"service,omitempty"`
}

// ArgoCDServerServiceSpec defines the Service options for Argo CD Server component.
type ArgoCDServerServiceSpec struct {
	// Type is the ServiceType to use for the Service resource.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Service Type'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:Server","urn:alm:descriptor:com.tectonic.ui:text"}
	Type corev1.ServiceType `json:"type"`
}

// SSOProviderType string defines the type of SSO provider.
type SSOProviderType string

const (
	// SSOProviderTypeKeycloak means keycloak will be Installed and Integrated with Argo CD. A new realm with name argocd
	// will be created in this keycloak. This realm will have a client with name argocd that uses OpenShift v4 as Identity Provider.
	SSOProviderTypeKeycloak SSOProviderType = "keycloak"
)

// ArgoCDSSOSpec defines SSO provider.
type ArgoCDSSOSpec struct {
	// Image is the SSO container image.
	Image string `json:"image,omitempty"`
	// Provider installs and configures the given SSO Provider with Argo CD.
	Provider SSOProviderType `json:"provider,omitempty"`
	// Resources defines the Compute Resources required by the container for SSO.
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
	// VerifyTLS set to false disables strict TLS validation.
	VerifyTLS *bool `json:"verifyTLS,omitempty"`
	// Version is the SSO container image tag.
	Version string `json:"version,omitempty"`
}

// KustomizeVersionSpec is used to specify information about a kustomize version to be used within ArgoCD.
type KustomizeVersionSpec struct {
	// Version is a configured kustomize version in the format of vX.Y.Z
	Version string `json:"version,omitempty"`
	// Path is the path to a configured kustomize version on the filesystem of your repo server.
	Path string `json:"path,omitempty"`
}

// ArgoCDSpec defines the desired state of ArgoCD
// +k8s:openapi-gen=true
type ArgoCDSpec struct {

	// ArgoCDApplicationSet defines whether the Argo CD ApplicationSet controller should be installed.
	ApplicationSet *ArgoCDApplicationSet `json:"applicationSet,omitempty"`

	// ApplicationInstanceLabelKey is the key name where Argo CD injects the app name as a tracking label.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Application Instance Label Key'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:advanced"}
	ApplicationInstanceLabelKey string `json:"applicationInstanceLabelKey,omitempty"`

	// ConfigManagementPlugins is used to specify additional config management plugins.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Config Management Plugins'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:advanced"}
	ConfigManagementPlugins string `json:"configManagementPlugins,omitempty"`

	// Controller defines the Application Controller options for ArgoCD.
	Controller ArgoCDApplicationControllerSpec `json:"controller,omitempty"`

	// Dex defines the Dex server options for ArgoCD.
	Dex ArgoCDDexSpec `json:"dex,omitempty"`

	// DisableAdmin will disable the admin user.
	DisableAdmin bool `json:"disableAdmin,omitempty"`

	// GATrackingID is the google analytics tracking ID to use.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Google Analytics Tracking ID'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:advanced"}
	GATrackingID string `json:"gaTrackingID,omitempty"`

	// GAAnonymizeUsers toggles user IDs being hashed before sending to google analytics.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Google Analytics Anonymize Users'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:booleanSwitch","urn:alm:descriptor:com.tectonic.ui:advanced"}
	GAAnonymizeUsers bool `json:"gaAnonymizeUsers,omitempty"`

	// Grafana defines the Grafana server options for ArgoCD.
	Grafana ArgoCDGrafanaSpec `json:"grafana,omitempty"`

	// HA options for High Availability support for the Redis component.
	HA ArgoCDHASpec `json:"ha,omitempty"`

	// HelpChatURL is the URL for getting chat help, this will typically be your Slack channel for support.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Help Chat URL'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:advanced"}
	HelpChatURL string `json:"helpChatURL,omitempty"`

	// HelpChatText is the text for getting chat help, defaults to "Chat now!"
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Help Chat Text'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:advanced"}
	HelpChatText string `json:"helpChatText,omitempty"`

	// Image is the ArgoCD container image for all ArgoCD components.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Image",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:ArgoCD","urn:alm:descriptor:com.tectonic.ui:text"}
	Image string `json:"image,omitempty"`

	// Import is the import/restore options for ArgoCD.
	Import *ArgoCDImportSpec `json:"import,omitempty"`

	// InitialRepositories to configure Argo CD with upon creation of the cluster.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Initial Repositories'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:advanced"}
	InitialRepositories string `json:"initialRepositories,omitempty"`

	// InitialSSHKnownHosts defines the SSH known hosts data upon creation of the cluster for connecting Git repositories via SSH.
	InitialSSHKnownHosts SSHHostsSpec `json:"initialSSHKnownHosts,omitempty"`

	// KustomizeBuildOptions is used to specify build options/parameters to use with `kustomize build`.
	KustomizeBuildOptions string `json:"kustomizeBuildOptions,omitempty"`

	// KustomizeVersions is a listing of configured versions of Kustomize to be made available within ArgoCD.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Kustomize Build Options'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:advanced"}
	KustomizeVersions []KustomizeVersionSpec `json:"kustomizeVersions,omitempty"`

	// OIDCConfig is the OIDC configuration as an alternative to dex.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="OIDC Config'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:advanced"}
	OIDCConfig string `json:"oidcConfig,omitempty"`

	// Prometheus defines the Prometheus server options for ArgoCD.
	Prometheus ArgoCDPrometheusSpec `json:"prometheus,omitempty"`

	// RBAC defines the RBAC configuration for Argo CD.
	RBAC ArgoCDRBACSpec `json:"rbac,omitempty"`

	// Redis defines the Redis server options for ArgoCD.
	Redis ArgoCDRedisSpec `json:"redis,omitempty"`

	// Repo defines the repo server options for Argo CD.
	Repo ArgoCDRepoSpec `json:"repo,omitempty"`

	// RepositoryCredentials are the Git pull credentials to configure Argo CD with upon creation of the cluster.
	RepositoryCredentials string `json:"repositoryCredentials,omitempty"`

	// ResourceCustomizations customizes resource behavior. Keys are in the form: group/Kind.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Resource Customizations'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:advanced"}
	ResourceCustomizations string `json:"resourceCustomizations,omitempty"`

	// ResourceExclusions is used to completely ignore entire classes of resource group/kinds.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Resource Exclusions'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text","urn:alm:descriptor:com.tectonic.ui:advanced"}
	ResourceExclusions string `json:"resourceExclusions,omitempty"`

	// ResourceInclusions is used to only include specific group/kinds in the
	// reconciliation process.
	ResourceInclusions string `json:"resourceInclusions,omitempty"`

	// Server defines the options for the ArgoCD Server component.
	Server ArgoCDServerSpec `json:"server,omitempty"`

	// SSO defines the Single Sign-on configuration for Argo CD
	SSO *ArgoCDSSOSpec `json:"sso,omitempty"`

	// StatusBadgeEnabled toggles application status badge feature.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Status Badge Enabled'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:booleanSwitch","urn:alm:descriptor:com.tectonic.ui:advanced"}
	StatusBadgeEnabled bool `json:"statusBadgeEnabled,omitempty"`

	// TLS defines the TLS options for ArgoCD.
	TLS ArgoCDTLSSpec `json:"tls,omitempty"`

	// UsersAnonymousEnabled toggles anonymous user access.
	// The anonymous users get default role permissions specified argocd-rbac-cm.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Anonymous Users Enabled'",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:booleanSwitch","urn:alm:descriptor:com.tectonic.ui:advanced"}
	UsersAnonymousEnabled bool `json:"usersAnonymousEnabled,omitempty"`

	// Version is the tag to use with the ArgoCD container image for all ArgoCD components.
	//+operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Version",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldGroup:ArgoCD","urn:alm:descriptor:com.tectonic.ui:text"}
	Version string `json:"version,omitempty"`
}

// ArgoCDStatus defines the observed state of ArgoCD
// +k8s:openapi-gen=true
type ArgoCDStatus struct {
	// ApplicationController is a simple, high-level summary of where the Argo CD application controller component is in its lifecycle.
	// There are five possible ApplicationController values:
	// Pending: The Argo CD application controller component has been accepted by the Kubernetes system, but one or more of the required resources have not been created.
	// Running: All of the required Pods for the Argo CD application controller component are in a Ready state.
	// Failed: At least one of the  Argo CD application controller component Pods had a failure.
	// Unknown: For some reason the state of the Argo CD application controller component could not be obtained.
	//+operator-sdk:csv:customresourcedefinitions:type=status,displayName="ApplicationController",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text"}
	ApplicationController string `json:"applicationController,omitempty"`

	// Dex is a simple, high-level summary of where the Argo CD Dex component is in its lifecycle.
	// There are five possible dex values:
	// Pending: The Argo CD Dex component has been accepted by the Kubernetes system, but one or more of the required resources have not been created.
	// Running: All of the required Pods for the Argo CD Dex component are in a Ready state.
	// Failed: At least one of the  Argo CD Dex component Pods had a failure.
	// Unknown: For some reason the state of the Argo CD Dex component could not be obtained.
	//+operator-sdk:csv:customresourcedefinitions:type=status,displayName="Dex",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text"}
	Dex string `json:"dex,omitempty"`

	// Phase is a simple, high-level summary of where the ArgoCD is in its lifecycle.
	// There are five possible phase values:
	// Pending: The ArgoCD has been accepted by the Kubernetes system, but one or more of the required resources have not been created.
	// Available: All of the resources for the ArgoCD are ready.
	// Failed: At least one resource has experienced a failure.
	// Unknown: For some reason the state of the ArgoCD phase could not be obtained.
	//+operator-sdk:csv:customresourcedefinitions:type=status,displayName="Phase",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text"}
	Phase string `json:"phase,omitempty"`

	// Redis is a simple, high-level summary of where the Argo CD Redis component is in its lifecycle.
	// There are five possible redis values:
	// Pending: The Argo CD Redis component has been accepted by the Kubernetes system, but one or more of the required resources have not been created.
	// Running: All of the required Pods for the Argo CD Redis component are in a Ready state.
	// Failed: At least one of the  Argo CD Redis component Pods had a failure.
	// Unknown: For some reason the state of the Argo CD Redis component could not be obtained.
	//+operator-sdk:csv:customresourcedefinitions:type=status,displayName="Redis",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text"}
	Redis string `json:"redis,omitempty"`

	// Repo is a simple, high-level summary of where the Argo CD Repo component is in its lifecycle.
	// There are five possible repo values:
	// Pending: The Argo CD Repo component has been accepted by the Kubernetes system, but one or more of the required resources have not been created.
	// Running: All of the required Pods for the Argo CD Repo component are in a Ready state.
	// Failed: At least one of the  Argo CD Repo component Pods had a failure.
	// Unknown: For some reason the state of the Argo CD Repo component could not be obtained.
	//+operator-sdk:csv:customresourcedefinitions:type=status,displayName="Repo",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text"}
	Repo string `json:"repo,omitempty"`

	// Server is a simple, high-level summary of where the Argo CD server component is in its lifecycle.
	// There are five possible server values:
	// Pending: The Argo CD server component has been accepted by the Kubernetes system, but one or more of the required resources have not been created.
	// Running: All of the required Pods for the Argo CD server component are in a Ready state.
	// Failed: At least one of the  Argo CD server component Pods had a failure.
	// Unknown: For some reason the state of the Argo CD server component could not be obtained.
	//+operator-sdk:csv:customresourcedefinitions:type=status,displayName="Server",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:text"}
	Server string `json:"server,omitempty"`

	// RepoTLSChecksum contains the SHA256 checksum of the latest known state of tls.crt and tls.key in the argocd-repo-server-tls secret.
	RepoTLSChecksum string `json:"repoTLSChecksum,omitempty"`
}

// ArgoCDTLSSpec defines the TLS options for ArgCD.
type ArgoCDTLSSpec struct {
	// CA defines the CA options.
	CA ArgoCDCASpec `json:"ca,omitempty"`

	// InitialCerts defines custom TLS certificates upon creation of the cluster for connecting Git repositories via HTTPS.
	InitialCerts map[string]string `json:"initialCerts,omitempty"`
}

type SSHHostsSpec struct {
	// ExcludeDefaultHosts describes whether you would like to include the default
	// list of SSH Known Hosts provided by ArgoCD.
	ExcludeDefaultHosts bool `json:"excludedefaulthosts,omitempty"`

	// Keys describes a custom set of SSH Known Hosts that you would like to
	// have included in your ArgoCD server.
	Keys string `json:"keys,omitempty"`
}

// IsDeletionFinalizerPresent checks if the instance has deletion finalizer
func (argocd *ArgoCD) IsDeletionFinalizerPresent() bool {
	for _, finalizer := range argocd.GetFinalizers() {
		if finalizer == common.ArgoCDDeletionFinalizer {
			return true
		}
	}
	return false
}

// ApplicationInstanceLabelKey returns either the custom application instance
// label key if set, or the default value.
func (a *ArgoCD) ApplicationInstanceLabelKey() string {
	if a.Spec.ApplicationInstanceLabelKey != "" {
		return a.Spec.ApplicationInstanceLabelKey
	} else {
		return common.ArgoCDDefaultApplicationInstanceLabelKey
	}
}
