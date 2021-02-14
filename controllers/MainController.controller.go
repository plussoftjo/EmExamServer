// Package controllers ...
package controllers

import (
	"net/http"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// Indexquestions ...
func Indexquestions(c *gin.Context) {
	categoriesID := c.Param("categories_id")
	var questions []models.Questions

	config.DB.
		Where("categories_id = ?", categoriesID).
		Preload("Answers").
		Find(&questions)

	c.JSON(http.StatusOK, gin.H{
		"questions": questions,
	})
}

// Index ..
func Index(c *gin.Context) {

	var categories []models.Categories

	config.DB.Find(&categories)

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})

}

// IndexWithAuth ...
func IndexWithAuth(c *gin.Context) {
	var categories []models.Categories
	var userResults []models.UserResults

	config.DB.Find(&categories)
	userID := c.Param("user_id")
	config.DB.Where("user_id = ?", userID).Find(&userResults)

	c.JSON(http.StatusOK, gin.H{
		"categories":  categories,
		"userResults": userResults,
	})
}

// StoreNotificationToken ...
func StoreNotificationToken(c *gin.Context) {
	var notificationsToken models.NotificationsToken
	if err := c.ShouldBindJSON(&notificationsToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&notificationsToken).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Success",
	})
}
