package wallet

import (
	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
)

type MysqlDao interface {
}

type dao struct {
	db mysql.Client
}

func newDao(db mysql.Client) MysqlDao {
	return &dao{db}
}
