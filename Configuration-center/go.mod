module Configuration-center

go 1.14

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/fsnotify/fsnotify v1.4.9
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.6.3
	github.com/go-openapi/spec v0.19.12 // indirect
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.3 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/mojocn/base64Captcha v1.3.1
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.7.1
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.6.9
	github.com/tebeka/strftime v0.1.5 // indirect
	github.com/unrolled/secure v1.0.8
	go.uber.org/zap v1.10.0
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/tools v0.0.0-20201111224557-41a3a589386c // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.6
	k8s.io/api v0.19.3
	k8s.io/apimachinery v0.19.3
	k8s.io/client-go v0.19.3
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920 // indirect
)

replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.3.0
