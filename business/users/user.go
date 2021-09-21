package users

import (
	"fmt"

	"github.com/kevin-untrojb/users-wallet-api/business/wallet"
	"github.com/kevin-untrojb/users-wallet-api/utils"
)

type user struct {
	ID        int64           `json:"id"`
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Alias     string          `json:"alias"`
	Email     string          `json:"email"`
	Wallets   []wallet.Wallet `json:"wallets"`
}

func (u *user) ValidateFields() error {
	if !utils.IsValidField(u.FirstName) || !utils.IsValidField(u.LastName) {
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
