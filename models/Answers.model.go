// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Answers ..
type Answers struct {
	Title       string `json:"title" binding:"required"`
	QuestionsID uint   `json:"questions_id"`
	Correct     bool   `json:"correct"`
	gorm.Model
}
