package main

import (
	"server/config"
	"server/ginengine"
)

func main() {
	config.CrcViper = config.Viper()
	ginengine.RunServer()
}
