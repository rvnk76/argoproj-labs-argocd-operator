//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	routev1 "github.com/openshift/api/route/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	"k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCD) DeepCopyInto(out *ArgoCD) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCD.
func (in *ArgoCD) DeepCopy() *ArgoCD {
	if in == nil {
		return nil
	}
	out := new(ArgoCD)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArgoCD) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDApplicationControllerProcessorsSpec) DeepCopyInto(out *ArgoCDApplicationControllerProcessorsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDApplicationControllerProcessorsSpec.
func (in *ArgoCDApplicationControllerProcessorsSpec) DeepCopy() *ArgoCDApplicationControllerProcessorsSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDApplicationControllerProcessorsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDApplicationControllerShardSpec) DeepCopyInto(out *ArgoCDApplicationControllerShardSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDApplicationControllerShardSpec.
func (in *ArgoCDApplicationControllerShardSpec) DeepCopy() *ArgoCDApplicationControllerShardSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDApplicationControllerShardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDApplicationControllerSpec) DeepCopyInto(out *ArgoCDApplicationControllerSpec) {
	*out = *in
	out.Processors = in.Processors
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.AppSync != nil {
		in, out := &in.AppSync, &out.AppSync
		*out = new(metav1.Duration)
		**out = **in
	}
	out.Sharding = in.Sharding
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDApplicationControllerSpec.
func (in *ArgoCDApplicationControllerSpec) DeepCopy() *ArgoCDApplicationControllerSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDApplicationControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDApplicationSet) DeepCopyInto(out *ArgoCDApplicationSet) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	in.WebhookServer.DeepCopyInto(&out.WebhookServer)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDApplicationSet.
func (in *ArgoCDApplicationSet) DeepCopy() *ArgoCDApplicationSet {
	if in == nil {
		return nil
	}
	out := new(ArgoCDApplicationSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDCASpec) DeepCopyInto(out *ArgoCDCASpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDCASpec.
func (in *ArgoCDCASpec) DeepCopy() *ArgoCDCASpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDCASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDCertificateSpec) DeepCopyInto(out *ArgoCDCertificateSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDCertificateSpec.
func (in *ArgoCDCertificateSpec) DeepCopy() *ArgoCDCertificateSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDCertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDDexOAuthSpec) DeepCopyInto(out *ArgoCDDexOAuthSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDDexOAuthSpec.
func (in *ArgoCDDexOAuthSpec) DeepCopy() *ArgoCDDexOAuthSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDDexOAuthSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDDexSpec) DeepCopyInto(out *ArgoCDDexSpec) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDDexSpec.
func (in *ArgoCDDexSpec) DeepCopy() *ArgoCDDexSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDDexSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDExport) DeepCopyInto(out *ArgoCDExport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDExport.
func (in *ArgoCDExport) DeepCopy() *ArgoCDExport {
	if in == nil {
		return nil
	}
	out := new(ArgoCDExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArgoCDExport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDExportList) DeepCopyInto(out *ArgoCDExportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArgoCDExport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDExportList.
func (in *ArgoCDExportList) DeepCopy() *ArgoCDExportList {
	if in == nil {
		return nil
	}
	out := new(ArgoCDExportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArgoCDExportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDExportSpec) DeepCopyInto(out *ArgoCDExportSpec) {
	*out = *in
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(ArgoCDExportStorageSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDExportSpec.
func (in *ArgoCDExportSpec) DeepCopy() *ArgoCDExportSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDExportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDExportStatus) DeepCopyInto(out *ArgoCDExportStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDExportStatus.
func (in *ArgoCDExportStatus) DeepCopy() *ArgoCDExportStatus {
	if in == nil {
		return nil
	}
	out := new(ArgoCDExportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDExportStorageSpec) DeepCopyInto(out *ArgoCDExportStorageSpec) {
	*out = *in
	if in.PVC != nil {
		in, out := &in.PVC, &out.PVC
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDExportStorageSpec.
func (in *ArgoCDExportStorageSpec) DeepCopy() *ArgoCDExportStorageSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDExportStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDGrafanaSpec) DeepCopyInto(out *ArgoCDGrafanaSpec) {
	*out = *in
	in.Ingress.DeepCopyInto(&out.Ingress)
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	in.Route.DeepCopyInto(&out.Route)
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDGrafanaSpec.
func (in *ArgoCDGrafanaSpec) DeepCopy() *ArgoCDGrafanaSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDGrafanaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDHASpec) DeepCopyInto(out *ArgoCDHASpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDHASpec.
func (in *ArgoCDHASpec) DeepCopy() *ArgoCDHASpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDHASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDImportSpec) DeepCopyInto(out *ArgoCDImportSpec) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDImportSpec.
func (in *ArgoCDImportSpec) DeepCopy() *ArgoCDImportSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDImportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDIngressSpec) DeepCopyInto(out *ArgoCDIngressSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.IngressClassName != nil {
		in, out := &in.IngressClassName, &out.IngressClassName
		*out = new(string)
		**out = **in
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = make([]networkingv1.IngressTLS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDIngressSpec.
func (in *ArgoCDIngressSpec) DeepCopy() *ArgoCDIngressSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDIngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDKeycloakSpec) DeepCopyInto(out *ArgoCDKeycloakSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.VerifyTLS != nil {
		in, out := &in.VerifyTLS, &out.VerifyTLS
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDKeycloakSpec.
func (in *ArgoCDKeycloakSpec) DeepCopy() *ArgoCDKeycloakSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDKeycloakSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDList) DeepCopyInto(out *ArgoCDList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArgoCD, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDList.
func (in *ArgoCDList) DeepCopy() *ArgoCDList {
	if in == nil {
		return nil
	}
	out := new(ArgoCDList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArgoCDList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDNodePlacementSpec) DeepCopyInto(out *ArgoCDNodePlacementSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDNodePlacementSpec.
func (in *ArgoCDNodePlacementSpec) DeepCopy() *ArgoCDNodePlacementSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDNodePlacementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDNotifications) DeepCopyInto(out *ArgoCDNotifications) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDNotifications.
func (in *ArgoCDNotifications) DeepCopy() *ArgoCDNotifications {
	if in == nil {
		return nil
	}
	out := new(ArgoCDNotifications)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDPrometheusSpec) DeepCopyInto(out *ArgoCDPrometheusSpec) {
	*out = *in
	in.Ingress.DeepCopyInto(&out.Ingress)
	in.Route.DeepCopyInto(&out.Route)
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDPrometheusSpec.
func (in *ArgoCDPrometheusSpec) DeepCopy() *ArgoCDPrometheusSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDPrometheusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDRBACSpec) DeepCopyInto(out *ArgoCDRBACSpec) {
	*out = *in
	if in.DefaultPolicy != nil {
		in, out := &in.DefaultPolicy, &out.DefaultPolicy
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = new(string)
		**out = **in
	}
	if in.PolicyMatcherMode != nil {
		in, out := &in.PolicyMatcherMode, &out.PolicyMatcherMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDRBACSpec.
func (in *ArgoCDRBACSpec) DeepCopy() *ArgoCDRBACSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDRBACSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDRedisSpec) DeepCopyInto(out *ArgoCDRedisSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDRedisSpec.
func (in *ArgoCDRedisSpec) DeepCopy() *ArgoCDRedisSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDRedisSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDRepoSpec) DeepCopyInto(out *ArgoCDRepoSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.ExecTimeout != nil {
		in, out := &in.ExecTimeout, &out.ExecTimeout
		*out = new(int)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SidecarContainers != nil {
		in, out := &in.SidecarContainers, &out.SidecarContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDRepoSpec.
func (in *ArgoCDRepoSpec) DeepCopy() *ArgoCDRepoSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDRepoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDRouteSpec) DeepCopyInto(out *ArgoCDRouteSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(routev1.TLSConfig)
		**out = **in
	}
	if in.WildcardPolicy != nil {
		in, out := &in.WildcardPolicy, &out.WildcardPolicy
		*out = new(routev1.WildcardPolicyType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDRouteSpec.
func (in *ArgoCDRouteSpec) DeepCopy() *ArgoCDRouteSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDSSOSpec) DeepCopyInto(out *ArgoCDSSOSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.VerifyTLS != nil {
		in, out := &in.VerifyTLS, &out.VerifyTLS
		*out = new(bool)
		**out = **in
	}
	if in.Dex != nil {
		in, out := &in.Dex, &out.Dex
		*out = new(ArgoCDDexSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Keycloak != nil {
		in, out := &in.Keycloak, &out.Keycloak
		*out = new(ArgoCDKeycloakSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDSSOSpec.
func (in *ArgoCDSSOSpec) DeepCopy() *ArgoCDSSOSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDSSOSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDServerAutoscaleSpec) DeepCopyInto(out *ArgoCDServerAutoscaleSpec) {
	*out = *in
	if in.HPA != nil {
		in, out := &in.HPA, &out.HPA
		*out = new(autoscalingv1.HorizontalPodAutoscalerSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDServerAutoscaleSpec.
func (in *ArgoCDServerAutoscaleSpec) DeepCopy() *ArgoCDServerAutoscaleSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDServerAutoscaleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDServerGRPCSpec) DeepCopyInto(out *ArgoCDServerGRPCSpec) {
	*out = *in
	in.Ingress.DeepCopyInto(&out.Ingress)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDServerGRPCSpec.
func (in *ArgoCDServerGRPCSpec) DeepCopy() *ArgoCDServerGRPCSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDServerGRPCSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDServerServiceSpec) DeepCopyInto(out *ArgoCDServerServiceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDServerServiceSpec.
func (in *ArgoCDServerServiceSpec) DeepCopy() *ArgoCDServerServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDServerServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDServerSpec) DeepCopyInto(out *ArgoCDServerSpec) {
	*out = *in
	in.Autoscale.DeepCopyInto(&out.Autoscale)
	in.GRPC.DeepCopyInto(&out.GRPC)
	in.Ingress.DeepCopyInto(&out.Ingress)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	in.Route.DeepCopyInto(&out.Route)
	out.Service = in.Service
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraCommandArgs != nil {
		in, out := &in.ExtraCommandArgs, &out.ExtraCommandArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDServerSpec.
func (in *ArgoCDServerSpec) DeepCopy() *ArgoCDServerSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDSpec) DeepCopyInto(out *ArgoCDSpec) {
	*out = *in
	if in.ApplicationSet != nil {
		in, out := &in.ApplicationSet, &out.ApplicationSet
		*out = new(ArgoCDApplicationSet)
		(*in).DeepCopyInto(*out)
	}
	in.Controller.DeepCopyInto(&out.Controller)
	if in.Dex != nil {
		in, out := &in.Dex, &out.Dex
		*out = new(ArgoCDDexSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraConfig != nil {
		in, out := &in.ExtraConfig, &out.ExtraConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Grafana.DeepCopyInto(&out.Grafana)
	in.HA.DeepCopyInto(&out.HA)
	if in.Import != nil {
		in, out := &in.Import, &out.Import
		*out = new(ArgoCDImportSpec)
		(*in).DeepCopyInto(*out)
	}
	out.InitialSSHKnownHosts = in.InitialSSHKnownHosts
	if in.KustomizeVersions != nil {
		in, out := &in.KustomizeVersions, &out.KustomizeVersions
		*out = make([]KustomizeVersionSpec, len(*in))
		copy(*out, *in)
	}
	if in.NodePlacement != nil {
		in, out := &in.NodePlacement, &out.NodePlacement
		*out = new(ArgoCDNodePlacementSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Notifications.DeepCopyInto(&out.Notifications)
	in.Prometheus.DeepCopyInto(&out.Prometheus)
	in.RBAC.DeepCopyInto(&out.RBAC)
	in.Redis.DeepCopyInto(&out.Redis)
	in.Repo.DeepCopyInto(&out.Repo)
	if in.ResourceHealthChecks != nil {
		in, out := &in.ResourceHealthChecks, &out.ResourceHealthChecks
		*out = make([]ResourceHealthCheck, len(*in))
		copy(*out, *in)
	}
	if in.ResourceIgnoreDifferences != nil {
		in, out := &in.ResourceIgnoreDifferences, &out.ResourceIgnoreDifferences
		*out = make([]ResourceIgnoreDifference, len(*in))
		copy(*out, *in)
	}
	if in.ResourceActions != nil {
		in, out := &in.ResourceActions, &out.ResourceActions
		*out = make([]ResourceAction, len(*in))
		copy(*out, *in)
	}
	in.Server.DeepCopyInto(&out.Server)
	if in.SSO != nil {
		in, out := &in.SSO, &out.SSO
		*out = new(ArgoCDSSOSpec)
		(*in).DeepCopyInto(*out)
	}
	in.TLS.DeepCopyInto(&out.TLS)
	if in.Banner != nil {
		in, out := &in.Banner, &out.Banner
		*out = new(Banner)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDSpec.
func (in *ArgoCDSpec) DeepCopy() *ArgoCDSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDStatus) DeepCopyInto(out *ArgoCDStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDStatus.
func (in *ArgoCDStatus) DeepCopy() *ArgoCDStatus {
	if in == nil {
		return nil
	}
	out := new(ArgoCDStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDTLSSpec) DeepCopyInto(out *ArgoCDTLSSpec) {
	*out = *in
	out.CA = in.CA
	if in.InitialCerts != nil {
		in, out := &in.InitialCerts, &out.InitialCerts
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDTLSSpec.
func (in *ArgoCDTLSSpec) DeepCopy() *ArgoCDTLSSpec {
	if in == nil {
		return nil
	}
	out := new(ArgoCDTLSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Banner) DeepCopyInto(out *Banner) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Banner.
func (in *Banner) DeepCopy() *Banner {
	if in == nil {
		return nil
	}
	out := new(Banner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KustomizeVersionSpec) DeepCopyInto(out *KustomizeVersionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KustomizeVersionSpec.
func (in *KustomizeVersionSpec) DeepCopy() *KustomizeVersionSpec {
	if in == nil {
		return nil
	}
	out := new(KustomizeVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceAction) DeepCopyInto(out *ResourceAction) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceAction.
func (in *ResourceAction) DeepCopy() *ResourceAction {
	if in == nil {
		return nil
	}
	out := new(ResourceAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceHealthCheck) DeepCopyInto(out *ResourceHealthCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceHealthCheck.
func (in *ResourceHealthCheck) DeepCopy() *ResourceHealthCheck {
	if in == nil {
		return nil
	}
	out := new(ResourceHealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceIgnoreDifference) DeepCopyInto(out *ResourceIgnoreDifference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceIgnoreDifference.
func (in *ResourceIgnoreDifference) DeepCopy() *ResourceIgnoreDifference {
	if in == nil {
		return nil
	}
	out := new(ResourceIgnoreDifference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHHostsSpec) DeepCopyInto(out *SSHHostsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHHostsSpec.
func (in *SSHHostsSpec) DeepCopy() *SSHHostsSpec {
	if in == nil {
		return nil
	}
	out := new(SSHHostsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookServerSpec) DeepCopyInto(out *WebhookServerSpec) {
	*out = *in
	in.Ingress.DeepCopyInto(&out.Ingress)
	in.Route.DeepCopyInto(&out.Route)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookServerSpec.
func (in *WebhookServerSpec) DeepCopy() *WebhookServerSpec {
	if in == nil {
		return nil
	}
	out := new(WebhookServerSpec)
	in.DeepCopyInto(out)
	return out
}
