// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// NotificationsToken ..
type NotificationsToken struct {
	Token  string `json:"token" binding:"required"`
	userID uint   `json:"user_id"`
	Active bool   `json:"user_id" gorm:"default:1"`
	gorm.Model
}
