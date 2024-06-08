package middleware

import (
	"basic-go/webook/internel/web"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
)

type LoginJWTMiddlewareBuilder struct {
}

func (m *LoginJWTMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == "/users/signup" || path == "/users/login" {
			//根据约定 token 在Authorization头部
			return
		}
		authCode := ctx.GetHeader("Authorization")
		if authCode == "" {
			//没登陆 在Authorization 内容是乱传的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		segs := strings.Split(authCode, " ")
		if len(segs) != 2 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenstr := segs[1]
		var uc web.UserClaims
		token, err := jwt.ParseWithClaims(tokenstr, &uc, func(token *jwt.Token) (interface{}, error) {
			return web.JWTKey, nil
		})
		if err != nil {
			//token 不对或者伪造的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token == nil || !token.Valid {
			//token 解析出来了但是过期了 也可能是非法的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if uc.UserAgent != ctx.GetHeader("User-Agent") {
			// 监控告警的时候要埋点
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		expireTime := uc.ExpiresAt
		//剩余过期时间 <50 刷新
		//if expireTime.Before(time.Now()) {
		//	ctx.AbortWithStatus(http.StatusUnauthorized)
		//	return
		//}
		if expireTime.Sub(time.Now()) < time.Second*50 {
			uc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute))
			tokenstr, err = token.SignedString(web.JWTKey)
			ctx.Header("x-jwt-token", tokenstr)
			if err != nil {
				// 不能中断
				log.Println(err)
			}
		}
		ctx.Set("user", uc)
	}

}
