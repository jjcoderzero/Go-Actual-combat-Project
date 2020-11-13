package k8s

import (
	"Configuration-center/k8s/initialize"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)



func GetPods()  {
	kubeconfig := initialize.InitConfig()
	config, err := clientcmd.BuildConfigFromFlags("192.168.99.103:6443", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)

	clientDeployment := clientset.CoreV1()
}