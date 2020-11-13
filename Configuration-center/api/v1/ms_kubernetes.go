package v1


// @Tags GetAssignPodList
// @Summary 获取指定Pod的列表
// @Description 通过给定的集群和namespace获取指定的deployment的pod List
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param ClusterName string
func GetAssignPodList(ClusterName string, NameSpace string, DeploymentName string) (PodList []string) { //
	
}
