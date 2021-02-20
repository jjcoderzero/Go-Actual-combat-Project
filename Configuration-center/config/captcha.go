package config

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`
	ImgHeigth int `mapstructure:"img-height" json:"imgHeigth" yaml:"img-height"`
}
