package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"messageBox/internal/model"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Group{})
	database.AutoMigrate(&model.GroupUser{})
	database.AutoMigrate(&model.Reply{})
	database.AutoMigrate(&model.Message{})

	DB = database
}
