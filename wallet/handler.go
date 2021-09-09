package wallet

import "github.com/gin-gonic/gin"

type Handler interface {
	NewMovement(c *gin.Context)
	GetAllMovements(c *gin.Context)
}

type handler struct {
	gtw Gateway
}

func (h handler) NewMovement(c *gin.Context) {
	panic("implement me")
}

func (h handler) GetAllMovements(c *gin.Context) {
	panic("implement me")
}

func NewHandler(gtw Gateway) Handler {
	return &handler{gtw}
}
