package utils

import (
	"Configuration-center/global"
	"go.uber.org/zap"
	"os"
)

// 判断文件目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.CRC_LOG.Debug("create directory" + v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				global.CRC_LOG.Error("create directory"+v, zap.Any("error:", err))
			}
		}
	}
	return err
}
