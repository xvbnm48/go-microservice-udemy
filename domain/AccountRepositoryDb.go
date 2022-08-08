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

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount FROM accounts WHERE account_id = ?"

	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while getting account " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return &account, nil
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting transaction " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// inserting bank account transaction
	result, _ := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) 
	values (?, ?, ?, ?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	//updating account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}

	// in case or error rollback, and changes from both the tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error("Error while updating account balance " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		logger.Error("Error while commiting transaction for bank account " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// getting the last transaction ID from transaction table
	transactionid, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new transaction" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// getting the latest account information from the accounts table
	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}

	t.TransactionId = strconv.FormatInt(transactionid, 10)

	// updating the transaction struct with the latest ballance
	t.Amount = account.Amount
	return &t, nil
}
