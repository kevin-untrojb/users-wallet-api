package wallet

import (
	"time"
)

type Wallet struct {
	ID             int64         `json:"id"`
	CurrencyName   string        `json:"curency_name"`
	CurrentBalance string        `json:"current_balance"`
	Currency       Currency      `json:"-"`
	Transactions   []Transaction `json:"transactions"`
}
type Transaction struct {
	ID              int64     `json:"id"`
	WalletID        string    `json:"-"`
	TransactionType string    `json:"movement_type"`
	UserID          string    `json:"-"`
	Date            time.Time `json:"date_create"`
	Amount          string    `json:"amount"`
	CurrencyName    string    `json:"currency"`
}

func (w Wallet) ToUserWallet() Wallet {
	return Wallet{
		CurrencyName:   w.CurrencyName,
		CurrentBalance: w.Currency.GetAmount(w.CurrentBalance),
	}
}

type NewTransactionResponse struct {
	ID int64 `json:"transaction_id"`
}
