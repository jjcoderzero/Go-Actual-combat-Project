package initconfig

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"server/config"
)

func InitConfig() (db *gorm.DB) {
	user := config.CrcConfig.Mysql.User
	password := config.CrcConfig.Mysql.Password
	host := config.CrcConfig.Mysql.Host
	database := config.CrcConfig.Mysql.Database
	DRIVER := config.CrcConfig.Mysql.Driver
	DSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, database)
	db, err := gorm.Open(DRIVER, DSN)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	return db
}
