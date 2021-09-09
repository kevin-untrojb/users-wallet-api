package user_crud

import (
	"context"
	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
)
//go:generate mockgen -destination=mock_dao.go -package=user_crud -source=dao.go MySql
type MysqlDao interface {
	InsertUser(context.Context, user) (int64, error)

	GetUser(context.Context, string) (user, error)
}

type dao struct {
	db mysql.Client
}

func (d dao) InsertUser(ctx context.Context, u user) (int64, error) {
	panic("implement me")
}

func (d dao) GetUser(ctx context.Context, s string) (user, error) {
	panic("implement me")
}

func newDao (db mysql.Client) MysqlDao{
	return &dao{db}
}
