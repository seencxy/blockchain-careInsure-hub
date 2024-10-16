package middleware

import (
	"backed/app/core/consts"
	"backed/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// JwtTokenValid 验证token
func JwtTokenValid(ctx *gin.Context) {
	//  从请求头部中获取auth
	auth_header := ctx.Request.Header.Get("Authorization")
	if auth_header == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    401,
			"message": "请携带token ",
		})
		ctx.Abort() //中止程序
	} else {
		auths := strings.Split(auth_header, " ")

		bearer := auths[0]
		token := auths[1]

		if len(token) == 0 || len(bearer) == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "请携带正确格式的token ",
			})

			ctx.Abort() //中止程序

		} else {
			user, err := utils.AuthToken(token, []byte(consts.TokenSecret))

			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 401,
					"msg":  "无效的token ",
				})

				ctx.Abort() //中止程序

			} else {
				ctx.Set("username", user.UserName)
				ctx.Next()
			}
		}
	}
}
