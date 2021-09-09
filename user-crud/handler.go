package user_crud

import "github.com/gin-gonic/gin"

type Handler interface {
	Post(c *gin.Context)
	Get(c *gin.Context)
}
type handler struct {
	gtw Gateway
}

func (h handler) Post(c *gin.Context) {
	panic("implement me")
}

func (h handler) Get(c *gin.Context) {
	panic("implement me")
}

func NewHandler(gtw Gateway) Handler {
	return &handler{
		gtw: gtw,
	}
}