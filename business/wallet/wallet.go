package wallet

import (
	"fmt"
	"log"
	"time"
)
const (
	extractionType = "extraction"
	depositType = "deposit"
)
type Transaction struct {
	ID              int64     `json:"id,omitempty"`
	WalletID        int64     `json:"wallet_id,omitempty"`
	TransactionType string    `json:"transaction_type"`
	UserID          int64     `json:"user_id,omitempty"`
	Date            time.Time `json:"date_create"`
	Amount          string    `json:"amount"`
	CurrencyName    string    `json:"currency"`
}

type Wallet struct {
	ID             int64         `json:"id"`
	CurrencyName   string        `json:"currency_name"`
	CurrentBalance string        `json:"current_balance"`
	CointExponent  int           `json:"-"`
	Coin           *Coin         `json:"-"`
	Transactions   []Transaction `json:"transactions,omitempty"`
}

func newWallet(CurrentBalance string, CoinExponent int) Wallet {
	newCoin, _ := newCoin(CurrentBalance, CoinExponent)
	return Wallet{
		CurrentBalance: CurrentBalance,
		CointExponent:  CoinExponent,
		Coin:           newCoin,
	}
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
		w.Coin, ok = newCoin(w.CurrentBalance, w.CointExponent)
		if !ok {
			log.Println(fmt.Sprintf("error: converting %s into a number", w.CurrentBalance))
			return fmt.Errorf("error: converting %s into a number", w.CurrentBalance)
		}
	}

	switch transaction.TransactionType {
	case depositType:
		ok = w.Coin.Add(transaction.Amount)
		if !ok {
			log.Println(fmt.Sprintf("error: adding %s into amount %s ",transaction.Amount, w.Coin.GetAmount()))
			return fmt.Errorf("error: adding %s into a %s val", transaction.Amount, w.CurrentBalance)
		}
	case extractionType:
		ok = w.Coin.Sub(transaction.Amount)
		if !ok {
			log.Println(fmt.Sprintf("error: subtract %s from amount %s ",transaction.Amount, w.Coin.GetAmount()))
			return fmt.Errorf("error: substracting %s into a %s val", transaction.Amount, w.CurrentBalance)
		}
		if w.Coin.IsNegative() {
			return fmt.Errorf("error: balance is negative")
		}
	default:
		log.Println(fmt.Sprintf("error: invalid transaction type %s ",transaction.TransactionType))
		return fmt.Errorf("error: invalid transaction type")
	}
	return nil
}

type NewTransactionResponse struct {
	ID int64 `json:"transaction_id"`
}
