package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
	"github.com/kevin-untrojb/users-wallet-api/users"
	"github.com/kevin-untrojb/users-wallet-api/wallet"
)

func routerMapping(router *gin.Engine) {
	dbClient := mysql.Connect()

	userGateway := users.NewGateway(dbClient)
	userHandler := users.NewHandler(userGateway)

	accountGtw := wallet.NewGateway(dbClient)
	accountHandler := wallet.NewHandler(accountGtw)

	router.GET("", userHandler.Get)
	router.POST("", userHandler.Post)

	router.GET("", accountHandler.GetAllMovements)
	router.POST("", accountHandler.NewMovement)

}
