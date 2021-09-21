package wallet

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
)

//go:generate mockgen -destination=mock_dao.go -package=wallet -source=dao.go MySql
type MysqlDao interface {
	SearchTransactions(ctx context.Context, userID int64, params *SearchRequestParams) (SearchResponse, error)
	NewTransaction(ctx context.Context, transaction Transaction) (Transaction, error)
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

	getWalletsAndCurrenciesForUser = "SELECT w.id, w.current_balance, c.name, c.exponent FROM wallet w INNER JOIN user u ON u.id = w.user_id INNER JOIN currency c on c.id = w.currency_id WHERE u.id = ?"
	getWalletAnCurrencyByWalletID  = "SELECT w.id, w.current_balance, c.name, c.exponent FROM wallet w INNER JOIN currency c on c.id = w.currency_id WHERE w.ID = ?"
	updateBalanceOFAWallet         = "UPDATE wallet w SET w.current_balance = ? WHERE w.id = ?"
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
			log.Println(fmt.Sprintf("error searching transactins query:%s, params %v, error: %s", searchQuery, searchQueryParams, err.Error()))
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
				log.Println(fmt.Sprintf("row error: %s", err.Error()))
				return err
			}
			transactionsResults = append(transactionsResults, currentTransaction)
		}
		response.Results = transactionsResults

		CountQuery, countQueryParams, err := CreateSearchQuery(userID, params, true)
		if err != nil {
			log.Println(fmt.Sprintf("error counting transactins query:%s, params %v, error: %s", CountQuery, countQueryParams, err.Error()))
			return err
		}
		ctx, cancel = context.WithTimeout(ctx, mysql.MediumTimeout)
		defer cancel()

		row := d.db.RawQueryRow(ctx, trx, CountQuery, countQueryParams...)
		err = row.Scan(&response.Paging.Total)
		if err != nil {
			log.Println(fmt.Sprintf("row error: %s", err.Error()))
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
			&currentWallet.CurrencyName, &currentWallet.CoinExponent)

		if err != nil {
			// todo handle
			return walletResults, err
		}

		currentWallet.Coin, ok = newCoin(currentWallet.CurrentBalance, currentWallet.CoinExponent)
		if !ok {
			return nil, fmt.Errorf("error converting %s into a number", currentWallet.CurrentBalance)
		}

		walletResults = append(walletResults, currentWallet)
	}
	return walletResults, nil
}

func (d dao) NewTransaction(ctx context.Context, transaction Transaction) (Transaction, error) {
	var w Wallet
	var ok bool

	err := d.db.WithTransaction(func(trx *sql.Tx) error {
		ctx, cancel := context.WithTimeout(ctx, mysql.MediumTimeout)
		defer cancel()

		row := d.db.RawQueryRow(ctx, trx, getWalletAnCurrencyByWalletID, transaction.WalletID)
		err := row.Scan(&w.ID, &w.CurrentBalance, &w.CurrencyName, &w.CoinExponent)
		if err != nil {
			log.Println(fmt.Sprintf("error getting wallet %d : %s", transaction.WalletID, err.Error()))
			return err
		}
		w.Coin, ok = newCoin(w.CurrentBalance, w.CoinExponent)
		if !ok {
			log.Println(fmt.Sprintf("error creating a new coin"))
			return fmt.Errorf("error creating a new coin")
		}
		if err := w.TryNewTransaction(transaction); err != nil {
			log.Println(fmt.Sprintf("error tring transaction wallet %d : %s", transaction.WalletID, err.Error()))
			return err
		}
		exec, err := d.db.RawExec(ctx, trx, updateBalanceOFAWallet, w.GetCurrentBalance(), w.ID)
		if err != nil {
			log.Println(fmt.Sprintf("error updating wallet %d: %s", w.ID, err.Error()))
			return err
		}
		_, err = exec.LastInsertId()
		if err != nil {
			log.Println(fmt.Sprintf("error updating wallet id %d: %s", w.ID, err.Error()))
			return err
		}

		exec, err = d.db.RawExec(ctx, trx, insertTransaction, transaction.WalletID, transaction.TransactionType, transaction.Amount)
		if err != nil {
			log.Println(fmt.Sprintf("error inserting new transaction :%s", err.Error()))
			return err
		}
		transaction.ID, err = exec.LastInsertId()
		if err != nil {
			log.Println(fmt.Sprintf("error inserting new transaction %d:%s", transaction.ID, err.Error()))
			return err
		}

		return nil
	})

	return transaction, err
}

func CreateSearchQuery(userID int64, params *SearchRequestParams, isCountQuery bool) (string, []interface{}, error) {
	var query bytes.Buffer
	queryParams := make([]interface{}, 0)

	if params == nil {
		return "", nil, fmt.Errorf("nil params")
	}
	if isCountQuery {
		query.WriteString("SELECT count(*) FROM transaction t")
	} else {
		query.WriteString("SELECT t.id, t.transaction_type, t.amount, t.date_created, c.name FROM transaction t")
	}

	query.WriteString(" INNER JOIN wallet w ON t.wallet_id = w.id")
	query.WriteString(" INNER JOIN user u ON u.id = w.user_id")
	query.WriteString(" INNER JOIN currency c on c.id = w.currency_id")
	query.WriteString(" WHERE")

	if params.Currency != "" {

		query.WriteString(" c.name=? AND")
		queryParams = append(queryParams, params.Currency)
	}
	query.WriteString(" u.id =?")
	queryParams = append(queryParams, userID)

	if params.MovementType != "" {
		query.WriteString(" AND t.transaction_type=?")
		queryParams = append(queryParams, params.MovementType)
	}

	if !isCountQuery {
		query.WriteString(" ORDER BY t.date_created DESC")

		query.WriteString(" LIMIT ? OFFSET ?") // Required for pagination
		queryParams = append(queryParams, params.Limit, params.Offset)
	}

	return query.String(), queryParams, nil
}
