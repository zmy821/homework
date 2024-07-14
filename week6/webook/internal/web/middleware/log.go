package middleware

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type LogMiddlewareBuilder struct {
	logFn         func(ctx context.Context, l AccessLog)
	allowReqBody  bool
	allowRespBody bool
}

func (l *LogMiddlewareBuilder) AllowReqBody() *LogMiddlewareBuilder {
	l.allowReqBody = true
	return l
}
func (l *LogMiddlewareBuilder) AllowRespBody() *LogMiddlewareBuilder {
	l.allowRespBody = true
	return l
}
func NewLogMiddlewareBuilder(logFn func(ctx context.Context, l AccessLog)) *LogMiddlewareBuilder {
	return &LogMiddlewareBuilder{
		logFn: logFn,
	}
}

func (l *LogMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if len(path) > 1024 {
			path = path[:1024]
		}
		method := ctx.Request.Method
		al := AccessLog{
			Path:   path,
			Method: method,
		}
		if l.allowReqBody {
			body, _ := ctx.GetRawData()
			if len(body) > 2048 {
				al.ReqBody = string(body)
			} else {
				al.ReqBody = string(body)
			}
			// 可以忽略 Request.body 是一个stream对象
			// 放回去
			ctx.Request.Body = io.NopCloser(bytes.NewReader(body))
			//ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}
		start := time.Now()
		if l.allowRespBody {
			ctx.Writer = &responseWriter{
				ResponseWriter: ctx.Writer,
				al:             &al,
			}
		}
		defer func() {
			al.Duration = time.Since(start)
			// duration:=time.Now().Sub(start)
			l.logFn(ctx, al)
		}()
		// 直接执行下一个middleware ... 知道业务逻辑
		ctx.Next()
		// 在这里，就拿到了响应
	}
}

type AccessLog struct {
	Path     string        `json:"path"`
	Method   string        `json:"method"`
	ReqBody  string        `json:"req_body"`
	Status   int           `json:"status"`
	RespBody string        `json:"resp_body"`
	Duration time.Duration `json:"duration"`
}

type responseWriter struct {
	gin.ResponseWriter
	//data []bytes
	al *AccessLog
}

func (w *responseWriter) Write(data []byte) (int, error) {
	//w.data = data
	w.al.RespBody = string(data)
	return w.ResponseWriter.Write(data)
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.al.Status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}