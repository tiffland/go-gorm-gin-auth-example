package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Advice string
}
