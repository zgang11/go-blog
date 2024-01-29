package controller

import (
	"blog/dao/topic"
	"blog/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type TopicController struct {
}

// CreateTopic 创建连接列表
func (a *TopicController) CreateTopic(c *gin.Context) {
	topicName := c.PostForm("topicName")
	url := c.PostForm("url")
	serialNumber := c.PostForm("serialNumber")

	params := &models.Topic{TopicName: topicName, Url: url, SerialNumber: serialNumber, CreateTime: time.Now().UnixNano() / 1e6,
		UpdateTime: time.Now().UnixNano() / 1e6}
	err := topic.CreateTopic(params)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1000,
		"message": "success",
	})
}

// GetTopic 获取连接列表
func (a *TopicController) GetTopic(c *gin.Context) {
	topicList, err := topic.GetTopic()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     1000,
		"message":  "success",
		"linkList": topicList,
	})
	return
}

func (a *TopicController) DeleteTopic(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": "id为空",
		})
		return
	}
	err := topic.DeleteTopic(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1000,
		"message": "success",
	})
	return
}
