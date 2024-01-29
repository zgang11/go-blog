package company

import (
	"blog/database"
	"blog/models"
	"strconv"
)

// CreateLink create link name
func CreateCompany(company *models.Company) (err error) {
	company_ := &models.Company{}
	err = database.DB.Debug().Where("url = ?", company.CompanyName).Find(company).Error
	if company_.Id != 0 {
		return
	}
	err = database.DB.Debug().Create(&company).Error
	if err != nil {
		return
	}
	return
}

func GetCompanies() (companyList []models.Company, err error) {
	if err = database.DB.Find(&companyList).Error; err != nil {
		return
	}
	return
}

func DeleteCompany(companyId string) (err error) {
	var id int
	id, err = strconv.Atoi(companyId)
	if err != nil {
		return
	}
	err = database.DB.Debug().Where("id = ?", id).Delete(&models.Company{}).Error
	if err != nil {
		return
	}
	return
}
