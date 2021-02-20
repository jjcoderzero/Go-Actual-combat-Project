package global

import (
	"Configuration-center/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CRC_CONFIG config.Server
	CRC_VIPER  *viper.Viper
	CRC_LOG    *zap.Logger
	CRC_DB     *gorm.DB
	CRC_REDIS  *redis.Client
)
