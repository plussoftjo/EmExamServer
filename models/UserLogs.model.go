// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// UserLogs ..
type UserLogs struct {
	CategoryTitle string `json:"categoryTitle"`
	UserID        uint   `json:"userID"`
	Mark          string `json:"mark"`
	gorm.Model
}
