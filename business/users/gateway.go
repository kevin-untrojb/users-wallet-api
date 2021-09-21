package users

import (
	"context"
	"fmt"
	"log"

	"github.com/kevin-untrojb/users-wallet-api/business/wallet"
	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
)

//go:generate mockgen -destination=mock_gateway.go -package=users -source=gateway.go Gateway

type Gateway interface {
	Create(context.Context, user) (user, error)

	Get(context.Context, int64) (user, error)
}

type gateway struct {
	db     MysqlDao
	wallet wallet.Gateway
}

func (g gateway) Create(ctx context.Context, u user) (user, error) {
	return g.db.InsertUser(ctx, u)
}

func (g gateway) Get(ctx context.Context, userID int64) (user, error) {
	user, err := g.db.GetUser(ctx, userID)
	if err != nil {
		log.Println(fmt.Sprintf("error gettting user %d: %s", userID, err.Error()))
		return user, fmt.Errorf("users_error: error getting user form db")
	}
	user.Wallets, err = g.wallet.GetWalletsFroUser(ctx, userID)
	if err != nil {
		log.Println(fmt.Sprintf("error gettting wallets of user %d: %s", userID, err.Error()))
		return user, fmt.Errorf("users_error: error getting transactions form db")
	}
	return user, nil
}

func NewGateway(dbClient mysql.Client, wallet wallet.Gateway) Gateway {
	return &gateway{db: newDao(dbClient), wallet: wallet}
}
