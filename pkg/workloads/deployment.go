package workloads

import (
	"errors"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	cntrlClient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/argoproj-labs/argocd-operator/pkg/mutation"
	"github.com/argoproj-labs/argocd-operator/pkg/resource"
)

// DeploymentRequest objects contain all the required information to produce a deployment object in return
type DeploymentRequest struct {
	ObjectMeta metav1.ObjectMeta
	Spec       appsv1.DeploymentSpec

	// array of functions to mutate role before returning to requester
	Mutations []mutation.MutateFunc
	Client    cntrlClient.Client
}

// newDeployment returns a new Deployment instance for the given ArgoCD.
func newDeployment(objMeta metav1.ObjectMeta, spec appsv1.DeploymentSpec) *appsv1.Deployment {

	return &appsv1.Deployment{
		ObjectMeta: objMeta,
		Spec:       spec,
	}
}

func RequestDeployment(request DeploymentRequest) (*appsv1.Deployment, error) {
	var (
		mutationErr error
	)
	deployment := newDeployment(request.ObjectMeta, request.Spec)

	if len(request.Mutations) > 0 {
		for _, mutation := range request.Mutations {
			err := mutation(nil, deployment, request.Client)
			if err != nil {
				mutationErr = err
			}
		}
		if mutationErr != nil {
			return deployment, fmt.Errorf("RequestDeployment: one or more mutation functions could not be applied: %s", mutationErr)
		}
	}

	return deployment, nil
}

// TriggerDeploymentRollout will update the label with the given key to trigger a new rollout of the Deployment.
func TriggerDeploymentRollout(client cntrlClient.Client, name, namespace string, updateChangedTime func(name, namespace string)) error {
	currentDeployment, err := GetDeployment(name, namespace, client)
	if err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
		return nil
	}

	updateChangedTime(currentDeployment.Name, currentDeployment.Namespace)
	return UpdateDeployment(currentDeployment, client)
// CreateDeployment creates the specified Deployment using the provided client.
func CreateDeployment(deployment *appsv1.Deployment, client cntrlClient.Client) error {
	return resource.CreateObject(deployment, client)
}

// UpdateDeployment updates the specified Deployment using the provided client.
func UpdateDeployment(deployment *appsv1.Deployment, client cntrlClient.Client) error {
	return resource.UpdateObject(deployment, client)
}

// DeleteDeployment deletes the Deployment with the given name and namespace using the provided client.
func DeleteDeployment(name, namespace string, client cntrlClient.Client) error {
	deployment := &appsv1.Deployment{}
	return resource.DeleteObject(name, namespace, deployment, client)
}

// GetDeployment retrieves the Deployment with the given name and namespace using the provided client.
func GetDeployment(name, namespace string, client cntrlClient.Client) (*appsv1.Deployment, error) {
	deployment := &appsv1.Deployment{}
	obj, err := resource.GetObject(name, namespace, deployment, client)
	if err != nil {
		return nil, err
	}
	// Assert the object as an appsv1.Deployment
	deployment, ok := obj.(*appsv1.Deployment)
	if !ok {
		return nil, errors.New("failed to assert the object as an appsv1.Deployment")
	}
	return deployment, nil
}

// ListDeployments returns a list of Deployment objects in the specified namespace using the provided client and list options.
func ListDeployments(namespace string, client cntrlClient.Client, listOptions []cntrlClient.ListOption) (*appsv1.DeploymentList, error) {
	deploymentList := &appsv1.DeploymentList{}
	obj, err := resource.ListObjects(namespace, deploymentList, client, listOptions)
	if err != nil {
		return nil, err
	}
	// Assert the object as an appsv1.DeploymentList
	deploymentList, ok := obj.(*appsv1.DeploymentList)
	if !ok {
		return nil, errors.New("failed to assert the object as an appsv1.DeploymentList")
	}
	return deploymentList, nil
}
