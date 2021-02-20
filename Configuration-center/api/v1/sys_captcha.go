package v1

import (
	"Configuration-center/global"
	"Configuration-center/model/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/captcha [post]
func Captcha(context *gin.Context) {
	driver := base64Captcha.NewDriverDigit(global.CRC_CONFIG.Captcha.ImgHeigth, global.CRC_CONFIG.Captcha.ImgWidth, global.CRC_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.CRC_LOG.Error("验证码获取失败", zap.Any("err", err))
		response.FailWithMessage("验证码获取失败", context)
	} else {
		response.OkWithDetailed(response.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", context)
	}

}
