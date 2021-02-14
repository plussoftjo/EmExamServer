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
