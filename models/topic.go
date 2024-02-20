package models

import "github.com/jinzhu/gorm"

type Topic struct {
	gorm.Model
	//Id           int     `json:"id" gorm:"not null unique pk INT(11) 'id'"`
	TopicName    string      `json:"topic_name" gorm:"default 'NULL' CHAR(20) 'topic_name'"`
	Url          string      `json:"url" gorm:"default 'NULL' CHAR(20) 'url'"`
	SerialNumber string      `json:"serial_number" gorm:"default 'NULL' CHAR(20) 'serial_number'"`
	Links        []*Link     `gorm:"many2many:link_topic;"`
	Tags         []*TopicTag `gorm:"many2many:topic_tag;"`
	CreateTime   int64       `json:"create_time" gorm:"default 'NULL' BIGINT 'create_time'"`
	UpdateTime   int64       `json:"update_time" gorm:"default 'NULL' BIGINT 'update_time'"`
}
