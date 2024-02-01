package controller

import (
	"blog/dao/link"
	"blog/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LinkController struct {
}

var linkLogic = logic.LinkLogic{}

// CreateLink 创建连接列表
func (a *LinkController) CreateLink(c *gin.Context) {
	err := linkLogic.CreateLinkLogic(c)
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
