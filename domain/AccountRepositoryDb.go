package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/xvbnm48/go-microservice-udemy/errs"
	"github.com/xvbnm48/go-microservice-udemy/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type,amount, status) values(? , ? , ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error("Error while inserting account " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	a.AccoutnId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepository(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}
