package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
	user "github.com/kevin-untrojb/users-wallet-api/user-crud"
)

func routerMapping(router *gin.Engine) {
	dbClient := mysql.Connect()

	userGateway := user.NewGateway(dbClient)
	userHandler := user.NewHandler(userGateway)

	router.GET("", userHandler.Get)
	router.POST("", userHandler.Post)
}
