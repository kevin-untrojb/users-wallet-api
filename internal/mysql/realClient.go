package mysql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

type realClient struct {
	client *sql.DB
}

func (c realClient) RawQuery(ctx context.Context, tx *sql.Tx, query string, params ...interface{}) (*sql.Rows, error) {
	if tx != nil {
		return tx.QueryContext(ctx, query, params...)
	}
	return c.client.QueryContext(ctx, query, params...)
}

func (c realClient) RawQueryRow(ctx context.Context, tx *sql.Tx, query string, params ...interface{}) *sql.Row {
	if tx != nil {
		return tx.QueryRowContext(ctx, query, params...)
	}
	return c.client.QueryRowContext(ctx, query, params...)
}

func (c realClient) RawExec(ctx context.Context, tx *sql.Tx, query string, params ...interface{}) (driver.Result, error) {
	if tx != nil {
		return tx.ExecContext(ctx, query, params...)
	}
	return c.client.ExecContext(ctx, query, params...)
}

func (c realClient) AddExpectedQueryWithRows(rawSQL string, rows *sqlmock.Rows, params ...driver.Value) {

}

func (c realClient) AddExpectedQueryWithError(rawSQL string, err error, params ...driver.Value) {

}

func (c realClient) AddExpectedExec(rawSQL string, result driver.Result, params ...driver.Value) {

}

func (c realClient) AddExpectedExecWithError(rawSQL string, err error, params ...driver.Value) {

}

func (c realClient) WithTransaction(txFunc func(*sql.Tx) error) error {
	tx, err := c.client.Begin()
	if err != nil {
		return errors.New("error beginning transaction")
	}
	var errTx error
	defer func() {
		if p := recover(); p != nil {
			err = tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if errTx != nil {
			rbErr := tx.Rollback() // err is non-nil; don't change it
			if rbErr != nil {
				err = rbErr
			}
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()

	errTx = txFunc(tx)
	return errTx
}

func makeRealClient(driver, dataSourceName string, maxConnections int, connMaxLifeTime time.Duration) (Client, error) {
	time.Sleep(1*time.Second)
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
