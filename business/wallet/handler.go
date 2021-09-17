package wallet

import (
	"errors"
	"net/http"

	"github.com/kevin-untrojb/users-wallet-api/utils"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	NewTransaction(c *gin.Context)
	SearchTransactions(c *gin.Context)
}

type handler struct {
	gtw Gateway
}

func NewHandler(gtw Gateway) Handler {
	return &handler{gtw}
}

func (h handler) NewTransaction(c *gin.Context) {
	ctx := c.Request.Context()
	params, err := getTransactionParams(c)
	if err != nil {
		// todo handler
		c.JSON(http.StatusBadRequest, err)
		return
	}

	transactionID, err := h.gtw.NewTransaction(ctx, params)
	if err != nil {
		// todo handler
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, NewTransactionResponse{transactionID})
}

func (h handler) SearchTransactions(c *gin.Context) {
	ctx := c.Request.Context()

	userID, err := utils.ConvertStringToInt64(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("bad_request error"))
		return
	}

	searchParams, err := NewSearchParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := h.gtw.SearchTransactionsForUser(ctx, userID, searchParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func getTransactionParams(c *gin.Context) (Transaction, interface{}) {
	return Transaction{}, nil
}
