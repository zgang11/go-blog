package models

// 公司分类表
type Company struct {
	Id          int    `json:"id" gorm:"not null unique pk INT(11) 'id'"`
	CompanyName string `json:"link_name" gorm:"default 'NULL' CHAR(20) 'link_name'"`
	CreateTime  int64  `json:"create_time" gorm:"default 'NULL' BIGINT 'create_time'"`
	UpdateTime  int64  `json:"update_time" gorm:"default 'NULL' BIGINT 'update_time'"`
}
