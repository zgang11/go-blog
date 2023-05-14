package routers

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func LoadTag(e *gin.Engine) {
	TagController := &controller.TagController{}
	Tag := e.Group("/api/v1/tag")
	{
		// 创建文章
		Tag.POST("create", TagController.CreateTag)
		// 后台文章管理查询
		Tag.GET("all", TagController.GetTag)
		// 删除文章
		Tag.DELETE("delete", TagController.DeleteTag)
	}
}
