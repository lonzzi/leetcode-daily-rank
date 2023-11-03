package data

import (
	"github.com/lonzzi/leetcode-daily-rank/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open("data/db/info.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&models.User{},
		&models.UserCalendar{},
		&models.RecentACSubmission{},
		&models.Question{},
	)
	DB = db
}

func Get() *gorm.DB {
	return DB
}
