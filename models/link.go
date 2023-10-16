package models

// 连接表
type Link struct {
	Id         int    `json:"id" gorm:"not null unique pk INT(11) 'id'"`
	LinkName   string `json:"link_name" gorm:"default 'NULL' CHAR(20) 'link_name'"`
	Url        string `json:"url" gorm:"default 'NULL' CHAR(20) 'url'"`
	CreateTime int64  `json:"create_time" gorm:"default 'NULL' BIGINT 'create_time'"`
	UpdateTime int64  `json:"update_time" gorm:"default 'NULL' BIGINT 'update_time'"`
}
