package initrouter

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

func InitGithubRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	GithubGroup := Router.Group("github")
	{
		GithubGroup.POST("repo/", v1.NewRepo)
		GithubGroup.GET("repo/", v1.ListRepos)
		GithubGroup.POST("custom-webhook-trigger/", v1.CustomWebHooks)
	}
	return GithubGroup
}
