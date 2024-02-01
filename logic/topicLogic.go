package logic

import (
	"blog/dao/topic"
	"blog/models"
	"github.com/gin-gonic/gin"
)

type TopicLogic struct {
}

func (this *TopicLogic) GetTopicLogic(c *gin.Context) (topicList []models.Topic, err error) {
	keyword := c.Query("keyword")
	if keyword == "" {
		topicList, err = topic.GetTopic()
	} else {
		topicList, err = topic.GetTopicQuery(keyword)
	}
	return
}
