package redis

import (
	"fmt"
	"reflect"
	"strconv"

	argoproj "github.com/argoproj-labs/argocd-operator/api/v1beta1"
	"github.com/argoproj-labs/argocd-operator/common"
	"github.com/argoproj-labs/argocd-operator/controllers/argocd/argocdcommon"
	"github.com/argoproj-labs/argocd-operator/pkg/argoutil"
	"github.com/argoproj-labs/argocd-operator/pkg/util"
	"github.com/argoproj-labs/argocd-operator/pkg/workloads"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

func (rr *RedisReconciler) TLSVerificationDisabled() bool {
	return rr.Instance.Spec.Redis.DisableTLSVerification
}

// UseTLS decides whether Redis component should communicate with TLS or not
func (rr *RedisReconciler) UseTLS() bool {
	tlsSecret, err := workloads.GetSecret(common.ArgoCDRedisServerTLSSecretName, rr.Instance.Namespace, rr.Client)
	if err != nil {
		rr.Logger.Error(err, "UseTLS: failed to retrieve tls secret", "name", common.ArgoCDRedisServerTLSSecretName, "namespace", rr.Instance.Namespace)
		rr.Logger.V(1).Info("useTLS: skipping TLS enforcement")
		return false
	}

	secretOwner, err := argocdcommon.FindSecretOwnerInstance(types.NamespacedName{Name: tlsSecret.Name, Namespace: tlsSecret.Namespace}, rr.Client)
	if err != nil {
		rr.Logger.Error(err, "UseTLS: failed to find secret owning instance")
		return false
	}

	if !reflect.DeepEqual(secretOwner, types.NamespacedName{}) {
		return true
	}

	return false
}

// getRedisServerAddress will return the Redis service address for the given ArgoCD instance
func (rr *RedisReconciler) GetServerAddress() string {
	if rr.Instance.Spec.Redis.Remote != nil && *rr.Instance.Spec.Redis.Remote != "" {
		return *rr.Instance.Spec.Redis.Remote
	}
	if rr.Instance.Spec.HA.Enabled {
		return rr.GetHAProxyAddress()
	}
	return argoutil.FqdnServiceRef(resourceName, rr.Instance.Namespace, common.DefaultRedisPort)
}

// GetContainerImage will return the container image for the Redis server.
func (rr *RedisReconciler) GetContainerImage() string {
	fn := func(cr *argoproj.ArgoCD) (string, string) {
		return cr.Spec.Redis.Image, cr.Spec.Redis.Version
	}
	return argocdcommon.GetContainerImage(fn, rr.Instance, common.RedisImageEnvVar, common.DefaultRedisImage, common.DefaultRedisVersion)
}

// GetResources will return the ResourceRequirements for the Redis container.
func (rr *RedisReconciler) GetResources() corev1.ResourceRequirements {
	resources := corev1.ResourceRequirements{}

	// Allow override of resource requirements from CR
	if rr.Instance.Spec.Redis.Resources != nil {
		resources = *rr.Instance.Spec.Redis.Resources
	}
	return resources
}

// GetConf will load the redis configuration from a template on disk for the given ArgoCD.
// If an error occurs, an empty string value will be returned.
func (rr *RedisReconciler) GetConf() string {
	path := fmt.Sprintf("%s/%s", getConfigPath(), redisConfTpl)
	params := map[string]string{
		UseTLSKey: strconv.FormatBool(rr.TLSEnabled),
	}

	conf, err := util.LoadTemplateFile(path, params)
	if err != nil {
		rr.Logger.Error(err, "GetConf: failed to load redis configuration")
		return ""
	}
	return conf
}

// GetInitScript will load the redis init script from a template on disk for the given ArgoCD.
// If an error occurs, an empty string value will be returned.
func (rr *RedisReconciler) GetInitScript() string {
	path := fmt.Sprintf("%s/%s", getConfigPath(), initShTpl)
	params := map[string]string{
		ServiceNameKey: HAResourceName,
		UseTLSKey:      strconv.FormatBool(rr.TLSEnabled),
	}

	script, err := util.LoadTemplateFile(path, params)
	if err != nil {
		rr.Logger.Error(err, "GetInitScript: failed to load redis init script")
		return ""
	}
	return script
}

// GetLivenessScript will load the redis liveness script from a template on disk for the given ArgoCD.
// If an error occurs, an empty string value will be returned.
func (rr *RedisReconciler) GetLivenessScript() string {
	path := fmt.Sprintf("%s/%s", getConfigPath(), livenessShTpl)
	params := map[string]string{
		UseTLSKey: strconv.FormatBool(rr.TLSEnabled),
	}
	script, err := util.LoadTemplateFile(path, params)
	if err != nil {
		rr.Logger.Error(err, "GetLivenessScript: failed to load redis liveness script")
		return ""
	}
	return script
}

// GetReadinessScript will load the redis readiness script from a template on disk for the given ArgoCD.
// If an error occurs, an empty string value will be returned.
func (rr *RedisReconciler) GetReadinessScript() string {
	path := fmt.Sprintf("%s/%s", getConfigPath(), readinessShTpl)
	params := map[string]string{
		UseTLSKey: strconv.FormatBool(rr.TLSEnabled),
	}
	script, err := util.LoadTemplateFile(path, params)
	if err != nil {
		rr.Logger.Error(err, "GetLivenessScript: failed to load redis readiness script")
		return ""
	}
	return script
}

// GetSentinelConf will load the redis sentinel configuration from a template on disk for the given ArgoCD.
// If an error occurs, an empty string value will be returned.
func (rr *RedisReconciler) GetSentinelConf() string {
	path := fmt.Sprintf("%s/%s", getConfigPath(), sentinelConfTpl)
	params := map[string]string{
		UseTLSKey: strconv.FormatBool(rr.TLSEnabled),
	}

	conf, err := util.LoadTemplateFile(path, params)
	if err != nil {
		rr.Logger.Error(err, "GetSentinelConf: failed to load redis sentinel configuration")
		return ""
	}
	return conf
}

// GetSentinelLivenessScript will load the redis liveness script from a template on disk for the given ArgoCD.
// If an error occurs, an empty string value will be returned.
func (rr *RedisReconciler) GetSentinelLivenessScript() string {
	path := fmt.Sprintf("%s/%s", getConfigPath(), sentinelLivenessShTpl)
	params := map[string]string{
		UseTLSKey: strconv.FormatBool(rr.TLSEnabled),
	}

	script, err := util.LoadTemplateFile(path, params)
	if err != nil {
		rr.Logger.Error(err, "GetSentinelConf: failed to load sentinel liveness script")
		return ""
	}
	return script
}

// getConfigPath will return the path for the Redis configuration templates.
func getConfigPath() string {
	path := common.DefaultRedisConfigPath
	if _, val := util.CaseInsensitiveGetenv(common.RedisConfigPathEnvVar); val != "" {
		path = val
	}
	return path
}
