package controller

import (
	"blog/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TopicTagController struct {
}

var TopicTagLogic = logic.TopicTagLogic{}

// CreateTopicTag 创建题目标签
func (a *TopicTagController) CreateTopicTag(c *gin.Context) {
	err, message := TopicTagLogic.CreateTopicTagLogic(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": message,
			"err":     err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1000,
		"message": "添加成功",
	})
}

// GetTopicTag 获取题目标签
func (a *TopicTagController) GetTopicTag(c *gin.Context) {
	list, err := TopicTagLogic.GetTopicTagLogic(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"list": list,
			"err":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1000,

		"companyList": list,
	})
	return
}

// DeleteTopicTag 删除题目标签
func (a *TopicTagController) DeleteTopicTag(c *gin.Context) {
	err := TopicTagLogic.DeleteTopicTagLogic(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": "删除失败",
			"err":     err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1000,
		"message": "删除成功",
	})
	return
}

func (a *TopicTagController) UpdateTopicTag(c *gin.Context) {
	err := TopicTagLogic.UpdateTopicTagLogic(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": "更新失败",
			"err":     err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1000,
		"message": "更新成功",
	})
	return
}
