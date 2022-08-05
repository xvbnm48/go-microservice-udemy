package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xvbnm48/go-microservice-udemy/errs"
	"github.com/xvbnm48/go-microservice-udemy/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	// var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name , city, zipcode, date_of_birth, status  from customers"
		err = d.client.Select(&customers, findAllSql)
		// rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := "select customer_id, name , city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
		// rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// customers := make([]Customer, 0)
	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error while scanning customer table database " + err.Error())
	// 	return nil, errs.NewUnexpectedError("unexpected database error")
	// }

	return customers, nil

}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name , date_of_birth,city, zipcode, status from customers where customer_id = ?"

	// row := d.client.QueryRow(customerSql, id)

	var c Customer
	// err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateofBitrh, &c.Status)
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer" + err.Error())
			return nil, errs.NewNotFoundError("unexpected database error")

		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {

	return CustomerRepositoryDb{
		client: dbClient,
	}
}
