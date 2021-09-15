package main

import (
	"github.com/gin-gonic/gin"
	users2 "github.com/kevin-untrojb/users-wallet-api/business/users"
	wallet2 "github.com/kevin-untrojb/users-wallet-api/business/wallet"
	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
)

func routerMapping(router *gin.Engine) {
	dbClient := mysql.Connect()

	accountGtw := wallet2.NewGateway(dbClient)
	accountHandler := wallet2.NewHandler(accountGtw)

	usersGateway := users2.NewGateway(dbClient, accountGtw)
	usersHandler := users2.NewHandler(usersGateway)

	router.GET("/users/:user_id/account", accountHandler.SearchTransactions)
	router.POST("/users/:user_id/account", accountHandler.NewMovement)

	router.GET("/users/:user_id", usersHandler.Get)
	router.POST("/users", usersHandler.Post)

}
