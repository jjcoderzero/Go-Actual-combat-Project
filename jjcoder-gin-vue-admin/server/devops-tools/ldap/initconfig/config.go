package initconfig

import (
	"crypto/tls"
	"github.com/go-ldap/ldap/v3"
	"server/config"
	"time"
)

func InitConfig() (conn *ldap.Conn) {
	var (
		err error
	)
	if config.CrcConfig.Ldap.Port == "636" {
		conn, err = ldap.DialTLS("tcp", config.CrcConfig.Ldap.Host+":"+config.CrcConfig.Ldap.Port, &tls.Config{InsecureSkipVerify: false})
	} else {
		conn, err = ldap.Dial("tcp", config.CrcConfig.Ldap.Host+":"+config.CrcConfig.Ldap.Port)
	}
	if err != nil {
		err.Error()
	}
	conn.SetTimeout(5 * time.Second)
	defer conn.Close()
	return conn
}
