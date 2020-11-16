package initialize

import (
	"flag"
	"os/user"
	"path/filepath"
)

func InitConfig() (kubconfig *string) {
	var kubeconfig *string
	if home := GetHomePath(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	//configPath := fmt.Sprintf("%s", "/Users/wanj016/Desktop/source-code-reading/Go-Actual-combat-Project/Configuration-center/k8s/kube/config")
	//kubeconfig = flag.String("kubeconfig", configPath, "(optional) absolute path to the kubeconfig file")

	flag.Parse()
	return kubeconfig
}

func GetHomePath() string {
	u, err := user.Current()
	if err == nil {
		return u.HomeDir
	}
	return ""
}
