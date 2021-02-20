package initrouter

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

func InitJenkinsRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	JenkinsGroup := Router.Group("jenkins")
	{
		JenkinsGroup.GET("node/", v1.GetNodeName)
	}
	return JenkinsGroup
}
