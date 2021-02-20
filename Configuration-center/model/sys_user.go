package model

import (
	"Configuration-center/global"
	"github.com/satori/go.uuid"
)

type SysUser struct {
	global.CRC_MODEL
	UUID        uuid.UUID    `json:"uuid" gorm:"comment:用户UUID"`
	UserName    string       `json:"userName" gorm:"comment:用户登录名"`
	Password    string       `json:"-" gorm:"comment:用户登录密码"`
	NickName    string       `json:"nickName" gorm:"default:系统用户";comment:"用户昵称"`
	HeaderImag  string       `json:"headerImag" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
}
