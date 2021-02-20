package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/devops-tools/mysql/resource"
	"strconv"
)

func CreateProduct(c *gin.Context) {
	code := c.Query("code")
	price := c.Query("price")
	priceint, _ := strconv.Atoi(price)
	resource.CreateProduct(code, priceint)
	c.String(http.StatusOK, "创建成功")
}

func ReadProduct(c *gin.Context) {
	code := c.Query("code")
	product := resource.ReadProduct(code)
	c.JSON(http.StatusOK, product)
}
