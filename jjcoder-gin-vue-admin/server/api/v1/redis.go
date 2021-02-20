package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/devops-tools/redis/resource"
	"strconv"
	"time"
)

func GetRedis(c *gin.Context) {
	key := c.Query("key")
	kvmap := resource.GetCtx(key)
	c.JSON(http.StatusOK, kvmap)
}

func SetRedis(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")
	expir := c.DefaultQuery("expir", "0")
	intexpir, err := strconv.ParseInt(expir, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	resource.SetCtx(key, value, time.Duration(intexpir))
	c.JSON(http.StatusOK, "set success")
}
