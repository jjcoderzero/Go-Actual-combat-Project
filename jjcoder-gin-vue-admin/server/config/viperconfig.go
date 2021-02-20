package config

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

const (
	ConfigEnv  = "CRC_CONFIG"
	ConfigFile = "config.yaml"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file")
		flag.Parse()
		if config == "" {
			if configEnv := os.Getenv(ConfigEnv); configEnv == "" {
				config = ConfigFile
				fmt.Printf("You are using the default value of config, the path of config is: %v\n", ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("You are using the CRC_CONFIG environment variable, and the path of config is:%v\n", config)
			}
		} else {
			fmt.Printf("The value you are passing using the -c parameter on the command line, the path of config is:%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("You are using the value passed by Func Viper(), the path of config is:%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file:%v\n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err := v.Unmarshal(&CrcConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&CrcConfig); err != nil {
		fmt.Println(err)
	}
	return v
}
