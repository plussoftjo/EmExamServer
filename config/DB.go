// Package config ...
package config

import (
	"github.com/jinzhu/gorm"
	// Connect mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// models
	"server/models"
)

// SetupDB ...

// DB ..
var DB *gorm.DB

// SetupDB ..
func SetupDB() {
	database, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3307)/questions?charset=utf8mb4&parseTime=True&loc=Local")

	// If Error in Connect
	if err != nil {
		panic(err)
	}
	// User Models Setup
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.AuthClients{})
	database.AutoMigrate(&models.AuthTokens{})
	database.AutoMigrate(&models.Categories{})
	database.AutoMigrate(&models.Questions{})
	database.AutoMigrate(&models.Answers{})
	database.AutoMigrate(&models.UserResults{})
	database.AutoMigrate(&models.NotificationsToken{})

	// Confirm the DB variables
	DB = database

}
