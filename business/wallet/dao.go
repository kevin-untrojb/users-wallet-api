package wallet

import (
	"bytes"
	"context"
	"fmt"
	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
)

type MysqlDao interface {
	SearchTransactions(ctx context.Context, userID string, params *SearchRequestParams)(SearchResponse,error)

	GetWalletsForUser(ctx context.Context,userID string)
}

type dao struct {
	db mysql.Client
}


func newDao(db mysql.Client) MysqlDao {
	return &dao{db}
}

func (d dao) SearchTransactions(ctx context.Context, userID string, params *SearchRequestParams) (SearchResponse, error) {
	var response SearchResponse
	transactionsResults := make([]Transaction,0)

	query, queryParams, err := CreateSearchQuery(userID, params)
	if err != nil {
		// todo log
		return response, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), mysql.MediumTimeout)
	defer cancel()

	rows, err := d.db.RawQuery(ctx, nil, query, queryParams...)
	if err != nil {
		// todo log
		return response,err
	}
	defer rows.Close()

	for rows.Next() {
		var rank int
		currentTransaction := Transaction{}
		// todo completar
		err := rows.Scan(&currentTransaction.ID,&currentTransaction.UserID,&currentTransaction.TransactionType)

		if err != nil{
			// todo compeltar
			return response, err
		}

	}



}

func (d dao) GetWalletsForUser(ctx context.Context, userID string) {
	panic("implement me")
}


func CreateSearchQuery(userID string, params *SearchRequestParams) (string,[]interface{}, error) {
	var query bytes.Buffer
	queryParams := make([]interface{}, 0)

	if params == nil {
		return "", nil, fmt.Errorf("nil params")
	}

	query.WriteString("SELECT q.* FROM (")

	query.WriteString("SELECT m.ID, m.USER_ID m.TRANSACTION_TYPE m.CURRENCY_ID, m.CURRENT_BALANCE, m.DATE_CREATED")
	query.WriteString(", c.ID AS CURRENCY_ID, c.NAME, c.EXPONENT, pr.")
	query.WriteString(", DENSE_RANK() OVER (ORDER BY m.DATE_CREATED DESC) rnk")

	query.WriteString(" FROM TRANSACTION AS m INNER JOIN WALLET AS w ON (w.ID = m.WALLET_ID)")
	query.WriteString(" FROM CURRENCY AS c INNER JOIN WALLET AS w ON (c.ID = m.CURRENCY_ID)")

	query.WriteString(" WHERE m.USER_ID= ? ")
	queryParams = append(queryParams, userID)

	if params.Currency != ""{
		query.WriteString(" AND c.name=?")
		queryParams = append(queryParams, params.Currency)
	}
	if params.MovementType != ""{
		query.WriteString(" AND m.TRANSACTION_TYPE=?")
		queryParams = append(queryParams, params.MovementType)
	}

	query.WriteString(" ORDER BY m.DATE_CREATED DESC")

	query.WriteString(") AS q WHERE q.rnk BETWEEN ? AND ?")                        // Required for pagination
	queryParams = append(queryParams, params.Offset+1, params.Offset+params.Limit)

	return query.String(), queryParams, nil
}