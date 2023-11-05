package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBConnTemplate = "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"
)

var db *gorm.DB

func init() {
	zone, _ := time.Now().Zone()
	db = MustConnect(fmt.Sprintf(DBConnTemplate, "0.0.0.0", "root", "123456", "castanha", "9090", "disable", zone))
}

func GetDB() *gorm.DB {
	return db
}

func MustConnect(conn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic(err)
	}
	return db
}
