package resource

import "server/devops-tools/jenkins/initconfig"

func GetNodes() (nodeName string) {
	jenkins := initconfig.InitConfig()
	nodes, _ := jenkins.GetAllNodes()
	for _, node := range nodes {
		node.Poll()
		if ok, _ := node.IsOnline(); ok {
			nodeName = node.GetName()
		}
	}
	return nodeName
}
