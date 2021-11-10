package middleware

import (
	"github.com/gin-gonic/gin"
	"user/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {

		const bearer = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < len(bearer) {
			c.AbortWithStatus(400)
			return
		}
		tokenString := authHeader[len(bearer):]

		if !config.ValidateToken(tokenString) {
			c.AbortWithStatus(401)
			return
		}

	}
}
