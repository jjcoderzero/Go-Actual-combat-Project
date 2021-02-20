package resource

import (
	"context"
	"k8s.io/api/apps/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"server/devops-tools/kubernetes/initconfig"
)

var clientset *kubernetes.Clientset

func init() {
	kubeconfig := initconfig.InitConfig()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

func GetDeployment(namespace, deploymentName string) (deployment *v1beta1.Deployment) {
	deploymentClient := clientset.AppsV1beta1().Deployments(namespace)
	deployment, err := deploymentClient.Get(context.TODO(), deploymentName, v1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	return deployment
}

func ListDeployment(namespace string) (deploymentList *v1beta1.DeploymentList) {
	deploymentClient := clientset.AppsV1beta1().Deployments(namespace)
	deploymentList, err := deploymentClient.List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	return deploymentList
}
