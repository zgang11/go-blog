package logic

import (
	"blog/dao/topicTag"
	"blog/models"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type TopicTagLogic struct {
}

func (*TopicTagLogic) CreateTopicTagLogic(c *gin.Context) (err error, message string) {
	topicTagName := c.PostForm("topicTagName")
	if topicTagName == "" {
		message = "参数为空"
		err = errors.New(message)
		return
	}
	parma := models.TopicTag{TopicTagName: topicTagName, UpdateTime: time.Now().UnixNano() / 1e6, CreateTime: time.Now().UnixNano() / 1e6}
	err, message = topicTag.CreateTopicTag(&parma)
	return
}

func (*TopicTagLogic) GetTopicTagLogic(c *gin.Context) (res []models.TopicTag, err error) {
	res, err = topicTag.GetTopicTag()
	return
}

func (*TopicTagLogic) DeleteTopicTagLogic(c *gin.Context) (err error) {
	id := c.Query("id")
	err = topicTag.DeleteTopicTag(id)
	return
}

func (*TopicTagLogic) UpdateTopicTagLogic(c *gin.Context) (err error) {
	topicTagId := c.PostForm("id")
	id, err := strconv.Atoi(topicTagId)
	if err != nil {
		return
	}
	topicTagName := c.PostForm("topicTagName")
	topicTagObj := &models.TopicTag{Id: id, TopicTagName: topicTagName, CreateTime: time.Now().UnixNano() / 1e6,
		UpdateTime: time.Now().UnixNano() / 1e6}
	err = topicTag.UpdateTopicTagDao(topicTagObj)
	return
}
