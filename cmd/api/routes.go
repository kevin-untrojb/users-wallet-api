package main

import (
	"github.com/gin-gonic/gin"

	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"
)

func routerMapping(router *gin.Engine) {
	dbClient := mysql.Connect()


}