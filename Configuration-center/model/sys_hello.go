package model

import "Configuration-center/global"

type HelloApi struct {
	global.CRC_MODEL
	Path        string `json:"path" gorm:"comment:api路径"`
	Description string `json:"description" gorm:"comment:api中文描述"`
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组"`
	Method      string `json:"method" gorm:"default:GET" gorm:"comment:方法"`
}
