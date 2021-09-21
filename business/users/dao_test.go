package users

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGetUserOK(t *testing.T) {
	ctx := context.Background()
	userID := int64(4)
	userMock := user{
		ID:    4,
		Alias: "roberto",
	}
	dbMockClient := mysql.Connect()
	rows := sqlmock.NewRows([]string{"ID", "NAME", "SURNAME", "ALIAS", "EMAIL"})
	rows.AddRow(userMock.ID, userMock.FirstName, userMock.LastName, userMock.Alias, userMock.Email)
	dbMockClient.AddExpectedQueryWithRows(getUserQuery, rows, userID)

	mysql := newDao(dbMockClient)

	userToCompare, err := mysql.GetUser(ctx, userID)
	assert.Nil(t, err)
	assert.Equal(t, userToCompare.ID, userMock.ID)
	assert.Equal(t, userToCompare.Alias, userMock.Alias)
}

func TestGetUserError(t *testing.T) {
	ctx := context.Background()
	userID := int64(4)

	dbMockClient := mysql.Connect()
	rows := sqlmock.NewRows([]string{"ID", "NAME", "SURNAME", "ALIAS", "EMAIL"}).RowError(0, fmt.Errorf("error"))

	dbMockClient.AddExpectedQueryWithRows(getUserQuery, rows, userID)

	mysql := newDao(dbMockClient)

	_, err := mysql.GetUser(ctx, userID)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "get_user: error getting user")

}

func TestInsertUserOK(t *testing.T) {
	ctx := context.Background()
	insertedID := int64(1)
	userMock := user{
		Email:     "asads@gmail.com",
		Alias:     "robertito",
		FirstName: "roberto",
		LastName:  "asd",
	}
	dbMockClient := mysql.Connect()
	checkRows := sqlmock.NewRows([]string{"COUNT"}).AddRow(0)
	dbMockClient.AddExpectedQueryWithRows(checkUserQuery, checkRows, userMock.Email, userMock.Alias)

	dbMockClient.AddExpectedExec(insertUserQuery, sqlmock.NewResult(insertedID, 1),
		userMock.FirstName, userMock.LastName, userMock.Alias, userMock.Email)
	newUser, err := newDao(dbMockClient).InsertUser(ctx, userMock)
	assert.Nil(t, err)
	assert.Equal(t, newUser.ID, insertedID)
}

func TestInsertUserErrorInsertingUser(t *testing.T) {
	ctx := context.Background()
	userMock := user{
		Email:     "asads@gmail.com",
		Alias:     "robertito",
		FirstName: "roberto",
		LastName:  "asd",
	}
	dbMockClient := mysql.Connect()
	checkRows := sqlmock.NewRows([]string{"COUNT"}).AddRow(0)
	dbMockClient.AddExpectedQueryWithRows(checkUserQuery, checkRows, userMock.Email, userMock.Alias)

	dbMockClient.AddExpectedExecWithError(insertUserQuery, errors.New("db error"),
		userMock.FirstName, userMock.LastName, userMock.Alias, userMock.Email)
	_, err := newDao(dbMockClient).InsertUser(ctx, userMock)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "db error")
}

func TestInsertUserErrorUserAlreadyExist(t *testing.T) {
	ctx := context.Background()
	count := int64(1)
	userMock := user{
		Email: "asads@gmail.com",
		Alias: "roberto",
	}
	dbMockClient := mysql.Connect()
	rows := sqlmock.NewRows([]string{"COUNT"}).AddRow(count)

	dbMockClient.AddExpectedQueryWithRows(checkUserQuery, rows, userMock.Email, userMock.Alias)

	_, err := newDao(dbMockClient).InsertUser(ctx, userMock)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "error, email or alias is already used")
}
