package router

import (
	"github.com/gin-gonic/gin"
	"user/middleware"
)
import "user/controller"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	customers := api.Group("/customer")
	customers.Use(middleware.AuthMiddleware())
	customers.GET("/", controller.FindAllCustomers)
	customers.GET("/:id", controller.FindCustomer)
	customers.POST("/", controller.AddCustomer)
	customers.PUT("/", controller.UpdateCustomer)
	customers.DELETE("/:id", controller.DeleteCustomer)

	admin := api.Group("/admin")
	admin.POST("/login", controller.Login)
	return r
}
