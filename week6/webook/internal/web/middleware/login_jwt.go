package middleware

import (
	ijwt "basic-go/webook/internal/web/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type LoginJWTMiddlewareBuilder struct {
	ijwt.Handler
}

func NewLoginJWTMiddlewareBuilder(hdl ijwt.Handler) *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{
		Handler: hdl,
	}
}
func (m *LoginJWTMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == "/users/signup" ||
			path == "/users/login" ||
			path == "/users/login_sms/code/send" ||
			path == "/users/login_sms" ||
			path == "/oauth2/wechat/authurl" ||
			path == "/oauth2/wechat/callback" {
			// 不需要登录校验
			return
		}
		// 根据约定，token 在 Authorization 头部
		// Bearer XXXX
		tokenStr := m.ExtractToken(ctx)
		var uc ijwt.UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
			return ijwt.JWTKey, nil
		})
		if err != nil {
			// token 不对，token 是伪造的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token == nil || !token.Valid {
			// token 解析出来了，但是 token 可能是非法的，或者过期了的
			// 在这里发现access_token过期了，自动刷新
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//这里看Redis
		err = m.CheckSession(ctx, uc.Ssid)
		if err != nil {
			// token无效或者redis有问题
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("user", uc)
	}
}
