package config

type Server struct {
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Local   Local   `mapstructure:"local" json:"local" yaml:"local"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
