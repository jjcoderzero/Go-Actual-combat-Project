package resource

import (
	"fmt"
	"golang.org/x/net/context"
	"server/devops-tools/etcd/initconfig"
)

func WatchKey(key string) {
	kvmap := make(map[string]string)
	cli := initconfig.InitConfig()
	defer cli.Close()
	rch := cli.Watch(context.Background(), key)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			if ev.Type == 0 {
				kvmap[string(ev.Kv.Key)] = string(ev.Kv.Value)
				fmt.Println(kvmap)
			}
		}
	}
}
