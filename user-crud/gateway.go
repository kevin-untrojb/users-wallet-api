package user_crud

import (
	"context"
	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
)

//go:generate mockgen -destination=mock_gateway.go -package=user_crud -source=gateway.go Gateway

type Gateway interface {
	Create(context.Context, user) error

	Get(context.Context, string) error

}

type gateway struct {
	db MysqlDao
}

func (g gateway) Create(ctx context.Context, u user) error {
	panic("implement me")
}

func (g gateway) Get(ctx context.Context, s string) error {
	panic("implement me")
}

func NewGateway(dbClient mysql.Client) Gateway{
	return &gateway{db: newDao(dbClient)}
}
