package model

import "Configuration-center/global"

type JwtBlacklist struct {
	global.CRC_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
