// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// CallUs ..
type CallUs struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
	gorm.Model
}
