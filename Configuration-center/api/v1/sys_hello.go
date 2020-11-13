package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
// @Tags HelloApi
// @Summary 创建Hello api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelloApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /api/v1/base/hello [GET]
func Hello(context *gin.Context) {
	context.String(http.StatusOK, "hello world")
}
