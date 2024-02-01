package models

import "github.com/jinzhu/gorm"

// 连接表
type Link struct {
	gorm.Model
	Id         int      `json:"id" gorm:"not null unique pk INT(11) 'id'"`
	LinkName   string   `json:"link_name" gorm:"default 'NULL' CHAR(20) 'link_name'"`
	Url        string   `json:"url" gorm:"default 'NULL' CHAR(20) 'url'"`
	CompanyId  int      `json:"company_id" gorm:"not null INT(11) 'company_id'"`
	Type       int      `json:"type" gorm:"not null comment('状态 0:全部  1:面经  2:题目')  'type'"`
	CreateTime int64    `json:"create_time" gorm:"default 'NULL' BIGINT 'create_time'"`
	UpdateTime int64    `json:"update_time" gorm:"default 'NULL' BIGINT 'update_time'"`
	Topics     []*Topic `gorm:"many2many:link_topic;"`
}

type Topic struct {
	gorm.Model
	//Id           int     `json:"id" gorm:"not null unique pk INT(11) 'id'"`
	TopicName    string  `json:"topic_name" gorm:"default 'NULL' CHAR(20) 'topic_name'"`
	Url          string  `json:"url" gorm:"default 'NULL' CHAR(20) 'url'"`
	SerialNumber string  `json:"serial_number" gorm:"default 'NULL' CHAR(20) 'serial_number'"`
	Links        []*Link `gorm:"many2many:link_topic;"`
	CreateTime   int64   `json:"create_time" gorm:"default 'NULL' BIGINT 'create_time'"`
	UpdateTime   int64   `json:"update_time" gorm:"default 'NULL' BIGINT 'update_time'"`
}
