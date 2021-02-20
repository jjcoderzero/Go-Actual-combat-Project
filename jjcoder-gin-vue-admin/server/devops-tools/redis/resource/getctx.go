package resource

import (
	"fmt"
	"github.com/go-redis/redis"
	"server/devops-tools/redis/initconfig"
)

func GetCtx(key string) (kvmap map[string]string) {
	cli := initconfig.InitConfig()
	kvmap = make(map[string]string)
	value, err := cli.Get(key).Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		kvmap[key] = value
	}
	return kvmap
}
