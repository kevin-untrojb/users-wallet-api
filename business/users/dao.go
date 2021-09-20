package users

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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
	getUserQuery    = "select id, first_name, last_name, alias, email from user where id = ?"
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
			log.Println(fmt.Sprintf("error checking if user into db email %s, alias: %s, count %d, error: %s", u.Email, u.Alias, count, err.Error()))
			return err
		}
		if count != 0 {
			return fmt.Errorf("error, email or alias is already used")
		}

		exec, err := d.db.RawExec(ctx, trx, insertUserQuery, u.FirstName, u.LastName, u.Alias, u.Email)
		if err != nil {
			log.Println(fmt.Sprintf("error inserting user into db query:%s, error: %s", insertUserQuery, err.Error()))
			return err
		}
		lastUserID, err = exec.LastInsertId()
		if err != nil {
			log.Println(fmt.Sprintf("error last inserting id inserting error: %s", err.Error()))
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
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Alias, &u.Email)
	if err != nil {
		log.Println(fmt.Sprintf("db error: gettin user %d: %s", userID, err.Error()))
		return u, fmt.Errorf("get_user: error getting user %w", err)
	}
	return u, nil
}
