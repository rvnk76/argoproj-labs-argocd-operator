package common

// app-controller
const (
	AppControllerComponent = "application-controller"

	// ArgoCDApplicationControllerDefaultShardReplicas is the default number of replicas that the ArgoCD Application Controller Should Use
	ArgocdApplicationControllerDefaultReplicas = 1

	// ArgoCDDefaultControllerParellelismLimit is the default parallelism limit for application controller
	ArgoCDDefaultControllerParallelismLimit = int32(10)

	// ArgoCDDefaultControllerResourceLimitCPU is the default CPU limit when not specified for the Argo CD application
	// controller contianer.
	ArgoCDDefaultControllerResourceLimitCPU = "1000m"

	// ArgoCDDefaultControllerResourceLimitMemory is the default memory limit when not specified for the Argo CD
	// application controller contianer.
	ArgoCDDefaultControllerResourceLimitMemory = "64Mi"

	// ArgoCDDefaultControllerResourceRequestCPU is the default CPU requested when not specified for the Argo CD
	// application controller contianer.
	ArgoCDDefaultControllerResourceRequestCPU = "250m"

	// ArgoCDDefaultControllerResourceRequestMemory is the default memory requested when not specified for the Argo CD
	// application controller contianer.
	ArgoCDDefaultControllerResourceRequestMemory = "32Mi"
)
