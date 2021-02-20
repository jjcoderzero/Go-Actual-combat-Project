package initrouter

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

func InitKubernetesRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	KubernetesGroup := Router.Group("k8s")
	{
		KubernetesGroup.GET("deployment/", v1.GetAssignDeploymentInfo)
		KubernetesGroup.GET("deploymentlist/", v1.ListDeployments)
	}
	return KubernetesGroup
}
