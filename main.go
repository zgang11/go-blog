package main

import (
	"blog/database"
	"blog/middleware"
	"blog/models"
	"blog/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 数据库初始化
	err := database.InitMySql()
	if err != nil {
		panic(err)
	}
	// 关闭连接
	defer database.DB.Close()
	// 模型和数据库表映射 创建表
	database.DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Article{}, &models.Tag{}, &models.ArticleTag{}, &models.Link{}, &models.Topic{}, &models.Company{}, &models.TopicTag{})
	// 注册路由
	r := gin.Default()
	// 将中间件注册到全局 对所有路由生效
	r.Use(middleware.Cors())
	// 注册登录路由
	routers.LoadLogin(r)
	//// 文章分类路由
	//routers.LoadCategory(r)
	//// 图片上传
	//routers.ImgUpload(r)
	//// 文章管理
	//routers.LoadArticle(r)
	// 标签
	routers.LoadTag(r)
	// 前端面经链接
	routers.LoadLink(r)
	// 公司分类
	routers.LoadCompany(r)
	// leetcode题目链接
	routers.LoadTopic(r)
	// 题目类别标签
	routers.LoadTopicTag(r)

	r.Run(":8081")

}
