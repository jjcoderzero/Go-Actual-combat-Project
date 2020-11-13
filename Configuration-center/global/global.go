package global

import (
	"Configuration-center/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	CRC_CONFIG config.Server
	CRC_VIPER  *viper.Viper
	CRC_LOG    *zap.Logger
)
