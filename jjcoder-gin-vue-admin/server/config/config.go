package config

import "github.com/spf13/viper"

// global variable
var (
	CrcConfig Config
	CrcViper  *viper.Viper
)

type Config struct {
	Server  Server  `mapstructure:"server" json:"server" yaml:"server"`
	Jenkins Jenkins `mapstructure:"jenkins" json:"jenkins" yaml:"jenkins"`
	Etcd    Etcd    `mapstructure:"etcd" json:"etcd" yaml:"etcd"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Github  Github  `mapstructure:"github" json:"github" yaml:"github"`
	Ldap    Ldap    `mapstructure:"ldap" json:"ldap" yaml:"ldap"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

type Server struct {
	ServerPort uint `mapstructure:"serverPort" json:"serverPort" yaml:"serverPort"`
}

type Mysql struct {
	Driver   string `mapstructure:"dirver" json:"driver" yaml:"dirver"`
	User     string `mapstructure:"user" json:"user" yaml:"user"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
}

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}

type Kubernetes struct {
}

type Ldap struct {
	Host           string         `mapstructure:"host" json:"host" yaml:"host"`
	Port           string         `mapstructure:"port" json:"port" yaml:"port"`
	BaseDn         string         //基础DN
	LdapAttributes LdapAttributes `mapstructure:"ldapattributes" json:"ldapattributes" yaml:"ldapattributes"` //结果集
}

type LdapAttributes struct {
	UName      string `mapstructure:"uname" json:"uname" yaml:"uname"`                //ldap中用户名的key
	Name       string `mapstructure:"name" json:"name" yaml:"name"`                   //ldap中姓名的key
	Email      string `mapstructure:"email" json:"email" yaml:"email"`                //ldap中email的key
	Department string `mapstructure:"department" json:"department" yaml:"department"` //ldap中部门的key
}

type Jenkins struct {
	Url      string `mapstructure:"url" json:"url" yaml:"url"`
	UserName string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
type Etcd struct {
	Url string `mapstructure:"url" json:"url" yaml:"url"`
}

type Github struct {
	AccessToken string `mapstructure:"accesstoken" json:"accesstoken" yaml:"accesstoken"`
}
