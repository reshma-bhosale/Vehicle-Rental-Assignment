package handler 

import (
	"github.com/gin-gonic/gin"
	"github.com/reshma-bhosale/Vehicle-Rental-Assignment/models"

)

func handle(c *gin.Context)  {
	var customer  model.Customer
	c.Bind(&customer)
	c.JSON(200, customer)

}