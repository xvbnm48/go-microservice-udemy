package domain

import "github.com/xvbnm48/go-microservice-udemy/errs"

type Customer struct {
	Id          string `db:"customer_id" json:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateofBitrh string `db:"date_of_birth" json:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	// status == 1 status == 0 status == ""
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
