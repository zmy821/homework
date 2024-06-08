package main

import (
	"basic-go/webook/config"
	"basic-go/webook/internel/repository"
	"basic-go/webook/internel/repository/dao"
	"basic-go/webook/internel/service"
	"basic-go/webook/internel/web"
	"basic-go/webook/internel/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

func main() {
	db := initDB()

	server := initWebServer()
	initUserHdl(db, server)
	//server := gin.Default()
	//server.GET("/hello", func(ctx *gin.Context) {
	//	ctx.String(http.StatusOK, "hello启动成功了")
	//})
	server.Run(":8081")
}

func initUserHdl(db *gorm.DB, server *gin.Engine) {
	ud := dao.NewUserDAO(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)
	hdl := web.NewUserHandler(us)
	hdl.RegisterRoutes(server)
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	if err != nil {
		panic(err)
	}

	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello启动成功了")
	})

	server.Use(cors.New(cors.Config{
		//AllowAllOrigins: true,
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,

		AllowHeaders:  []string{"Authorization", "Content-Type"},
		ExposeHeaders: []string{"X-Jwt-Token"},

		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				//if strings.Contains(origin, "localhost") {
				return true
			}
			return strings.Contains(origin, "your_company.com")
		},
		MaxAge: 12 * time.Hour,
	}), func(ctx *gin.Context) {
		println("这是我的 Middleware")
	})
	_ = redis.NewClient(&redis.Options{
		Addr: config.Config.Redis.Addr,
	})
	//server.Use(ratelimit.NewBuilder(redisCilent,
	//	time.Second, 100).Build())
	useJWT(server)
	//useSession(server)
	return server
}

func useJWT(server *gin.Engine) {
	login := &middleware.LoginJWTMiddlewareBuilder{}
	server.Use(login.CheckLogin())
}

func useSession(server *gin.Engine) {
	login := &middleware.LoginMiddlewareBuilder{}
	// 存储数据的，也就是你 userId 存哪里
	// 直接存 cookie
	store := cookie.NewStore([]byte("secret"))
	//基于内存实现
	//store := memstore.NewStore([]byte("q'/?nR]m0ngm|#6KsXlj&h>LgDya8mwV"), []byte("XWyI)|>XwW;00pqYY,W9&YkSf:\\+\"94i"))
	//store, err := redis.NewStore(16, "tcp",
	//	"localhost:6379",
	//	"",
	//	[]byte("q'/?nR]m0ngm|#6KsXlj&h>LgDya8mwV"),
	//	[]byte("XWyI)|>XwW;00pqYY,W9&YkSf:\\+\"94i"))
	//if err != nil {
	//	panic(err)
	//}
	server.Use(sessions.Sessions("ssid", store), login.CheckLogin())
}
