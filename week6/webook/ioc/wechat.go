package ioc

import (
	"basic-go/webook/internal/service/oauth2/wechat"
	"basic-go/webook/pkg/logger"
	"os"
)

func InitWechatService(l logger.LoggerV1) wechat.Service {
	//appID, ok := os.LookupEnv("WECHAT_APP_ID")
	//if !ok {
	//	panic("找不到环境变量WECHAT_APP_ID")
	//}
	//appSecret, ok := os.LookupEnv("WECHAT_APP_SECRET")
	//if !ok {
	//	panic("找不到环境变量WECHAT_APP_SECRET")
	//
	//}
	appID := os.Getenv("WECHAT_APP_ID")
	appSecret := os.Getenv("WECHAT_APP_SECRET")
	return wechat.NewService(appID, appSecret, l)
}
