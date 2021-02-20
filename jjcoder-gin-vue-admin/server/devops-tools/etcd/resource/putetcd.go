package resource

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"server/devops-tools/etcd/initconfig"
	"time"
)

func PutCtx(key, value string) (putResponse *clientv3.PutResponse) {
	cli := initconfig.InitConfig()
	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
	putResponse, err := cli.Put(ctx, key, value)
	cancle()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	return putResponse
}
