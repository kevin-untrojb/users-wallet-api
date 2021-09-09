package wallet

import (
	"time"
)

type Wallet struct {
	ID             string     `json:"id"`
	UserID         string     `json:"user_id"`
	CurrencyName   string     `json:"curency_name"`
	CurrentBalance string     `json:"current_balance"`
	Currency       Currency   `json:"-"`
	Movements      []Movement `json:"movements"`
}
type Movement struct {
	ID           string    `json:"id"`
	WalletID     string    `json:"-"`
	MovementType string    `json:"movement_type"`
	UserID       string    `json:"-"`
	Date         time.Time `json:"date_create"`
	Amount       string    `json:"amount"`
	CurrencyName string    `json:"currency"`
}
func (w Wallet) ToUserWallet() Wallet{
	return Wallet{
		CurrencyName: w.CurrencyName,
		CurrentBalance: w.Currency.GetAmount(w.CurrentBalance),
	}
}