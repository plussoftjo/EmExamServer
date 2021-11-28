// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// TypeCheck ..
type TypeCheck struct {
	Name          string `json:"name"`
	ExamType      string `json:"examType"`
	ExamTime      string `json:"examTime"`
	Success       string `json:"success"`
	Errors        string `json:"errors"`
	Words         string `json:"words"`
	WordPerSecond string `json:"wordPerSecond"`
	Country       string `json:"country"`
	gorm.Model
}
