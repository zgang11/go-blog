package routers

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func LoadTopicTag(e *gin.Engine) {
	TopicTagController := &controller.TopicTagController{}
	TopicTag := e.Group("/api/v1/topicTag")
	{
		// 创建
		TopicTag.POST("create", TopicTagController.CreateTopicTag)
		// 查询
		TopicTag.GET("all", TopicTagController.GetTopicTag)
		// 删除
		TopicTag.GET("delete", TopicTagController.DeleteTopicTag)
		// 更新
		TopicTag.POST("update", TopicTagController.UpdateTopicTag)
	}
}
