// Package controllers ...
package controllers

import (
	"net/http"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// StoreUserResults ...
func StoreUserResults(c *gin.Context) {
	var userResults models.UserResults
	if err := c.ShouldBindJSON(&userResults); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&userResults).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message":     "Success",
		"userResults": userResults,
	})
}

// StoreUserLogs ..
func StoreUserLogs(c *gin.Context) {
	var data models.UserLogs

	c.ShouldBindJSON(&data)

	err := config.DB.Create(&data).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
	})
}

// IndexUserLogs ..
func IndexUserLogs(c *gin.Context) {
	var data []models.UserLogs
	ID := c.Param("id")

	err := config.DB.Where("user_id = ?", ID).Find(&data).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	c.JSON(200, data)

}
