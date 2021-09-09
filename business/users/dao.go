package users

import (
	"context"
	"fmt"

	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
)

//go:generate mockgen -destination=mock_dao.go -package=users -source=dao.go MySql
type MysqlDao interface {
	InsertUser(context.Context, user) (int64, error)

	GetUser(context.Context, string) (user, error)
}

type dao struct {
	db mysql.Client
}

func newDao(db mysql.Client) MysqlDao {
	return &dao{db}
}

const (
	insertUserQuery = ""

	getUserQuery = ""
)

func (d dao) InsertUser(ctx context.Context, u user) (int64, error) {
	panic("implement me")
}

func (d dao) GetUser(ctx context.Context, userID string) (user, error) {
	var u user
	row := d.db.RawQueryRow(ctx, nil, getUserQuery, userID)
	err := row.Scan(&u.ID, &u.Name, &u.Surname, &u.Alias, &u.email)
	if err != nil {
		return u, fmt.Errorf("get_user: error getting user %w", err)
	}
	return u, nil
}
