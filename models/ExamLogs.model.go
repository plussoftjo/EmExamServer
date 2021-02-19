// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// ExamLogs ..
type ExamLogs struct {
	Number int64 `json:"number"`
	gorm.Model
}
