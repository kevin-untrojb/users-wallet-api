package users

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	ctx := c.Request.Context()
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, errors.New("bad_request error"))
	}
	user, err := h.gtw.Get(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func NewHandler(gtw Gateway) Handler {
	return &handler{
		gtw: gtw,
	}
}
