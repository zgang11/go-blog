package articleTag

import (
	"blog/database"
	"blog/models"
)

// 创建文章标签记录表
func CreateArticleTag(articleTag *models.ArticleTag) (err error) {
	articleAndTag := &models.ArticleTag{}
	err = database.DB.Debug().Where("article_id = ? AND tag_id = ?", articleTag.ArticleId, articleTag.TagId).Find(articleTag).Error
	if articleAndTag.Id != 0 {
		return
	}
	err = database.DB.Debug().Create(&articleTag).Error
	if err != nil {
		return
	}
	return
}

// 查询文章标签对应关系表
func QueryArticleTag() (ArticleTag []models.ArticleTag, err error) {
	if err = database.DB.Debug().Find(&ArticleTag).Error; err != nil {
		return
	}
	return
}
