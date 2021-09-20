package mysql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"os"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/kevin-untrojb/users-wallet-api/internal/host"
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

var (
	mysqlPassword         = os.Getenv("MYSQL_PASSWORD")
	mysqlUser             = os.Getenv("MYSQL_USER")
	mysqlDB               = os.Getenv("MYSQL_DATABASE")
	mysqlPort             = os.Getenv("MYSQL_PORT")
	mysqlConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mysqlUser, mysqlPassword, mysqlDB, mysqlPort, mysqlDB)
	dsn                   = fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s sslmode=disable", mysqlUser, mysqlPassword, "test_db", mysqlPort, "mysqldb")
)

func NewClient() Client {

	dbConfig := mysql.NewConfig()
	dbConfig.User = mysqlUser
	dbConfig.Passwd = mysqlPassword
	dbConfig.Addr = fmt.Sprintf("%s:%s", "mysqldb", mysqlPort)
	dbConfig.DBName = "test_db"
	dbConfig.Net = "tcp"
	MakeClient = makeMockClient

	if host.IsProduction() {
		MakeClient = makeRealClient
	}
	time.Sleep(2 * time.Second)
	db, err := MakeClient(driverName, "test_db:root@tcp(mysqldb:3306)/test_db?parseTime=true", MinConnections, LongTimeout)
	if err != nil {
		panic(err)
	}

	return db
}
