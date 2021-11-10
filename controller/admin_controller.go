package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"user/config"
	"user/model"
)

func Login(c *gin.Context) {
	var givenAdmin model.Admin
	var foundAdmin model.Admin
	if err := c.ShouldBindJSON(&givenAdmin); err != nil {
		c.JSON(400, givenAdmin)
	} else {
		if err := config.DB.Where("username = ?", givenAdmin.Username).First(&foundAdmin).Error; err != nil {
			c.Status(401)
		} else {
			if bcrypt.CompareHashAndPassword([]byte(foundAdmin.Password), []byte(givenAdmin.Password)) == nil {
				c.JSON(http.StatusOK, gin.H{
					"token": config.GenerateToken(foundAdmin.Username, foundAdmin.Email,foundAdmin.CanDelete),
				})
			} else {
				c.Status(401)
			}
		}
	}
}
