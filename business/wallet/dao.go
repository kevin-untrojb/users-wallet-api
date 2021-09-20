package wallet

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"

	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
)

//go:generate mockgen -destination=mock_dao.go -package=wallet -source=dao.go MySql
type MysqlDao interface {
	SearchTransactions(ctx context.Context, userID int64, params *SearchRequestParams) (SearchResponse, error)
	NewTransaction(ctx context.Context, transaction Transaction) (int64, error)
	GetWalletsForUser(ctx context.Context, userID int64) ([]Wallet, error)
}

type dao struct {
	db mysql.Client
}

func newDao(db mysql.Client) MysqlDao {
	return &dao{db}
}

const (
	insertTransaction = "insert into transaction (wallet_id, transaction_type, amount, date_created) VALUES ( ?, ?, ?, UTC_TIMESTAMP())"

	getWalletsAndCurrenciesForUser = "SELECT w.ID, w.CURRENT_BALANCE, c.NAME, c.EXPONENT FROM WALLET w INNER JOIN USER u ON u.ID = w.USER_ID INNER JOIN CURRENCY c on c.id = w.CURRENCY_ID WHERE u.ID = ?"
	getWalletAnCurencyByWalletID   = "SELECT w.ID, w.CURRENT_BALANCE, c.NAME, c.EXPONENT FROM WALLET w INNER JOIN CURRENCY c on c.id = w.CURRENCY_ID WHERE w.ID = ?"
	updateBalanceOFAWallet         = "UPDATE WALLET w SET w.CURRENT_BALANCE = ? WHERE w.ID = ?"
)

func (d dao) SearchTransactions(ctx context.Context, userID int64, params *SearchRequestParams) (SearchResponse, error) {
	var response SearchResponse
	transactionsResults := make([]Transaction, 0)

	err := d.db.WithTransaction(func(trx *sql.Tx) error {
		searchQuery, searchQueryParams, err := CreateSearchQuery(userID, params, false)
		if err != nil {
			// todo log
			return err
		}
		ctx, cancel := context.WithTimeout(ctx, mysql.MediumTimeout)
		defer cancel()

		rows, err := d.db.RawQuery(ctx, trx, searchQuery, searchQueryParams...)
		if err != nil {
			// todo log
			return err
		}
		defer rows.Close()

		for rows.Next() {
			currentTransaction := Transaction{}
			err := rows.Scan(&currentTransaction.ID,
				&currentTransaction.TransactionType,
				&currentTransaction.Amount,
				&currentTransaction.Date,
				&currentTransaction.CurrencyName)

			if err != nil {
				// todo compeltar
				return err
			}
			transactionsResults = append(transactionsResults, currentTransaction)
		}
		response.Results = transactionsResults

		CountQuery, countQueryParams, err := CreateSearchQuery(userID, params, true)
		if err != nil {
			// todo log
			return err
		}
		ctx, cancel = context.WithTimeout(ctx, mysql.MediumTimeout)
		defer cancel()

		row := d.db.RawQueryRow(ctx, trx, CountQuery, countQueryParams...)
		err = row.Scan(&response.Paging.Total)
		if err != nil {
			return err
		}

		return nil
	})
	response.Paging.Limit = params.Limit
	response.Paging.Offset = params.Offset

	return response, err
}

func (d dao) GetWalletsForUser(ctx context.Context, userID int64) ([]Wallet, error) {
	var walletResults []Wallet
	var ok bool

	ctx, cancel := context.WithTimeout(ctx, mysql.MediumTimeout)
	defer cancel()

	rows, err := d.db.RawQuery(ctx, nil, getWalletsAndCurrenciesForUser, userID)
	if err != nil {
		// todo log
		return walletResults, err
	}
	defer rows.Close()

	for rows.Next() {
		currentWallet := Wallet{}

		err := rows.Scan(&currentWallet.ID,
			&currentWallet.CurrentBalance,
			&currentWallet.CurrencyName, &currentWallet.CointExponent)

		if err != nil {
			// todo handle
			return walletResults, err
		}

		currentWallet.Coin, ok = newCoin(currentWallet.CurrentBalance, currentWallet.CointExponent)
		if !ok {
			return nil, fmt.Errorf("error converting %s into a number", currentWallet.CurrentBalance)
		}

		walletResults = append(walletResults, currentWallet)
	}
	return walletResults, nil
}

func (d dao) NewTransaction(ctx context.Context, transaction Transaction) (int64, error) {
	var lastID int64
	var w Wallet

	err := d.db.WithTransaction(func(trx *sql.Tx) error {
		ctx, cancel := context.WithTimeout(ctx, mysql.MediumTimeout)
		defer cancel()

		row := d.db.RawQueryRow(ctx, trx, getWalletAnCurencyByWalletID, transaction.WalletID)
		err := row.Scan(&w.ID, &w.CurrentBalance, &w.CurrencyName, &w.CointExponent)
		if err != nil {
			return err
		}
		if err := w.TryNewTransaction(transaction); err != nil {
			return err
		}
		exec, err := d.db.RawExec(ctx, trx, updateBalanceOFAWallet, w.GetCurrentBalance(), w.ID)
		if err != nil {
			// todo handler
			return err
		}
		_, err = exec.LastInsertId()
		if err != nil {
			// todo handler
			return err
		}

		exec, err = d.db.RawExec(ctx, trx, insertTransaction, transaction.WalletID, transaction.TransactionType, transaction.Amount)
		if err != nil {
			// todo handler
			return err
		}
		lastID, err = exec.LastInsertId()
		if err != nil {
			// todo handler
			return err
		}

		return nil
	})

	return lastID, err
}

func CreateSearchQuery(userID int64, params *SearchRequestParams, isCountQuery bool) (string, []interface{}, error) {
	var query bytes.Buffer
	queryParams := make([]interface{}, 0)

	if params == nil {
		return "", nil, fmt.Errorf("nil params")
	}
	if isCountQuery {
		query.WriteString("SELECT count(*) FROM TRANSACTION t")
	} else {
		query.WriteString("SELECT t.ID, t.TRANSACTION_TYPE, t.AMOUNT, t.DATE_CREATED, c.NAME FROM TRANSACTION")
	}

	query.WriteString(" INNER JOIN WALLET w ON t.WALLET_ID = w.ID")
	query.WriteString(" INNER JOIN USER u ON u.ID = w.USER_ID")
	query.WriteString(" INNER JOIN CURRENCY c on c.id = w.CURRENCY_ID")
	query.WriteString(" WHERE")

	if params.Currency != "" {

		query.WriteString(" c.NAME=? AND")
		queryParams = append(queryParams, params.Currency)
	}
	query.WriteString(" u.ID =?")
	queryParams = append(queryParams, userID)

	if params.MovementType != "" {
		query.WriteString(" AND m.TRANSACTION_TYPE=?")
		queryParams = append(queryParams, params.MovementType)
	}

	if !isCountQuery {
		query.WriteString(" ORDER BY m.DATE_CREATED DESC")

		query.WriteString(" LIMIT ? OFFSET ?") // Required for pagination
		queryParams = append(queryParams, params.Limit, params.Offset)
	}

	return query.String(), queryParams, nil
}
