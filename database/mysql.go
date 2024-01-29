package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

// 初始化连接
func InitMySql() (err error) {
	dsn := "root:shualeecode1314_CFZY@tcp(101.43.135.248:3306)/user?charset=utf8mb4&parseTime=True&loc=Local" //线上
	//dsn := "root:shualeecode1314_CFZY@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local" //线下
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return
	}

	if err := db.Debug().DB().Ping(); err != nil {
		panic(err)
	}

	DB = db
	return
}

func Close() {
	DB.Close()
}
