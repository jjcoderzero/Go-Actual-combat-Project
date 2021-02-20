package initrouter

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

func InitMysqlRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	MysqlGroup := Router.Group("mysql")
	{
		MysqlGroup.POST("product/", v1.CreateProduct)
		MysqlGroup.GET("product/", v1.ReadProduct)
	}
	return MysqlGroup
}
