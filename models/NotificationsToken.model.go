// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// NotificationsToken ..
type NotificationsToken struct {
	Token  string `json:"token" binding:"required"`
	UserID uint   `json:"userID"`
	Active bool   `json:"active" gorm:"default:1"`
	App    string `json:"app"`
	gorm.Model
}
