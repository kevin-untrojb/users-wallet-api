package mysql

import "database/sql"

type realClient struct {
	client *sql.DB
}