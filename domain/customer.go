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
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
