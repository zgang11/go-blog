package topic

import (
	"blog/database"
	"blog/models"
	"strconv"
)

// CreateTopic create link name
func CreateTopic(topic *models.Topic) (err error) {
	topics := &models.Topic{}
	err = database.DB.Debug().Where("url = ?", topic.Url).Find(&topics).Error
	if topics.Url != "" {
		return
	}
	err = database.DB.Debug().Create(&topic).Error
	if err != nil {
		return
	}
	return
}

func GetTopic() (topicList []models.Topic, err error) {
	if err = database.DB.Find(&topicList).Error; err != nil {
		return
	}
	return
}

func GetTopicQuery(keyword string) (topicList []models.Topic, err error) {
	serialNumberStr, topicNameStr := keyword, "%"+keyword+"%"
	//if err = database.DB.Where("serialNumber = ? or topicName like ?", serialNumberStr, topicNameStr).Error; err != nil {
	//	return
	//}
	if err = database.DB.Where("serial_number = ?", serialNumberStr).Or("topic_name LIKE ?", topicNameStr).Find(&topicList).Error; err != nil {
		return
	}
	return
}

func DeleteTopic(topicId string) (err error) {
	var id int
	id, err = strconv.Atoi(topicId)
	if err != nil {
		return
	}
	err = database.DB.Debug().Where("id = ?", id).Delete(&models.Topic{}).Error
	if err != nil {
		return
	}
	return
}

func GetTopicByIds(ids []int) (topics []*models.Topic) {
	database.DB.Find(&topics, ids)
	return
}
