package users

import (
	"context"
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
	mockDao.EXPECT().InsertUser(ctx, userMock).Return(newID, nil)

	gtw := &gateway{mockDao, nil}
	lastID, err := gtw.Create(ctx, userMock)
	assert.Nil(t, err)
	assert.Equal(t, lastID, newID)
}

func TestGetUserAndWalletShouldReturnOK(t *testing.T) {
}
