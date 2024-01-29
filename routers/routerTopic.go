package routers

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func LoadTopic(e *gin.Engine) {
	TopicController := &controller.TopicController{}
	Topic := e.Group("/api/v1/topic")
	{
		// 创建链接
		Topic.POST("create", TopicController.CreateTopic)
		// 查询链接
		Topic.GET("all", TopicController.GetTopic)
		// 删除链接
		Topic.GET("delete", TopicController.DeleteTopic)
	}
}
