// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// UserResults ..
type UserResults struct {
	UserID    uint `json:"user_id"`
	Questions int  `json:"questions"`
	Results   int  `json:"results"`
	gorm.Model
}
