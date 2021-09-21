package wallet

import (
	"fmt"
	"github.com/kevin-untrojb/users-wallet-api/utils"
	"log"
	"time"
)

const (
	extractionType = "extraction"
	depositType    = "deposit"
)

type Transaction struct {
	ID              int64     `json:"id,omitempty"`
	WalletID        int64     `json:"wallet_id,omitempty"`
	TransactionType string    `json:"transaction_type"`
	UserID          int64     `json:"user_id,omitempty"`
	Date            time.Time `json:"date_create,omitempty"`
	Amount          string    `json:"amount,omitempty"`
	CurrencyName    string    `json:"currency,omitempty"`
}

type Wallet struct {
	ID             int64         `json:"id"`
	CurrencyName   string        `json:"currency_name"`
	CurrentBalance string        `json:"current_balance"`
	CoinExponent   int           `json:"-"`
	Coin           *Coin         `json:"-"`
	Transactions   []Transaction `json:"transactions,omitempty"`
}

func (u *Transaction) ValidateFields() error {
	if !utils.IsValidField(u.TransactionType) || !utils.IsValidField(u.Amount) {
		return fmt.Errorf("bad_request: error transaction type  or amount")
	}
	return nil
}
func (w Wallet) ToUserWallet() Wallet {
	return Wallet{
		ID:             w.ID,
		CurrencyName:   w.CurrencyName,
		CurrentBalance: w.Coin.GetAmount(),
	}
}
func (w Wallet) GetCurrentBalance() string {
	return w.Coin.GetAmount()
}

func (w Wallet) TryNewTransaction(transaction Transaction) error {
	var ok bool
	if w.Coin == nil {
		w.Coin, ok = newCoin(w.CurrentBalance, w.CoinExponent)
		if !ok {
			log.Println(fmt.Sprintf("error: converting %s into a number", w.CurrentBalance))
			return fmt.Errorf("error: converting %s into a number", w.CurrentBalance)
		}
	}

	switch transaction.TransactionType {
	case depositType:
		ok = w.Coin.Add(transaction.Amount)
		if !ok {
			log.Println(fmt.Sprintf("error: adding %s into amount %s ", transaction.Amount, w.Coin.GetAmount()))
			return fmt.Errorf("error: adding %s into a %s val", transaction.Amount, w.CurrentBalance)
		}
	case extractionType:
		ok = w.Coin.Sub(transaction.Amount)
		if !ok {
			log.Println(fmt.Sprintf("error: subtract %s from amount %s ", transaction.Amount, w.Coin.GetAmount()))
			return fmt.Errorf("error: substracting %s into a %s val", transaction.Amount, w.CurrentBalance)
		}
		if w.Coin.IsNegative() {
			return fmt.Errorf("error: balance is negative")
		}
	default:
		log.Println(fmt.Sprintf("error: invalid transaction type %s ", transaction.TransactionType))
		return fmt.Errorf("error: invalid transaction type")
	}
	return nil
}