package wallet

import (
	"time"
)

type Wallet struct {
	id            string
	walletType    string
	currentAmount Coins
	Movements     []Movement
}
type Movement struct {
	id           string
	movementType string
	userID       string
	date         time.Time
	amount       Coins
}
