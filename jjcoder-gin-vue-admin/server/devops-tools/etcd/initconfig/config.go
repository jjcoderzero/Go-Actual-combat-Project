package initconfig

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"server/config"
)

func InitConfig() (cli *clientv3.Client) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{config.CrcConfig.Etcd.Url},
		DialTimeout: 0,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
	}
	fmt.Println("connect to etcd success")
	return cli
}
