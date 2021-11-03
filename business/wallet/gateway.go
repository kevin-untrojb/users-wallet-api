package wallet

import (
	"context"

	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
)

//go:generate mockgen -destination=mock_gateway.go -package=wallet -source=gateway.go Gateway
type Gateway interface {
	GetWalletsFroUser(context.Context, int64) ([]Wallet, error)
	CreateDefaultWalletsForUser(context.Context, int64) ([]Wallet, error)
	SearchTransactionsForUser(ctx context.Context, userID int64, params *SearchRequestParams) (SearchResponse, error)
	NewTransaction(ctx context.Context, nt Transaction) (Transaction, error)
}
type gateway struct {
	db MysqlDao
}

func (g gateway) CreateDefaultWalletsForUser(ctx context.Context, userID int64) ([]Wallet, error) {
	return g.db.CreateDefaultWalletsForUser(ctx, userID)
}

func (g gateway) GetWalletsFroUser(ctx context.Context, userID int64) ([]Wallet, error) {
	wallets, err := g.db.GetWalletsForUser(ctx, userID)
	if err != nil {
		// todo handler
		return nil, err
	}
	for i := range wallets {
		wallets[i] = wallets[i].ToUserWallet()
	}
	return wallets, nil
}

func (g gateway) SearchTransactionsForUser(ctx context.Context, userID int64, params *SearchRequestParams) (SearchResponse, error) {
	return g.db.SearchTransactions(ctx, userID, params)
}

func (g gateway) NewTransaction(ctx context.Context, nt Transaction) (Transaction, error) {
	return g.db.NewTransaction(ctx, nt)
}

func NewGateway(dbClient mysql.Client) Gateway {
	return &gateway{db: newDao(dbClient)}
}
