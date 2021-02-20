package ginengine

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"server/config"
	"server/routes"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	router := routes.Routers()
	address := fmt.Sprintf(":%d", config.CrcConfig.Server.ServerPort)
	s := initServer(address, router)
	fmt.Printf(`
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
`, address)
	logrus.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	server := endless.NewServer(address, router)
	server.ReadHeaderTimeout = 10 * time.Millisecond
	server.WriteTimeout = 10 * time.Second
	server.MaxHeaderBytes = 1 << 20
	return server
}
