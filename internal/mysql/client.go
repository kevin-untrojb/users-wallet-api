package mysql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

type Client interface {
	RawQuery(ctx context.Context, trx *sql.Tx, query string, params ...interface{}) (*sql.Rows, error)
	RawQueryRow(ctx context.Context, trx *sql.Tx, query string, params ...interface{}) *sql.Row
	RawExec(ctx context.Context, trx *sql.Tx, query string, params ...interface{}) (driver.Result, error)
	AddExpectedQueryWithRows(rawSQL string, rows *sqlmock.Rows, params ...driver.Value)
	AddExpectedQueryWithError(rawSQL string, err error, params ...driver.Value)
	AddExpectedExec(rawSQL string, result driver.Result, params ...driver.Value)
	AddExpectedExecWithError(rawSQL string, err error, params ...driver.Value)

	WithTransaction(txFunc func(*sql.Tx) error) error
}

var MakeClient func(driver, connection string, connections int, connMaxLifeTime time.Duration) (Client, error)

const (
	driverName = "mysql"


	MediumTimeout   = 30 * time.Second
	LongTimeout     = 5 * time.Minute
	MinConnections  = 10
	LimitSearchRows = 100
)

func NewClient() Client {

	MakeClient = makeRealClient
	db, err := MakeClient(driverName, "", MinConnections, LongTimeout)
	if err != nil {
		panic(err)
	}

	return db
}
