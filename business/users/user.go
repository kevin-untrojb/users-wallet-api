package users

import (
	"github.com/kevin-untrojb/users-wallet-api/business/wallet"
)

type user struct {
	ID      int64           `json:"id"`
	Name    string          `json:"name"`
	Surname string          `json:"surname"`
	Alias   string          `json:"alias"`
	email   string          `json:"email"`
	Wallets []wallet.Wallet `json:"wallets"`
}
