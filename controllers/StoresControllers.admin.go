// Package controllers ...
package controllers

import (
	"net/http"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// StoreQuestions ...
func StoreQuestions(c *gin.Context) {
	var question models.Questions
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message":  "Success",
		"question": question,
	})
}

// StoreCategories ...
func StoreCategories(c *gin.Context) {
	var category models.Categories
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message":  "Success",
		"category": category,
	})
}

// StoreAnswers ...
func StoreAnswers(c *gin.Context) {
	var answer models.Answers
	if err := c.ShouldBindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&answer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message":  "Success",
		"category": answer,
	})
}

// StoreQuestionType ..
type StoreQuestionType struct {
	Question models.Questions `json:"question"`
	Answers  []models.Answers `json:"answers"`
}

// StoreQuestionsWithAnswers ...
func StoreQuestionsWithAnswers(c *gin.Context) {
	var storeQuestionType StoreQuestionType
	if err := c.ShouldBindJSON(&storeQuestionType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question := storeQuestionType.Question
	answers := storeQuestionType.Answers

	config.DB.Create(&question)

	for _, answer := range answers {
		AnswerQuery := models.Answers{
			Title:       answer.Title,
			QuestionsID: question.ID,
			Correct:     answer.Correct,
		}
		config.DB.Create(&AnswerQuery)
	}

	var questionQuery models.Questions
	config.DB.Preload("Answers").Preload("Categories").Where("id = ?", question.ID).First(&questionQuery)

	c.JSON(200, gin.H{
		"question": questionQuery,
	})

}

// RemoveQuestion ..
func RemoveQuestion(c *gin.Context) {
	id := c.Param("id")

	config.DB.Delete(&models.Questions{}, id)
	config.DB.Delete(&models.Answers{}, "questions_id = ?", id)
}
