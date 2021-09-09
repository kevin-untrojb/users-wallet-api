package wallet

import (
	"time"
)

type Wallet struct {
	ID             string
	UserID         string
	currentBalance string
	currency       Currency
	Movements      []Movement
}
type Movement struct {
	ID           string
	WalletID     string
	MovementType string
	UserID       string
	Date         time.Time
	Amount       string
}
