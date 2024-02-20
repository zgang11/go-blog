package logic

import (
	"blog/dao/topic"
	"blog/dao/topicTag"
	"blog/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

type TopicLogic struct {
}

func (this *TopicLogic) CreateTopicLogic(c *gin.Context) (message string, err error) {
	topicName := c.PostForm("topicName")
	url := c.PostForm("url")
	serialNumber := c.PostForm("serialNumber")

	tagIdsStr := c.PostForm("tagIds")
	ids := []int{}
	tagIds := strings.Split(tagIdsStr, ",")
	for _, tagId := range tagIds {
		id, err := strconv.Atoi(tagId)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	tags := topicTag.GetTopicTagByIds(ids)

	params := &models.Topic{TopicName: topicName, Url: url, SerialNumber: serialNumber, Tags: tags, CreateTime: time.Now().UnixNano() / 1e6,
		UpdateTime: time.Now().UnixNano() / 1e6}
	message, err = topic.CreateTopic(params)
	return
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
