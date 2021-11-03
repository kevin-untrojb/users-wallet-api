package wallet

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetWalletsForUserOK(t *testing.T) {
	ctx := context.Background()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockWallets := []Wallet{
		{ID: int64(1), CurrentBalance: "0.12321000", CurrencyName: "BTC", CoinExponent: 8},
		{ID: int64(2), CurrentBalance: "12312.23", CurrencyName: "ARS", CoinExponent: 2},
	}
	for i := range mockWallets {
		mockWallets[i].Coin, _ = newCoin(mockWallets[i].CurrentBalance, mockWallets[i].CoinExponent)
	}

	userID := int64(1)
	mockDao := NewMockMysqlDao(mockCtrl)
	mockDao.EXPECT().GetWalletsForUser(ctx, userID).Return(mockWallets, nil)

	gtw := &gateway{mockDao}

	returnedWallets := []Wallet{
		{ID: int64(1), CurrentBalance: "0.12321000", CurrencyName: "BTC"},
		{ID: int64(2), CurrentBalance: "12312.23", CurrencyName: "ARS"},
	}

	wallets, err := gtw.GetWalletsFroUser(ctx, userID)
	assert.Nil(t, err)
	assert.Equal(t, wallets, returnedWallets)
}
