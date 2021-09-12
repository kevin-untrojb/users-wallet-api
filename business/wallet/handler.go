package wallet

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	NewMovement(c *gin.Context)
	SearchTransactions(c *gin.Context)
}

type handler struct {
	gtw Gateway
}

func NewHandler(gtw Gateway) Handler {
	return &handler{gtw}
}

func (h handler) NewMovement(c *gin.Context) {
	panic("implement me")
}

func (h handler) SearchTransactions(c *gin.Context) {
	ctx := c.Request.Context()

	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, errors.New("bad_request error"))
		return
	}

	searchParams, err := NewSearchParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	h.gtw.SearchTransactionsForUser(ctx, userID, searchParams)
}
