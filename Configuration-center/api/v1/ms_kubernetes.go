package v1

import (
	"Configuration-center/k8s"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags GetAssignPodList
// @Summary 获取指定Pod的列表
// @Description 通过给定的集群和namespace获取指定的deployment的pod List
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param ClusterName string
func GetAssignDeploymentInfo(context *gin.Context) {
	namespace := context.DefaultQuery("namespace", "default")
	deploymentName := context.Query("deploymentName")

	deployment := k8s.GetDeployment(namespace, deploymentName)
	fmt.Println(deployment.Status)
	context.JSON(http.StatusOK, deployment)
}

func ListDeployments(context *gin.Context) {
	namespace := context.DefaultQuery("namespace", "default")
	deploymentList := k8s.ListDeployment(namespace)
	context.JSON(http.StatusOK, deploymentList)
}
