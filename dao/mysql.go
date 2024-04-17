package dao

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

func InitDB() (err error) {
	dsn := "root:254137.w@(127.0.0.1:3306)/checklist?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return DB.DB().Ping()
}
