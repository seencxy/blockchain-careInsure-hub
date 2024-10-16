package utils

import (
	"backed/app/core/consts"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type UserToken struct {
	jwt.StandardClaims
	//自定义的用户信息
	UserName string `json:"user_name"`
}

// GenToken 生成token
func GenToken(UserName string, Expires time.Duration, secret_key []byte) (string, error) {
	user := UserToken{
		jwt.StandardClaims{
			//过期时间：现在 + 加上传的过期时间  // 过期时间1小时，n小时的话 * n
			ExpiresAt: time.Now().Add(Expires).Unix(),
			//签名
			Issuer: "BlockchainCareInsureHub",
		},
		UserName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)

	return token.SignedString(secret_key)
}

// AuthToken 认证token
func AuthToken(tokenString string, secretKey []byte) (*UserToken, error) {

	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserToken{}, func(token *jwt.Token) (key interface{}, err error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	clasims, is_ok := token.Claims.(*UserToken)

	//验证token
	if token.Valid && is_ok {
		return clasims, nil
	}

	return nil, errors.New("token valid err")
}

// GetUserInfoFromContext  从上下文中获取用户信息
func GetUserInfoFromContext(ctx *gin.Context) (string, error) {
	//  从请求头部中获取auth
	auth_header := ctx.Request.Header.Get("Authorization")
	auths := strings.Split(auth_header, " ")
	token := auths[1]
	user, err := AuthToken(token, []byte(consts.TokenSecret))
	if err != nil {
		return "", err
	}

	return user.UserName, nil
}
