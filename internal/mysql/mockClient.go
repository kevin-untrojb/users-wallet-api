package mysql

import (
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

type mockClient struct {
	realClient
	mock sqlmock.Sqlmock
}

func (mc *mockClient) AddExpectedQueryWithError(rawSQL string, err error, params ...driver.Value) {
	mc.mock.ExpectQuery(rawSQL).WithArgs(params...).WillReturnError(err)
}

func (mc *mockClient) AddExpectedQueryWithRows(rawSQL string, rows *sqlmock.Rows, params ...driver.Value) {
	mc.mock.ExpectQuery(rawSQL).WithArgs(params...).WillReturnRows(rows)
}

func (mc *mockClient) AddExpectedExec(rawSQL string, result driver.Result, params ...driver.Value) {
	eq := mc.mock.ExpectExec(rawSQL).WithArgs(params...)
	if result == nil {
		eq.WillReturnResult(sqlmock.NewResult(0, 0))
		return
	}
	eq.WillReturnResult(result)
}

func (mc *mockClient) AddExpectedExecWithError(rawSQL string, err error, params ...driver.Value) {
	mc.mock.ExpectExec(rawSQL).WithArgs(params...).WillReturnError(err)
}

//WithTransaction is a high order function that supplies a database transaction in order
//	to manage rollbacks and commits in one single place.
func (mc *mockClient) WithTransaction(txFunc func(*sql.Tx) error) error {
	mc.mock.ExpectBegin()
	tx, err := mc.realClient.client.Begin()
	if err != nil {
		return err
	}
	var errTx error
	defer func() {
		if p := recover(); p != nil {
			mc.mock.ExpectRollback()
			err = tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if errTx != nil {
			mc.mock.ExpectRollback()
			rbErr := tx.Rollback() // err is non-nil; don't change it
			if rbErr != nil {
				err = rbErr
			}
		} else {
			mc.mock.ExpectCommit()
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()
	errTx = txFunc(tx)

	return errTx
}
func (mc *mockClient) ExpectationsWereMet() error {
	return mc.mock.ExpectationsWereMet()
}

func makeMockClient(driver, dataSourceName string, maxConnections int, connMaxLifeTime time.Duration) (Client, error) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		return nil, err
	}
	mock.MatchExpectationsInOrder(false)
	mc := mockClient{
		realClient{db},
		mock,
	}
	return &mc, nil
}
