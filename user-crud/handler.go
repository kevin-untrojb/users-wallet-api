package user_crud

import "github.com/gin-gonic/gin"

type Handler interface {
	Post(c *gin.Context)
	Get(c *gin.Context)
}
