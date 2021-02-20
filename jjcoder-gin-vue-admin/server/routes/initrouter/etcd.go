package initrouter

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

func InitEtcdRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	EtcdGroup := Router.Group("etcd")
	{
		EtcdGroup.POST("ctx/", v1.PutEtcd)
		EtcdGroup.GET("ctx/", v1.GetEtcd)
		EtcdGroup.GET("watchctx/", v1.WatchEtcd)
		EtcdGroup.DELETE("ctx/", v1.DeleteEtcd)

	}
	return EtcdGroup
}
