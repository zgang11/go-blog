package topic

import (
	"blog/database"
	"blog/models"
	"strconv"
)

// CreateTopic create link name
func CreateTopic(topic *models.Topic) (message string, err error) {
	topics := &models.Topic{}
	err = database.DB.Debug().Where("serial_number = ?", topic.SerialNumber).Or("topic_name = ?", topic.TopicName).Find(&topics).Error
	if topics.ID != 0 {
		message = "添加重复"
		return
	}
	err = database.DB.Debug().Create(&topic).Error
	if err != nil {
		message = "插入数据库失败"
		return
	}
	return
}

func GetTopic() (topicList []models.Topic, err error) {
	if err = database.DB.Preload("Links").Find(&topicList).Error; err != nil {
		return
	}
	return
}

func GetTopicQuery(keyword string) (topicList []models.Topic, err error) {
	serialNumberStr, topicNameStr := keyword, "%"+keyword+"%"
	//if err = database.DB.Where("serialNumber = ? or topicName like ?", serialNumberStr, topicNameStr).Error; err != nil {
	//	return
	//}
	if err = database.DB.Where("serial_number = ?", serialNumberStr).Or("topic_name LIKE ?", topicNameStr).Preload("Links").Find(&topicList).Error; err != nil {
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
