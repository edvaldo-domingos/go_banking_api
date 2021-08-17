package domain

import "github.com/ashishjuyal/banking-lib/errs"
type Customer struct{
	// maps the name of the struct with the customer props
	Id string `db:"customer_id"`
	Name string
	City string
	Zipcode string
	DateofBirth string `db:"date_of_birth"`
	Status string
}

type CustomerRepository interface{
	FindAll(status string)([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}