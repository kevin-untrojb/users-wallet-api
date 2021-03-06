package users

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevin-untrojb/users-wallet-api/utils"
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
		log.Println("error json format")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err := u.ValidateFields(); err != nil {
		log.Println("error invalid fields")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	newUser, err := h.gtw.Create(ctx, u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func (h handler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	userID, err := utils.ConvertStringToInt64(c.Param("user_id"))
	if err != nil {
		log.Println(fmt.Sprintf("invalid user id %s", c.Param("user_id")))
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
