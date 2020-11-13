package main

import (
	"Configuration-center/core"
	"Configuration-center/global"
)

func main() {
	global.CRC_VIPER = core.Viper()
	global.CRC_LOG = core.Zap()
	core.RunWindowsServer()
}
