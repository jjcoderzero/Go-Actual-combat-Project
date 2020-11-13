// +build !windows

package utils

import (
	"Configuration-center/global"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

func GetWriteSyncer() (zapcore.WriteSyncer, error)  {
	fileWrite, err := zaprotatelogs.New(
		path.Join(global.CRC_CONFIG.Zap.Director, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName(global.CRC_CONFIG.Zap.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
		)
	if global.CRC_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWrite)), err
	}
	return zapcore.AddSync(fileWrite), err
}
