package users

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
)

//go:generate mockgen -destination=mock_dao.go -package=users -source=dao.go MySql
type MysqlDao interface {
	InsertUser(context.Context, user) (int64, error)

	GetUser(context.Context, int64) (user, error)
}

type dao struct {
	db mysql.Client
}

func newDao(db mysql.Client) MysqlDao {
	return &dao{db}
}

const (
	insertUserQuery = "insert into user (first_name, last_name, alias, email, date_created) VALUES ( ?, ?, ?, ?, UTC_TIMESTAMP())"
	checkUserQuery  = "select count(*) from user u where u.email = ? or u.alias = ?"
	getUserQuery    = "select ID, NAME, SURNAME,ALIAS, EMAIL from user where ID = ?"
)

func (d dao) InsertUser(ctx context.Context, u user) (int64, error) {
	var lastUserID int64
	ctx, cancel := context.WithTimeout(ctx, mysql.MediumTimeout)
	defer cancel()

	err := d.db.WithTransaction(func(trx *sql.Tx) error {
		var count int64
		row := d.db.RawQueryRow(ctx, trx, checkUserQuery, u.Email, u.Alias)
		err := row.Scan(&count)
		if err != nil {
			// todo handler
			return err
		}
		if count != 0 {
			return fmt.Errorf("error, email or alias is already used")
		}

		exec, err := d.db.RawExec(ctx, trx, insertUserQuery, u.Name, u.Surname, u.Alias, u.Email)
		if err != nil {
			// todo handler
			return err
		}
		lastUserID, err = exec.LastInsertId()
		if err != nil {
			// todo handler
			return err
		}
		return nil
	})
	return lastUserID, err
}

func (d dao) GetUser(ctx context.Context, userID int64) (user, error) {
	var u user
	ctx, cancel := context.WithTimeout(ctx, mysql.MediumTimeout)
	defer cancel()

	row := d.db.RawQueryRow(ctx, nil, getUserQuery, userID)
	err := row.Scan(&u.ID, &u.Name, &u.Surname, &u.Alias, &u.Email)
	if err != nil {
		return u, fmt.Errorf("get_user: error getting user %w", err)
	}
	return u, nil
}
