package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open(mysql.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	DB = db

	return DB
}
