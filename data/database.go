package data

import (
	"os"

	"github.com/lonzzi/leetcode-daily-rank/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	_, err := os.Stat("data/db")
	if os.IsNotExist(err) {
		os.MkdirAll("data/db", 0755)
	}

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
