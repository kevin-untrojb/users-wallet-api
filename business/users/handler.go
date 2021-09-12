package users

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Post(c *gin.Context)
	Get(c *gin.Context)
}

func NewHandler(gtw Gateway) Handler {
	return &handler{
		gtw: gtw,
	}
}

type handler struct {
	gtw Gateway
}

func (h handler) Post(c *gin.Context) {
	ctx := c.Request.Context()
	var u user

	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err := u.ValidateFields(); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	lasInsertedID, err := h.gtw.Create(ctx, u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, NewUserResponse{lasInsertedID})
}

func (h handler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	userID, err := getIDasInt64(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	user, err := h.gtw.Get(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func getIDasInt64(idStr string) (int64, error) {
	if idStr == "" {
		return 0, errors.New("error: bad request")
	}
	return strconv.ParseInt(idStr, 10, 64)
}
