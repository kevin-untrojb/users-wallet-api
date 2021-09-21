package users

import (
	"context"
	"github.com/kevin-untrojb/users-wallet-api/business/wallet"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestCreateUserShouldReturnOK(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()
	newID := int64(1)
	userMock := user{
		Email:     "asads@gmail.com",
		Alias:     "robertito",
		FirstName: "roberto",
		LastName:  "asd",
	}
	mockDao := NewMockMysqlDao(mockCtrl)
	mockDao.EXPECT().InsertUser(ctx, userMock).Return(user{ID:newID}, nil)

	gtw := &gateway{mockDao, nil}
	newUser, err := gtw.Create(ctx, userMock)
	assert.Nil(t, err)
	assert.Equal(t, newUser.ID, newID)
}

func TestGetUserAndWalletShouldReturnOK(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()
	userIDToGet := int64(1)
	userMock := user{
		ID : int64(1),
		Email:     "asads@gmail.com",
		Alias:     "robertito",
		FirstName: "roberto",
		LastName:  "asd",
	}
	mockDao := NewMockMysqlDao(mockCtrl)
	mockDao.EXPECT().GetUser(ctx, userIDToGet).Return(userMock, nil)

	mockWalletGtw:= wallet.NewMockGateway(mockCtrl)
	wallets := []wallet.Wallet{
		{ID: int64(1), CurrentBalance: "0.12321", CurrencyName: "BTC", CoinExponent: 8},
		{ID: int64(2), CurrentBalance: "12312.23", CurrencyName: "ARS", CoinExponent: 2},
	}
	mockWalletGtw.EXPECT().GetWalletsFroUser(ctx,userIDToGet).Return(wallets,nil)
	userMock.Wallets =	wallets
	gtw := &gateway{mockDao,mockWalletGtw}


	user,err := gtw.Get(ctx,userIDToGet)
	assert.Nil(t, err)
	assert.Equal(t, user, userMock)
}
