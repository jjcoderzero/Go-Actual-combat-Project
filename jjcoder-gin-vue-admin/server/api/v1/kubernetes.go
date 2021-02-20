package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/devops-tools/kubernetes/resource"
)

//
func GetAssignDeploymentInfo(c *gin.Context) {
	namespace := c.DefaultQuery("namespace", "default")
	deploymentName := c.Query("deploymentName")
	deployment := resource.GetDeployment(namespace, deploymentName)
	c.JSON(http.StatusOK, deployment)
}

func ListDeployments(c *gin.Context) {
	namespace := c.DefaultQuery("namespace", "default")
	deploymentList := resource.ListDeployment(namespace)
	c.JSON(http.StatusOK, deploymentList)
}
