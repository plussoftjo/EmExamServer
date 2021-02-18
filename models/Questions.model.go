// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Questions ..
type Questions struct {
	Title        string     `json:"title" binding:"required"`
	CategoriesID uint       `json:"categories_id"`
	Answers      []Answers  `json:"answers" gorm:"foreignKey:QuestionsID;references:ID"`
	Categories   Categories `json:"categories" gorm:"foreignKey:CategoriesID;references:ID"`
	gorm.Model
}
