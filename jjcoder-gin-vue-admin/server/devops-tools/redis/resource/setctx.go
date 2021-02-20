package resource

import (
	"fmt"
	"server/devops-tools/redis/initconfig"
	"time"
)

func SetCtx(key string, value string, expir time.Duration) {
	cli := initconfig.InitConfig()
	err := cli.Set(key, value, expir).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}
}
