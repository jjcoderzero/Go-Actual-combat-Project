package v1

import (
	"Configuration-center/global"
	"Configuration-center/middleware"
	"Configuration-center/model"
	"Configuration-center/model/request"
	"Configuration-center/model/response"
	"Configuration-center/service"
	"Configuration-center/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Base
// @Summary 用户登录
// @Produce application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
//@Router /base/login [post]
func Login(context *gin.Context) {
	var L request.Login
	_ = context.ShouldBindJSON(&L)
	if err := utils.Verify(L, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	if store.Verify(L.CaptchaId, L.Captcha, true) {
		U := &model.SysUser{UserName: L.Username, Password: L.Password}
		if err, user := service.Login(U); err != nil {
			global.CRC_LOG.Error("登录失败！用户名不存在或者密码错误", zap.Any("err", err))
			response.FailWithMessage("用户名不存在或者密码错误", context)
		} else {
			tokenNext(context, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", context)
	}
}

func tokenNext(context *gin.Context, user model.SysUser) {
	j := &middleware.JWT()
}

//@Tags SysUser
