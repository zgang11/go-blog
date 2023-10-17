package routers

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func LoadLink(e *gin.Engine) {
	LinkController := &controller.LinkController{}
	Article := e.Group("/api/v1/link")
	{
		// 创建链接
		Article.POST("create", LinkController.CreateLink)
		// 查询链接
		Article.GET("all", LinkController.GetLink)
		// 删除链接
		Article.GET("delete", LinkController.DeleteLink)
	}
}
