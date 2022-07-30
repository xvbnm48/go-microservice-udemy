package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	custoomers := []Customer{
		{Id: "1010", Name: "sakura endo", City: "japan", ZipCode: "2367", DateofBitrh: "2001-08-30", Status: "1"},
		{Id: "101", Name: "sakura miyawaki", City: "japan", ZipCode: "2367", DateofBitrh: "2001-08-30", Status: "1"},
		{Id: "110", Name: "akari kito", City: "japan", ZipCode: "2367", DateofBitrh: "2001-08-30", Status: "1"},
	}

	return CustomerRepositoryStub{customers: custoomers}
}
