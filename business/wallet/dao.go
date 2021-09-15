package wallet

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"

	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
)

type MysqlDao interface {
	SearchTransactions(ctx context.Context, userID string, params *SearchRequestParams) (SearchResponse, error)

	GetWalletsForUser(ctx context.Context, userID string)
}

type dao struct {
	db mysql.Client
}

func newDao(db mysql.Client) MysqlDao {
	return &dao{db}
}

func (d dao) SearchTransactions(ctx context.Context, userID string, params *SearchRequestParams) (SearchResponse, error) {
	var response SearchResponse
	transactionsResults := make([]Transaction, 0)

	err := d.db.WithTransaction(func(trx *sql.Tx) error {
		searchQuery, searchQueryParams, err := CreateSearchQuery(userID, params, false)
		if err != nil {
			// todo log
			return  err
		}
		ctx, cancel := context.WithTimeout(ctx, mysql.MediumTimeout)
		defer cancel()

		rows, err := d.db.RawQuery(ctx, trx, searchQuery, searchQueryParams...)
		if err != nil {
			// todo log
			return  err
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
			transactionsResults = append(transactionsResults,currentTransaction)
		}
		response.Results = transactionsResults

		CountQuery, countQueryParams, err := CreateSearchQuery(userID, params, true)
		if err != nil {
			// todo log
			return  err
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

	return response, err
}

func (d dao) GetWalletsForUser(ctx context.Context, userID string) {
	panic("implement me")
}

func CreateSearchQuery(userID string, params *SearchRequestParams, isCountQuery bool) (string, []interface{}, error) {
	var query bytes.Buffer
	queryParams := make([]interface{}, 0)

	if params == nil {
		return "", nil, fmt.Errorf("nil params")
	}
	if isCountQuery{
		query.WriteString("SELECT count(*) FROM TRANSACTION t")
	}else {
		query.WriteString("SELECT t.ID, t.TRANSACTION_TYPE, t.AMOUNT, t.DATE_CREATED, c.NAME from Transaction")
	}

	query.WriteString("INNER JOIN WALLET w ON t.WALLET_ID = w.ID")
	query.WriteString("INNER JOIN USER u ON u.ID = w.USER_ID")
	query.WriteString("INNER JOIN CURRENCY c on c.id = w.CURRENCY_ID")
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

	if !isCountQuery{
		query.WriteString(" ORDER BY m.DATE_CREATED DESC")

		query.WriteString(" LIMIT ? OFFSET ?") // Required for pagination
		queryParams = append(queryParams, params.Offset+1, params.Offset+params.Limit)
	}

	return query.String(), queryParams, nil
}
