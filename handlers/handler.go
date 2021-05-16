package handler 

import (
	"fmt"
	_"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/reshma-bhosale/Vehicle-Rental-Assignment/models"

)

func AddNewCustomer(c *gin.Context)  {
	var customer models.Customer
	c.Bind(&customer)
	c.JSON(200, customer)

}

func GetCustomers(c *gin.Context) {
	// res, err := os.Open("customer.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer res.Close()
	//c.JSON(200, res)
}