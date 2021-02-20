package service

import (
	"Configuration-center/global"
	"Configuration-center/model"
	"Configuration-center/utils"
)

//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser
func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.CRC_DB.Where("username = ? AND password = ?", u.UserName, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}
