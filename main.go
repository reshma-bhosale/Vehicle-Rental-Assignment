package main

import (
	"fmt"
	_"net/http"
	"github.com/gin-gonic/gin"
	"github.com/reshma-bhosale/Vehicle-Rental-Assignment/models"
	"github.com/reshma-bhosale/Vehicle-Rental-Assignment/handlers"
)

func main()  {
	fmt.Println("hii")
	var c models.Customer
	c.Name = "Reshma"
	fmt.Println(c.Name)

	r := gin.Default()
	r.POST("/customers", handler.AddNewCustomer)
	r.GET("/customer", handler.GetCustomers)
	r.Run()
}