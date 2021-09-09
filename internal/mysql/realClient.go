package mysql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"time"
)

type realClient struct {
	client *sql.DB
}

func (r realClient) RawQuery(ctx context.Context, trx *sql.Tx, query string, params ...interface{}) (*sql.Rows, error) {
	panic("implement me")
}

func (r realClient) RawQueryRow(ctx context.Context, trx *sql.Tx, query string, params ...interface{}) *sql.Row {
	panic("implement me")
}

func (r realClient) RawExec(ctx context.Context, trx *sql.Tx, query string, params ...interface{}) (driver.Result, error) {
	panic("implement me")
}

func (r realClient) AddExpectedQueryWithRows(rawSQL string, rows *sqlmock.Rows, params ...driver.Value) {
	panic("implement me")
}

func (r realClient) AddExpectedQueryWithError(rawSQL string, err error, params ...driver.Value) {
	panic("implement me")
}

func (r realClient) AddExpectedExec(rawSQL string, result driver.Result, params ...driver.Value) {
	panic("implement me")
}

func (r realClient) AddExpectedExecWithError(rawSQL string, err error, params ...driver.Value) {
	panic("implement me")
}

func (r realClient) WithTransaction(txFunc func(*sql.Tx) error) error {
	panic("implement me")
}

func makeRealClient(driver, dataSourceName string, maxConnections int, connMaxLifeTime time.Duration) (Client, error) {
	db, err := sql.Open(driver, dataSourceName)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxConnections)
	db.SetMaxIdleConns(maxConnections / 4)
	db.SetConnMaxLifetime(connMaxLifeTime)

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &realClient{db}, nil
}
