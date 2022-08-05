package domain

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionId   string  `db:"transaction_id" json:"transaction_id"`
	AccountId       string  `db:"account_id" json:"account_id"`
	Amount          float64 `db:"amount" json:"amount"`
	TransactionType string  `db:"transaction_type" json:"transaction_type"`
	TransactionDate string  `db:"transaction_date" json:"transaction_date"`
}

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}
