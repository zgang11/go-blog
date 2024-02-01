package link

import (
	"blog/database"
	"blog/models"
	"strconv"
)

// CreateLink create link name
func CreateLink(link *models.Link) (err error) {
	links := &models.Link{}
	err = database.DB.Debug().Where("url = ?", link.Url).Find(&links).Error
	if links.Url != "" {
		return
	}
	err = database.DB.Debug().Create(&link).Error
	if err != nil {
		return
	}
	return
}

func GetLinks() (linkList []models.Link, err error) {
	if err = database.DB.Preload("Topics").Find(&linkList).Error; err != nil {
		return
	}
	return
}

func DeleteLink(linkId string) (err error) {
	var id int
	id, err = strconv.Atoi(linkId)
	if err != nil {
		return
	}
	err = database.DB.Debug().Where("id = ?", id).Delete(&models.Link{}).Error
	if err != nil {
		return
	}
	return
}
