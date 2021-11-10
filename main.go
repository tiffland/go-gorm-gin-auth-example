package main

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"user/config"
	"user/model"
	"user/router"
)

func main() {
	r := router.SetupRouter()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"),bcrypt.DefaultCost)
	config.DB, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	config.DB.AutoMigrate(&model.Customer{})
	config.DB.AutoMigrate(&model.Admin{})
	config.DB.Create(&model.Admin{
		Username:  "admin",
		Password:  string(hashedPassword),
		Email:     "admin@company.com",
		CanDelete: false,
	})
	r.Run()
}
