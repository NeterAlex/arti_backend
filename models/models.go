package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db, tErr = gorm.Open(sqlite.Open("database/arti_database.db"), &gorm.Config{})

func InitDB() {
	if tErr != nil {
		panic("failed to connect sqlite database")
	}
	_ = db.AutoMigrate(&Auth{})
	_ = db.AutoMigrate(&Article{})
}
