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

	countQuery := "SELECT count(*) FROM transaction t INNER JOIN wallet w ON t.wallet_id = w.id INNER JOIN user u ON u.id = w.user_id INNER JOIN currency c on c.id = w.currency_id WHERE c.name=? AND u.id =? AND t.transaction_type=?"
	totalCount := 10

	searchQuery := "SELECT t.id, t.transaction_type, t.amount, t.date_created, c.name FROM transaction t INNER JOIN wallet w ON t.wallet_id = w.id INNER JOIN user u ON u.id = w.user_id INNER JOIN currency c on c.id = w.currency_id WHERE c.name=? AND u.id =? AND t.transaction_type=? ORDER BY t.date_created DESC LIMIT ? OFFSET ?"
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
		{ID: int64(1), CurrentBalance: "0.12321", CurrencyName: "BTC", CoinExponent: 8},
		{ID: int64(2), CurrentBalance: "12312.23", CurrencyName: "ARS", CoinExponent: 2},
	}

	rows := sqlmock.NewRows([]string{"ID", "CURRENT_BALANCE", "NAME", "EXPONENT"})
	for _, w := range wallets {
		rows.AddRow(w.ID, w.CurrentBalance, w.CurrencyName, w.CoinExponent)
	}

	dbMockClient.AddExpectedQueryWithRows(getWalletsAndCurrenciesForUser, rows, userID)

	res, err := newDao(dbMockClient).GetWalletsForUser(ctx, userID)
	assert.Nil(t, err)
	for i, walletResponse := range res {
		assert.Equal(t, walletResponse.ID, wallets[i].ID)
		assert.Equal(t, walletResponse.CurrentBalance, wallets[i].CurrentBalance)
		assert.Equal(t, walletResponse.CurrencyName, wallets[i].CurrencyName)
		assert.Equal(t, walletResponse.CoinExponent, wallets[i].CoinExponent)
	}

}

func TestNewTransactionOK(t *testing.T) {
	ctx := context.Background()
	dbMockClient := mysql.Connect()
	newTransactionID := int64(99)

	mockTransaction := Transaction{
		WalletID:        int64(1),
		TransactionType: "deposit",
		UserID:          int64(1),
		Amount:          "0.00231000",
	}

	w := Wallet{
		ID:             int64(1),
		CurrentBalance: "1.00000000",
		CurrencyName:   "BTC",
		CoinExponent:   8,
	}
	walletRows := sqlmock.NewRows([]string{"ID", "CURRENT_BALANCE", "NAME", "EXPONENT"})
	walletRows.AddRow(w.ID, w.CurrentBalance, w.CurrencyName, w.CoinExponent)
	dbMockClient.AddExpectedQueryWithRows(getWalletAnCurrencyByWalletID, walletRows, mockTransaction.WalletID)

	dbMockClient.AddExpectedExec(updateBalanceOFAWallet, sqlmock.NewResult(1, 1), "1.00231000", mockTransaction.WalletID)

	dbMockClient.AddExpectedExec(insertTransaction, sqlmock.NewResult(newTransactionID, 1),
		mockTransaction.WalletID, mockTransaction.TransactionType, mockTransaction.Amount)

	res, err := newDao(dbMockClient).NewTransaction(ctx, mockTransaction)
	assert.Nil(t, err)

	assert.Equal(t, res.ID, newTransactionID)

}
