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

func (j *JWT) GenerateToken(id int64, userName string) (token string, err error) {
	claims := jwt.Claims{
		ID:       id,
		UserName: userName,
		StandardClaims: gjwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expiresAt)).Unix(),
		},
	}

	return jwt.GenerateToken(&claims, key)
}

func (j *JWT) Parse(ctx context.Context) (claims *jwt.Claims, error error) {
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

	if claims, ok := c.(*jwt.Claims); ok {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// 获取 header 中的 Authorization:Bearer xxxxx
func getTokenFromHeader(ctx context.Context) (string, error) {
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
