package middleware

import (
	"Configuration-center/global"
	"Configuration-center/model"
	"Configuration-center/model/request"
	"Configuration-center/model/response"
	"Configuration-center/service"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法登录", context)
			context.Abort()
			return
		}
		if service.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", context)
			context.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", context)
				context.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), context)
			context.Abort()
			return
		}
		if err, _ = service.FindUserByUuid(claims.UUID.String()); err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), context)
			context.Abort()
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + 60*60*24*7
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			context.Header("new-token", newToken)
			context.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if global.CRC_CONFIG.System.UseMultipoint {
				err, RedisJwtToken := service.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.CRC_LOG.Error("get redis jwt failed", zap.Any("err", err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = service.JsonInBlacklist(model.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = service.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{SigningKey: []byte(global.CRC_CONFIG.JWT.SigningKey)}
}

func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
