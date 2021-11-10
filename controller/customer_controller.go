package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"user/config"
	"user/model"
)

func FindAllCustomers(c *gin.Context) {
	var customers []model.Customer
	if err := config.DB.Find(&customers).Error; err != nil {
		c.JSON(404, customers)
	} else {
		c.JSON(200, customers)
	}
}

func FindCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var customer model.Customer
	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(404, customer)
	} else {
		c.JSON(200, customer)
	}
}

func AddCustomer(c *gin.Context) {
	var customer model.Customer

	if e := c.ShouldBindJSON(&customer); e != nil {
		c.AbortWithStatusJSON(400,e.Error())
		return
	}

	adviceChan := make(chan string)
	go getAdvice(adviceChan)
	customer.Advice = <- adviceChan
	if err := config.DB.Create(&customer).Error; err == nil {
		c.JSON(200, customer)
	}

}

func UpdateCustomer(c *gin.Context) {
	var customer model.Customer
	if e := c.BindJSON(&customer); e != nil {
		c.AbortWithStatusJSON(400, e.Error())
		return
	}

	if err := config.DB.Save(&customer).Error; err == nil {
		c.JSON(200, customer)
	}


}

func DeleteCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var customer model.Customer
	if err := config.DB.Where("id = ?", id).Delete(&customer).Error; err != nil {
		c.Status(404)
	} else {
		c.Status(200)
	}
}

func getAdvice(c chan string) {
	get, _ := http.Get("https://api.adviceslip.com/advice")

	var advice model.Advice
	body, _ := ioutil.ReadAll(get.Body)
	fmt.Println(body)
	err := json.Unmarshal(body, &advice)
	if err != nil {
		fmt.Println(err)
		c <- ""
	}
	c <- advice.Slip.Advice
}
