package resource

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"server/devops-tools/etcd/initconfig"
	"time"
)

func GetCtx(key string) (getResponse *clientv3.GetResponse) {
	cli := initconfig.InitConfig()
	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
	getResponse, err := cli.Get(ctx, key)
	cancle()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	return getResponse
}
