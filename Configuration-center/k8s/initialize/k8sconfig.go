package initialize

import (
	"flag"
	"fmt"
)

func InitConfig() (kubconfig *string) {
	configPath := fmt.Sprintf("%s", "/Users/wanj016/Desktop/source-code-reading/Go-Actual-combat-Project/Configuration-center/k8s/kube/config")
	kubeconfig := flag.String("kubeconfig", configPath, "(optional) absolute path to the kubeconfig file")
	return kubeconfig
}
