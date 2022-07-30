package domain

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
}
