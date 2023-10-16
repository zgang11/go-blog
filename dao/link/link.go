package link

import (
	"blog/database"
	"blog/models"
)

// CreateLink create link name
func CreateLink(link *models.Link) (err error) {
	links := &models.Link{}
	err = database.DB.Debug().Where("url = ?", link.Url).Find(link).Error
	if links.Id != 0 {
		return
	}
	err = database.DB.Debug().Create(&link).Error
	if err != nil {
		return
	}
	return
}
