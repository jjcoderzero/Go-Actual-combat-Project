package resource

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"server/devops-tools/etcd/initconfig"
	"time"
)

func DeleteCtx(key string) (deleteResponse *clientv3.DeleteResponse) {
	cli := initconfig.InitConfig()
	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
	deleteResponse, err := cli.Delete(ctx, key)
	cancle()
	if err != nil {
		fmt.Printf("delete from etcd failed, err:%v\n", err)
		return
	}
	return deleteResponse
}
