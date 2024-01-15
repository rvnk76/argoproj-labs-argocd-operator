package redis

import (
	"github.com/argoproj-labs/argocd-operator/common"
	"github.com/argoproj-labs/argocd-operator/controllers/argocd/argocdcommon"
	"github.com/argoproj-labs/argocd-operator/pkg/argoutil"
	"github.com/argoproj-labs/argocd-operator/pkg/mutation"
	"github.com/argoproj-labs/argocd-operator/pkg/util"
	"github.com/argoproj-labs/argocd-operator/pkg/workloads"
	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (rr *RedisReconciler) reconcileDeployment() error {

}

func (rr *RedisReconciler) getDeploymentRequest() workloads.DeploymentRequest {
	depReq := workloads.DeploymentRequest{
		ObjectMeta: argoutil.GetObjMeta(resourceName, rr.Instance.Namespace, rr.Instance.Name, rr.Instance.Namespace, component),
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					common.AppK8sKeyName: resourceName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						common.AppK8sKeyName: resourceName,
					},
				},
				Spec: corev1.PodSpec{
					NodeSelector:       common.DefaultNodeSelector(),
					Containers:         rr.getDeploymentContainers(),
					ServiceAccountName: resourceName,
					Volumes: []corev1.Volume{
						{
							Name: common.ArgoCDRedisServerTLSSecretName,
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName: common.ArgoCDRedisServerTLSSecretName,
									Optional:   util.BoolPtr(true),
								},
							},
						},
					},
				},
			},
		},
		Instance:  rr.Instance,
		Mutations: []mutation.MutateFunc{mutation.ApplyReconcilerMutation},
	}

	return depReq
}

func (rr *RedisReconciler) getDeploymentContainers() []corev1.Container {
	containers := []corev1.Container{}

	redisContainer := corev1.Container{
		Args:            rr.getArgs(),
		Image:           rr.getContainerImage(),
		ImagePullPolicy: corev1.PullAlways,
		Name:            redisName,
		Ports: []corev1.ContainerPort{
			{
				ContainerPort: common.DefaultRedisPort,
			},
		},
		Resources: rr.getResources(),
		Env:       util.ProxyEnvVars(),
		SecurityContext: &corev1.SecurityContext{
			AllowPrivilegeEscalation: util.BoolPtr(allowPrivilegeEscalation),
			Capabilities: &corev1.Capabilities{
				Drop: []corev1.Capability{
					"ALL",
				},
			},
			RunAsNonRoot: util.BoolPtr(runAsNonRoot),
			RunAsUser:    util.Int64Ptr(999),
		},
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      common.ArgoCDRedisServerTLSSecretName,
				MountPath: TLSPath,
			},
		},
	}

	containers = append(containers, redisContainer)
	return containers
}

func (rr *RedisReconciler) getDesiredHAProxyDeployment() *appsv1.Deployment {
	desiredDeployment := &appsv1.Deployment{}

	return desiredDeployment
}

// TriggerDeploymentRollout starts redis deployment rollout by updating the given key
func (rr *RedisReconciler) TriggerDeploymentRollout(name, namespace, key string) error {
	return argocdcommon.TriggerDeploymentRollout(name, namespace, key, rr.Client)
}

func (rr *RedisReconciler) deleteDeployment(name, namespace string) error {
	if err := workloads.DeleteDeployment(name, namespace, rr.Client); err != nil {
		if apierrors.IsNotFound(err) {
			return nil
		}
		return errors.Wrapf(err, "deleteDeployment: failed to delete deployment %s", name)
	}
	rr.Logger.V(0).Info("deleteDeployment: deployment deleted", "name", name, "namespace", namespace)
	return nil
}
