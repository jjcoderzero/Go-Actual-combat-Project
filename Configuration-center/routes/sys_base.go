package routes

import (
	v1 "Configuration-center/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter)  {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("hello", v1.Hello)
	}
	return BaseRouter
}
