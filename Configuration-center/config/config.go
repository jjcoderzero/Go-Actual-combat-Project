package config

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	Local Local `mapstructure:"local" json:"local" yaml:"local"`

}