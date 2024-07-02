package web

import "github.com/gin-gonic/gin"

type ArticleHandle struct {
}

func (h *ArticleHandle) RegisterRoutes(server *gin.Engine) {
	g := server.Group("/article")
	g.POST("edit", h.Edit)
}

// 接收Article输入 返回一个id 文字的id
func (h *ArticleHandle) Edit(context *gin.Context) {

}
