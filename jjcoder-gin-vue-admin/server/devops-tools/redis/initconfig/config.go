package initconfig

import (
	"fmt"
	"github.com/go-redis/redis"
	"server/config"
)

func InitConfig() (cli *redis.Client) {
	cli = redis.NewClient(&redis.Options{
		Addr:     config.CrcConfig.Redis.Addr,
		Password: config.CrcConfig.Redis.Password,
		DB:       config.CrcConfig.Redis.DB,
	})
	_, err := cli.Ping().Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("redis connect success")
	return cli
}
