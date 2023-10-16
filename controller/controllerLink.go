package controller

import (
	"blog/dao/link"
	"blog/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LinkController struct {
}

// 创建连接列表
func (a *LinkController) CreateLink(c *gin.Context) {
	linkName := c.PostForm("linkName")
	url := c.PostForm("url")
	fmt.Println(linkName, url)
	params := &models.Link{LinkName: linkName, Url: url}
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

// 获取连接列表
func (a *LinkController) GetLink(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    1000,
		"message": "success",
	})
}
