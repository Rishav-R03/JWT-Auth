package models

import "gorm.io/gorm"

// ✅ Define User model with GORM fields
type User struct {
	gorm.Model
	Username string
	Email    string `gorm:"unique"`
	Password string
}
