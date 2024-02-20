package topicTag

import (
	"blog/database"
	"blog/models"
	"errors"
	"strconv"
)

func CreateTopicTag(topicTag *models.TopicTag) (err error, message string) {
	topicTag_ := &models.TopicTag{}
	err = database.DB.Debug().Where("topic_tag_name = ?", topicTag.TopicTagName).Find(&topicTag_).Error
	if topicTag_.ID != 0 {
		message = "重复添加"
		err = errors.New(message)
		return
	}
	err = database.DB.Debug().Create(&topicTag).Error
	if err != nil {
		message = "添加数据库失败"
		return
	}
	return
}

func GetTopicTag() (topicTagList []models.TopicTag, err error) {
	err = database.DB.Find(&topicTagList).Error
	return
}

func GetTopicTagByIds(ids []int) (tags []*models.TopicTag) {
	database.DB.Find(&tags, ids)
	return
}

func DeleteTopicTag(topicTagId string) (err error) {
	var id int
	id, err = strconv.Atoi(topicTagId)
	if err != nil {
		return
	}
	err = database.DB.Debug().Where("id = ?", id).Delete(&models.TopicTag{}).Error
	if err != nil {
		return
	}
	return
}

func UpdateTopicTagDao(topicTag *models.TopicTag) (err error) {
	err = database.DB.Debug().Model(&models.Article{}).Where("id = ?", topicTag.ID).Update(topicTag).Error
	if err != nil {
		return
	}
	return
}
