package wallet

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
	"github.com/stretchr/testify/assert"
)

func TestSearchTransactionForUserShouldReturnOK(t *testing.T) {
	ctx := context.Background()
	dbMockClient := mysql.Connect()
	userID := int64(123123)
	currencyName := "BTC"
	params := &SearchRequestParams{
		Limit:        10,
		Offset:       1,
		MovementType: "Deposit",
		Currency:     currencyName,
	}

	transactions := []Transaction{
		{ID: int64(1), TransactionType: "Deposit", Amount: "0.0034200", Date: time.Time{}, CurrencyName: "BTC"},
		{ID: int64(2), TransactionType: "Deposit", Amount: "0.0033200", Date: time.Time{}, CurrencyName: "BTC"},
	}

	countQuery := "SELECT count(*) FROM TRANSACTION t INNER JOIN WALLET w ON t.WALLET_ID = w.ID INNER JOIN USER u ON u.ID = w.USER_ID INNER JOIN CURRENCY c on c.id = w.CURRENCY_ID WHERE c.NAME=? AND u.ID =? AND m.TRANSACTION_TYPE=?"
	totalCount := 10

	searchQuery := "SELECT t.ID, t.TRANSACTION_TYPE, t.AMOUNT, t.DATE_CREATED, c.NAME FROM TRANSACTION INNER JOIN WALLET w ON t.WALLET_ID = w.ID INNER JOIN USER u ON u.ID = w.USER_ID INNER JOIN CURRENCY c on c.id = w.CURRENCY_ID WHERE c.NAME=? AND u.ID =? AND m.TRANSACTION_TYPE=? ORDER BY m.DATE_CREATED DESC LIMIT ? OFFSET ?\n"
	countRows := sqlmock.NewRows([]string{"COUNT"}).AddRow(totalCount)

	dbMockClient.AddExpectedQueryWithRows(countQuery, countRows, params.Currency, userID, params.MovementType)
	searchRows := sqlmock.NewRows([]string{"ID", "TRANSACTION_TYPE", "AMOUNT", "DATE_CREATED", "CURRENCY_NAME"})
	for _, transaction := range transactions {
		searchRows.AddRow(transaction.ID, transaction.TransactionType, transaction.Amount, transaction.Date, currencyName)
	}

	dbMockClient.AddExpectedQueryWithRows(searchQuery, searchRows, params.Currency, userID, params.MovementType, params.Limit, params.Offset)

	res, err := newDao(dbMockClient).SearchTransactions(ctx, userID, params)

	assert.Nil(t, err)
	assert.Equal(t, res.Paging.Total, totalCount)
	assert.Equal(t, res.Results, transactions)
}

func TestGetWalletsForUserOKResponse(t *testing.T) {
	ctx := context.Background()
	dbMockClient := mysql.Connect()
	userID := int64(123123)

	walletA := Wallet{}
	walletA.CurrentBalance = ""

	wallets := []Wallet{
		{ID: int64(1), CurrentBalance: "0.12321" , CurrencyName: "BTC"},
		{ID: int64(2), CurrentBalance: "12312.23" , CurrencyName: "ARS"},
	}

	rows := sqlmock.NewRows([]string{"ID","CURRENT_BALANCE", "NAME", "EXPONENT"})
	for _,w := range wallets{
		rows.AddRow(w.ID,w.CurrentBalance, w.CurrencyName, 3)
	}

	dbMockClient.AddExpectedQueryWithRows(getWalletsAndCurrenciesForUser,rows,userID)

	res, err := newDao(dbMockClient).GetWalletsForUser(ctx,userID)
	assert.Nil(t, err)
	for i, walletResponse := range res{
		assert.Equal(t, walletResponse.ID,wallets[i].ID)
		assert.Equal(t, walletResponse.CurrentBalance,wallets[i].CurrentBalance)
		assert.Equal(t, walletResponse.CurrencyName,wallets[i].CurrencyName)
	}

}