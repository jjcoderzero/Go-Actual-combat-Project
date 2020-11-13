package routes

import "github.com/gin-gonic/gin"

func InitKubernetesRouter(Router *gin.RouterGroup) (R *gin.IRouter)  {
	KubernetesGroup := Router.Group("k8s")
	{
		KubernetesGroup.POST("pods", v1.)
	}
}
