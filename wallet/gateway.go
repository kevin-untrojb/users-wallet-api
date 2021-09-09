package wallet

import "github.com/kevin-untrojb/users-wallet-api/internal/mysql"

type Gateway interface {
}
type gateway struct {
	db MysqlDao
}

func NewGateway(dbClient mysql.Client) Gateway {
	return &gateway{db: newDao(dbClient)}
}
