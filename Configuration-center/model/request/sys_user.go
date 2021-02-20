package request

import uuid "github.com/satori/go.uuid"

type Register struct {
	UserName    string `json:"userName"`
	Password    string `json:"passWord"`
	NickName    string `json:"nickName"`
	HeaderImg   string `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	AuthorityId string `json:"authorityId" gorm:"default:888"`
}

type Login struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

type ChangePasswordStruct struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`
	AuthorityId string    `json:"authorityId"`
}
