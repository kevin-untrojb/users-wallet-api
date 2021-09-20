package main

import (
	"net/http"
	"os"

	"github.com/kevin-untrojb/users-wallet-api/business/users"
	"github.com/kevin-untrojb/users-wallet-api/business/wallet"
	"github.com/kevin-untrojb/users-wallet-api/internal/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := run(port); err != nil {
		panic(err)
	}
}

func run(port string) error {
	router := gin.Default()

	routerMapping(router)
	return router.Run(":" + port)
}

func routerMapping(router *gin.Engine) {
	dbClient := mysql.Connect()

	accountGtw := wallet.NewGateway(dbClient)
	accountHandler := wallet.NewHandler(accountGtw)

	usersGateway := users.NewGateway(dbClient, accountGtw)
	usersHandler := users.NewHandler(usersGateway)

	router.GET("/ping", ping)

	router.POST("/users/:user_id/wallet/:wallet_id/transaction", accountHandler.NewTransaction)
	router.GET("/users/:user_id/wallet", accountHandler.SearchTransactions)


	router.GET("/users/:user_id", usersHandler.Get)
	router.POST("/users", usersHandler.Post)

}
func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
