package mysql

import (
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

type mockClient struct {
	realClient
	mock sqlmock.Sqlmock
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
