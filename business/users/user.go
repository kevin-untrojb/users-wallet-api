package users

import (
	"github.com/kevin-untrojb/users-wallet-api/business/wallet"
)

type user struct {
	ID      int64
	Name    string
	Surname string
	Alias   string
	email   string
	Wallets []wallet.Wallet
}
