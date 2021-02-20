package core

import (
	"Configuration-center/global"
	"Configuration-center/initialize"
	"fmt"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initialize.Routers()
	//Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", global.CRC_CONFIG.System.Addr)
	s := initServer(address, Router)
	global.CRC_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf(`
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
`, address)
	global.CRC_LOG.Error(s.ListenAndServe().Error())
}
