package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandlerPostOK(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()
	userID := int64(1)
	userMock := user{
		Email:     "asads@gmail.com",
		Alias:     "robertito",
		FirstName: "roberto",
		LastName:  "asd",
	}
	userBytes, err := json.Marshal(userMock)
	assert.Nil(t, err)

	gtwMock := NewMockGateway(mockCtrl)
	gtwMock.EXPECT().Create(ctx, userMock).Return(userID, nil)
	httpUsersHandler := NewHandler(gtwMock)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	c.Request, _ = http.NewRequest(http.MethodPost, "/users", strings.NewReader(string(userBytes)))

	httpUsersHandler.Post(c)
	assert.Nil(t, err)
	assert.Equal(t, c.Writer.Status(), http.StatusOK)

}

func TestHandlerGetOK(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()
	userID := int64(1)
	userMock := user{
		ID:        userID,
		Email:     "asads@gmail.com",
		Alias:     "robertito",
		FirstName: "roberto",
		LastName:  "asd",
	}

	gtwMock := NewMockGateway(mockCtrl)
	gtwMock.EXPECT().Get(ctx, userID).Return(userMock, nil)

	httpUsersHandler := NewHandler(gtwMock)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest(http.MethodPost, fmt.Sprintf("/users/%d", userID), strings.NewReader(""))
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "1"}}

	httpUsersHandler.Get(c)
	assert.Equal(t, c.Writer.Status(), http.StatusOK)

}
