package models

import "github.com/jinzhu/gorm"

// 题目标签列表
type TopicTag struct {
	gorm.Model
	Id           int      `json:"id" gorm:"not null unique pk INT(11) 'id'"`
	TopicTagName string   `json:"topic_tag_name" gorm:"default 'NULL' CHAR(20) 'topic_tag_name'"`
	Topics       []*Topic `gorm:"many2many:topic_tag;"`
	CreateTime   int64    `json:"create_time" gorm:"default 'NULL' BIGINT 'create_time'"`
	UpdateTime   int64    `json:"update_time" gorm:"default 'NULL' BIGINT 'update_time'"`
}
