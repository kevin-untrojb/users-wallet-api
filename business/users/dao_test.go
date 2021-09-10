package users

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserOK(t *testing.T){
	ctx := context.Background()
	userID := int64(4)
	userMock := user{
		ID : userID,
		Alias: "roberto",
	}
	dbMockClient := mysql.Connect()
	rows := sqlmock.NewRows([]string{"ID", "NAME", "SURNAME","ALIAS", "EMAIL"})
	rows.AddRow(userMock.ID,userMock.Name,userMock.Surname,userMock.Alias,userMock.Email)
	dbMockClient.AddExpectedQueryWithRows(getUserQuery,rows,userID)

	mysql := newDao(dbMockClient)

	userToCompare,err :=mysql.GetUser(ctx,"4")
	assert.Nil(t, err)
	assert.Equal(t, userToCompare.ID,userMock.ID)
	assert.Equal(t, userToCompare.Alias,userMock.Alias)
}

func TestGetUserError(t *testing.T){}

func TestInsertUserOK(t*testing.T){}

func TestInsertUserErrorUserAlreadyExist(t*testing.T){}

