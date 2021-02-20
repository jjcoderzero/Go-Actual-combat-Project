package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/devops-tools/jenkins/resource"
)

func GetNodeName(c *gin.Context) {
	nodeName := resource.GetNodes()
	c.JSON(http.StatusOK, nodeName)
}
