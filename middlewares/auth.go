package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-template/controller"
	"go-template/pkg/e"
	"go-template/pkg/jwt"
	"go-template/setting"
	"strings"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	// 初始化 secret
	jwt.NewJWT()
	return func(c *gin.Context) {
		path := c.FullPath()
		if !strings.Contains(path, "signup") {
			authHeader := c.Request.Header.Get(setting.Conf.JwtConfig.AuthHeader)
			if authHeader == "" {
				controller.ResponseError(c, e.CodeNeedLogin)
				c.Abort()
				return
			}
			// 按空格分割
			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				controller.ResponseError(c, e.CodeInvalidToken)
				c.Abort()
				return
			}
			// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
			mc, err := jwt.ParseToken(parts[1])
			if err != nil {
				controller.ResponseError(c, e.CodeInvalidToken)
				c.Abort()
				return
			}
			// 将当前请求的userID信息保存到请求的上下文c上
			c.Set(controller.CtxUserIDKey, mc.UserID)
		}
		c.Next() // 后续的处理请求的函数中 可以用过c.Get(CtxUserIDKey) 来获取当前请求的用户信息
	}
}
