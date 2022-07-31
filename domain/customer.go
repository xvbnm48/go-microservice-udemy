package domain

import "github.com/xvbnm48/go-microservice-udemy/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	ZipCode     string
	DateofBitrh string
	Status      string
}

type CustomerRepository interface {
	// status == 1 status == 0 status == ""
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
