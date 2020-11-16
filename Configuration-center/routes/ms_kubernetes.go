package routes

import (
	v1 "Configuration-center/api/v1"
	"github.com/gin-gonic/gin"
)

func InitKubernetesRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	KubernetesGroup := Router.Group("k8s")
	{
		KubernetesGroup.GET("deployment/", v1.GetAssignDeploymentInfo)
		KubernetesGroup.GET("deploymentList", v1.ListDeployments)
	}
	return KubernetesGroup
}
