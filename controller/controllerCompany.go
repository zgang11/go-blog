package controller

import (
	"blog/dao/company"
	"blog/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CompanyController struct {
}

// CreateCompany 创建公司列表
func (a *CompanyController) CreateCompany(c *gin.Context) {
	companyName := c.PostForm("companyName")
	if companyName == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": "公司名称为空",
		})
		return
	}
	params := &models.Company{CompanyName: companyName, CreateTime: time.Now().UnixNano() / 1e6,
		UpdateTime: time.Now().UnixNano() / 1e6}
	err := company.CreateCompany(params)
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

// GetCompany 获取公司列表
func (a *CompanyController) GetCompany(c *gin.Context) {
	companyList, err := company.GetCompanies()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":        1000,
		"message":     "success",
		"companyList": companyList,
	})
	return
}

// DeleteCompany 删除公司
func (a *CompanyController) DeleteCompany(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": "id为空",
		})
		return
	}
	err := company.DeleteCompany(id)
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
