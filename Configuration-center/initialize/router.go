package initialize

import (
	_ "Configuration-center/docs"
	"Configuration-center/global"
	"Configuration-center/middleware"
	"Configuration-center/routes"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	//Router.StaticFS(global.CRC_CONFIG.Local.Path, http.Dir(global.CRC_CONFIG.Local.Path))
	//Router.Use(middleware.LoadTls())
	global.CRC_LOG.Info("use middleware logger")
	Router.Use(middleware.Cors()) // 跨域
	global.CRC_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.CRC_LOG.Info("register swagger handler")

	ApiV1Group := Router.Group("/api/v1")
	routes.InitBaseRouter(ApiV1Group)
	routes.InitKubernetesRouter(ApiV1Group)
	global.CRC_LOG.Info("router register success")
	return Router
}
