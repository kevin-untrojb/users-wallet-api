package wallet

import (
	"errors"
	"fmt"
	"log"
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
	transaction, err := getTransactionParams(c)
	if err != nil {
		log.Println(fmt.Sprintf("error getting params %s", err.Error()))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	newTransaction, err := h.gtw.NewTransaction(ctx, transaction)
	if err != nil {
		log.Println(fmt.Sprintf("error creating new transaction %s", err.Error()))
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, newTransaction)
}

func (h handler) SearchTransactions(c *gin.Context) {
	ctx := c.Request.Context()

	userID, err := utils.ConvertStringToInt64(c.Param("user_id"))
	if err != nil {
		log.Println(fmt.Sprintf("error parsing id %s", c.Param("user_id")))
		c.JSON(http.StatusBadRequest, errors.New("bad_request error"))
		return
	}

	searchParams, err := CreateSearchParams(c)
	if err != nil {
		log.Println(fmt.Sprintf("error creating search params: %s", err.Error()))
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

func getTransactionParams(c *gin.Context) (Transaction, error) {
	var transaction Transaction
	var err error

	if err := c.BindJSON(&transaction); err != nil {
		log.Println(fmt.Sprintf("error json format: %s", err.Error()))
		return transaction, err
	}
	if err := transaction.ValidateFields(); err != nil {
		log.Println(fmt.Sprintf("error invalid body: %s", err.Error()))
		return transaction, err
	}
	transaction.UserID, err = utils.ConvertStringToInt64(c.Param("user_id"))
	if err != nil {
		log.Println(fmt.Sprintf("error parsing user_id: %s", err.Error()))
		return transaction, err
	}
	transaction.WalletID, err = utils.ConvertStringToInt64(c.Param("wallet_id"))
	if err != nil {
		log.Println(fmt.Sprintf("error parsing wallet_id: %s", err.Error()))
		return transaction, err
	}

	return transaction, nil
}
