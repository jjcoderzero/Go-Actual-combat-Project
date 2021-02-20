package initconfig

import (
	"github.com/bndr/gojenkins"
	"log"
	"server/config"
)

func InitConfig() (jenkins *gojenkins.Jenkins) {
	jenkins = gojenkins.CreateJenkins(nil, config.CrcConfig.Jenkins.Url, config.CrcConfig.Jenkins.UserName, config.CrcConfig.Jenkins.Password)
	// Provide CA certificate if server is using self-signed certificate
	// caCert, _ := ioutil.ReadFile("/tmp/ca.crt")
	// jenkins.Requester.CACert = caCert
	_, err := jenkins.Init()

	if err != nil {
		log.Printf("连接Jenkins失败, %v\n", err)
		return
	}
	log.Println("Jenkins连接成功")
	return jenkins
}
