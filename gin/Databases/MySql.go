package databases

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	db, err := gorm.Open(mysql.Open("hello"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}
