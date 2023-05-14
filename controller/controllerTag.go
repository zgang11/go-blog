package controller

import (
	"blog/dao/tags"
	"blog/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TagController struct {
}

// 查询文章
func (a *TagController) CreateTag(c *gin.Context) {
	TagName := c.PostForm("tagName")
	// 调用添加标签函数
	tag := &models.Tag{
		TagName:    TagName,
		CreateTime: time.Now().UnixNano() / 1e6,
		UpdateTime: time.Now().UnixNano() / 1e6,
	}

	err := tags.CreateTag(tag)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			// 失败返回code 1001
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

func (a *TagController) GetTag(c *gin.Context) {
	articleList, err := tags.QueryAllTagList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			// 失败返回code 1001
			"code":    1001,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":        1000,
		"message":     "success",
		"articleList": articleList,
	})
	return
}

func (a *TagController) DeleteTag(c *gin.Context) {
	tagId := c.Query("tagId")
	if tagId == "" {
		c.JSON(http.StatusOK, gin.H{
			// 失败返回code 1001
			"code":    1001,
			"message": "未选择id",
		})
		return
	}
	err := tags.DeleteTag(tagId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			// 失败返回code 1001
			"code":    1001,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		// 失败返回code 1001
		"code":    1000,
		"message": "success",
	})
	return
}
