package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email string
	CanDelete bool
}