package models

type Customer struct {
	Id string `json: "id"`
	Name string `json: "name"`
	Address string `json: "address"`
}

type Vehicle struct {
	ChessisNo string `json: "chessisNo"`
	regNO string `json: "regNo"`
	MakeYear string `json: "makeYear"`
	TypeId string `json: "typeId"`
}

type VehicleType struct {
	TypeId string `json: "string"`
	Name string `json: "name"`
	Company string `json: "company"`
}

type VehcileFeatures struct {
	FeatureId string `json: "string"`
	Name string `json:"name"`
	RegNo string `json: "regNo"`
}

type Driver struct {
	LicenseNo string `json: "licenseNo"`
	Name string `json:"name"`
	age string `json: "age"`
	Address string `json:"address"`
}

type AllowedDestination struct {
	Id string `json:"allowed_dest_id"`
	Name string `json: "name"`
	Distance string `json: "distance"`
}

type Destination struct {
	DestinationId string `json:"destinationId"`
	NoOfPassengers int `json:"noOfPassengers"`
	RentalCost float64 `json:"rentalCost"`
	AllowedDestId string `json:"allowedDestId"`
	CustomerId string `json: "customerId"`
	DriverId string `json: "driverId"`
}