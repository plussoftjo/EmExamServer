// Package config ...
package config

import (
	"github.com/jinzhu/gorm"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"

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
	database, err := gorm.Open("mysql", "root:00962s00962S!@tcp(127.0.0.1:3306)/emexam?charset=utf8mb4&parseTime=True&loc=Local")

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
	database.AutoMigrate(&models.ExamLogs{})
	database.AutoMigrate(&models.UserLogs{})
	database.AutoMigrate(&models.CallUs{})
	database.AutoMigrate(&models.TypeCheck{})

	converter := typescriptify.New().
		Add(&models.User{}).
		Add(&models.Questions{})
	converter.ConvertToFile("ts/models.ts")

	// Confirm the DB variables
	DB = database

}
