package routers

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func LoadCompany(e *gin.Engine) {
	CompanyController := &controller.CompanyController{}
	Article := e.Group("/api/v1/company")
	{
		// 创建链接
		Article.POST("create", CompanyController.CreateCompany)
		// 查询链接
		Article.GET("all", CompanyController.GetCompany)
		// 删除链接
		Article.GET("delete", CompanyController.DeleteCompany)
	}
}
