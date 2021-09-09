package main

import (
	"os"

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
