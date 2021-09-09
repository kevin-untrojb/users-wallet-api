package users

import (
	"fmt"

	"github.com/kevin-untrojb/users-wallet-api/business/wallet"
	"github.com/kevin-untrojb/users-wallet-api/utils"
)

type user struct {
	ID      int64           `json:"id"`
	Name    string          `json:"name"`
	Surname string          `json:"surname"`
	Alias   string          `json:"alias"`
	Email   string          `json:"email"`
	Wallets []wallet.Wallet `json:"wallets"`
}

func (u *user) ValidateFields() error {
	if !utils.IsValidField(u.Name) || !utils.IsValidField(u.Surname) {
		return fmt.Errorf("bad_request: error user name or surname")
	}
	if !utils.IsValidField(u.Alias) || !utils.IsValidField(u.Email) {
		return fmt.Errorf("bad_request: error user alias or email")
	}

	return nil
}

type NewUserResponse struct {
	ID int64 `json:"user_id"`
}
