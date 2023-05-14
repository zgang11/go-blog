package tags

import (
	"blog/database"
	"blog/models"
)

// 文章管理列表
type TagList struct {
	Id         int    `json:"id"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	TagName    string `json:"tag_name"`
	//models.Tag
	//models.Article
}

// 添加标签
func CreateTag(tag *models.Tag) (err error) {
	err = database.DB.Debug().Create(&tag).Error
	if err != nil {
		return
	}
	return
}

// 根据名字查询标签
func QueryAllTagList() (tagList []TagList, err error) {
	//db := database.DB.Debug().Where("tag_name = ?",tagName).Last(&tag)
	list := make([]models.Tag, 0)
	err = database.DB.Debug().Find(&list).Error
	if err != nil {
		return
	}
	for _, item := range list {
		tagList = append(tagList, TagList{Id: item.Id, TagName: item.TagName, CreateTime: item.CreateTime, UpdateTime: item.UpdateTime})
	}
	return
}

type Tags struct {
	TagName string
}

// 文章id 对应的标签
func QueryTags(articleIds []int) (map[int][]string, error) {
	tagMap := make(map[int][]string)
	for _, id := range articleIds {
		t := make([]Tags, 0)
		name := make([]string, 0)
		if err := database.DB.Debug().Select("tags.tag_name").Table("article_tags").Joins("left join tags on article_tags.tag_id = tags.id").
			Where("article_tags.article_id = ?", id).Find(&t).Error; err != nil {
			return tagMap, err
		}

		for _, v := range t {
			name = append(name, v.TagName)
		}

		tagMap[id] = name
	}

	return tagMap, nil
}

// 删除标签
func DeleteTag(tagId string) (err error) {
	err = database.DB.Debug().Delete(&models.Tag{}, tagId).Error
	if err != nil {
		return
	}
	return
}
