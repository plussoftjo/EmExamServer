// Package controllers ...
package controllers

import (
	"net/http"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// StoreTypeCheck ...
func StoreTypeCheck(c *gin.Context) {
	var data models.TypeCheck
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetTypeCheck ...
func GetTypeCheck(c *gin.Context) {
	var data []models.TypeCheck

	if err := config.DB.Take(100).Find(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

// IndexWithFilterTypeCheckType ..
type IndexWithFilterTypeCheckType struct {
	Country  string `json:"country"`
	ExamType string `json:"examType"`
}

// IndexWithFilterTypeCheck ..
func IndexWithFilterTypeCheck(c *gin.Context) {

	var data IndexWithFilterTypeCheckType

	c.ShouldBindJSON(&data)

	var results []models.TypeCheck
	err := config.DB.Take(100).Where("country = ?", data.Country).Where("exam_type = ?", data.ExamType).Find(&results).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	c.JSON(200, results)
}
