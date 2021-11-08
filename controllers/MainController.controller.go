// Package controllers ...
package controllers

import (
	"fmt"
	"net/http"
	"server/config"
	"server/models"
	"server/vendors"

	"github.com/gin-gonic/gin"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

// Indexquestions ...
func Indexquestions(c *gin.Context) {
	categoriesID := c.Param("categories_id")
	var questions []models.Questions

	config.DB.
		Where("categories_id = ?", categoriesID).
		Preload("Answers").
		Find(&questions)

	// StoreQuestionLogs
	startOfToday, endOfToday := vendors.BetwenToday()
	var todayCount int64
	config.DB.Model(&models.ExamLogs{}).
		Where("created_at BETWEEN ? AND ?", startOfToday, endOfToday).
		Count(&todayCount)

	if todayCount == 0 {
		config.DB.Create(&models.ExamLogs{
			Number: 1,
		})
	} else {
		var todayExamLogs models.ExamLogs
		config.DB.
			Where("created_at BETWEEN ? AND ?", startOfToday, endOfToday).
			First(&todayExamLogs)
		examNumber := todayExamLogs.Number
		examNumber = examNumber + 1
		config.DB.Model(&models.ExamLogs{}).
			Where("id = ?", todayExamLogs.ID).
			Update("number", examNumber)
	}

	c.JSON(http.StatusOK, questions)
}

// Index ..
func Index(c *gin.Context) {

	var categories []models.Categories

	config.DB.Find(&categories)

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})

}

// Indexx ..
func Indexx(c *gin.Context) {

	var categories []models.Categories

	config.DB.Find(&categories)

	c.JSON(http.StatusOK, categories)

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

	config.DB.Where("token = ?", notificationsToken.Token).FirstOrCreate(&notificationsToken)

	c.JSON(http.StatusOK, notificationsToken)
}

// ToggleNotification ...
func ToggleNotification(c *gin.Context) {
	type ToggleNotificationType struct {
		Token  string `json:"token"`
		Active bool   `json:"active"`
	}
	var data ToggleNotificationType
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var notificationToken models.NotificationsToken
	err := config.DB.Model(&models.NotificationsToken{}).Where("token = ?", data.Token).First(&notificationToken).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	if notificationToken.Active == true {
		notificationToken.Active = false
	} else {
		notificationToken.Active = true
	}

	config.DB.Save(&notificationToken)

	c.JSON(200, gin.H{"message": "success"})
}

// IndexAllQuestions ..
func IndexAllQuestions(c *gin.Context) {
	var questions []models.Questions
	var categories []models.Categories

	config.DB.Find(&categories)
	config.DB.Preload("Answers").Preload("Categories").Find(&questions)

	c.JSON(200, gin.H{
		"questions":  questions,
		"categories": categories,
	})
}

// NotificationBody ..
type NotificationBody struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// SendNotificationForAll ..
func SendNotificationForAll(c *gin.Context) {
	var notificationBody NotificationBody
	if err := c.ShouldBindJSON(&notificationBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var NotificationsToken []models.NotificationsToken
	config.DB.Find(&NotificationsToken)

	for _, notification := range NotificationsToken {
		var tokens []expo.ExponentPushToken
		pushToken, err := expo.NewExponentPushToken(notification.Token)
		if err != nil {
			fmt.Println("Error One")
		}

		tokens = append(tokens, pushToken)

		// Create a new Expo SDK client
		client := expo.NewPushClient(nil)

		// Publish message
		response, err := client.Publish(
			&expo.PushMessage{
				To:       tokens,
				Body:     notificationBody.Body,
				Data:     map[string]string{"date": "notification"},
				Sound:    "default",
				Title:    notificationBody.Title,
				Priority: expo.DefaultPriority,
			},
		)
		// Check errors
		if err != nil {
			fmt.Println("error")
		}
		// Validate responses
		if response.ValidateResponse() != nil {

		}
	}

	c.JSON(200, gin.H{
		"message": "Success",
	})
}

// StoreCallUs ..
func StoreCallUs(c *gin.Context) {
	var data models.CallUs

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
		"message": "success",
	})
}
