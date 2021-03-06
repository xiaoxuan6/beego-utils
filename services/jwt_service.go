package services

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	gjwt "github.com/golang-jwt/jwt"
	"github.com/xiaoxuan6/beego-utils/jwt"
	"strings"
	"time"
)

type JWT struct {
}

var key string
var expiresAt int64

var (
	ErrTokenExpired    error = fmt.Errorf("令牌已过期")
	ErrTokenInvalid    error = fmt.Errorf("请求令牌无效")
	ErrHeaderEmpty     error = fmt.Errorf("需要认证才能访问！")
	ErrHeaderMalformed error = fmt.Errorf("请求头中 Authorization 格式有误")
)

var JwtService = newJWTService()

func newJWTService() *JWT {
	return new(JWT)
}

func init() {
	key = beego.AppConfig.String("jwt_secret")
	expiresAt, _ = beego.AppConfig.Int64("expires_at")
}

// GenerateToken 生成 token
func (j *JWT) GenerateToken(id int64, userName string) (token string, err error) {
	claims := jwt.MyClaims{
		ID:       id,
		UserName: userName,
		StandardClaims: gjwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expiresAt) * time.Second).Unix(),
		},
	}

	return jwt.GenerateToken(claims, key)
}

// Parse 使用 jwt.ParseTokenString 解析 Token
func (j *JWT) Parse(ctx *context.Context) (claims gjwt.MapClaims, error error) {
	token, err := getTokenFromHeader(ctx)

	if err != nil {
		return nil, err
	}

	c, e := jwt.ParseTokenString(token, key)

	if e != nil {
		validationErr, ok := e.(gjwt.ValidationError)

		if ok {
			if validationErr.Errors == gjwt.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}

		return nil, e
	}

	if cl, ok := c.(gjwt.MapClaims); ok {
		return cl, nil
	}

	return nil, ErrTokenInvalid
}

// ParseWithClaims 使用 jwt.ParseWithClaims 解析 Token
func (j *JWT) ParseWithClaims(ctx *context.Context) (*jwt.MyClaims, error) {
	token, err := getTokenFromHeader(ctx)

	if err != nil {
		return nil, err
	}

	c, e := jwt.ParseWithClaims(token, &jwt.MyClaims{}, key)
	if e != nil {
		validationErr, ok := e.(gjwt.ValidationError)

		if ok {
			if validationErr.Errors == gjwt.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}

		return nil, e
	}

	return c, nil
}

// 获取 header 中的 Authorization:Bearer xxxxx
func getTokenFromHeader(ctx *context.Context) (string, error) {
	Authorization := ctx.Input.Header("Authorization")

	if Authorization == "" {
		return "", ErrHeaderEmpty
	}

	parts := strings.SplitN(Authorization, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}

	return parts[1], nil
}
