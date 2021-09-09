package wallet

import (
	"context"

	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
)

type Gateway interface {
	GetWalletsFroUser(context.Context, string) ([]Wallet, error)
	SearchMovementsForUser(ctx context.Context)
	NewMovement(ctx context.Context)
}
type gateway struct {
	db MysqlDao
}

func (g gateway) GetWalletsFroUser(ctx context.Context, s string) ([]Wallet, error) {
	panic("implement me")
}

func (g gateway) SearchMovementsForUser(ctx context.Context) {
	panic("implement me")
}

func (g gateway) NewMovement(ctx context.Context) {
	panic("implement me")
}

func NewGateway(dbClient mysql.Client) Gateway {
	return &gateway{db: newDao(dbClient)}
}
