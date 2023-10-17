package controller

import (
	"blog/dao/link"
	"blog/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LinkController struct {
}

// CreateLink 创建连接列表
func (a *LinkController) CreateLink(c *gin.Context) {
	linkName := c.PostForm("linkName")
	url := c.PostForm("url")
	params := &models.Link{LinkName: linkName, Url: url, CreateTime: time.Now().UnixNano() / 1e6,
		UpdateTime: time.Now().UnixNano() / 1e6}
	err := link.CreateLink(params)
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

// GetLink 获取连接列表
func (a *LinkController) GetLink(c *gin.Context) {
	linkList, err := link.GetLinks()
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
		"linkList": linkList,
	})
	return
}

func (a *LinkController) DeleteLink(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": "id为空",
		})
		return
	}
	err := link.DeleteLink(id)
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
