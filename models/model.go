package models

type Customer struct {
	Name string `json: "name"`
}

func GetData()  {
	var c Customer
	c.Name = "Reshma"
}