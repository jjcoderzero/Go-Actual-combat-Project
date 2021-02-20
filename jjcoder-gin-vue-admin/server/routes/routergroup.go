package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "server/docs"
	"server/middleware"
	"server/routes/initrouter"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.CrossDomain())
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ApiV1Group := Router.Group("/api/v1")
	initrouter.InitKubernetesRouter(ApiV1Group)
	initrouter.InitJenkinsRouter(ApiV1Group)
	initrouter.InitEtcdRouter(ApiV1Group)
	initrouter.InitRedisRouter(ApiV1Group)
	initrouter.InitGithubRouter(ApiV1Group)
	initrouter.InitMysqlRouter(ApiV1Group)
	fmt.Println("router register success")
	return Router
}
