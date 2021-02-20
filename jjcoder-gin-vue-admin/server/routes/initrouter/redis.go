package initrouter

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

func InitRedisRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	RedisGroup := Router.Group("redis")
	{
		RedisGroup.POST("ctx/", v1.SetRedis)
		RedisGroup.GET("ctx/", v1.GetRedis)
	}
	return RedisGroup
}
