package wallet

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type SearchRequestParams struct {
	Limit        int
	Offset       int
	MovementType string
	Currency     string
}

func CreateSearchParams(c *gin.Context) (*SearchRequestParams, error) {
	var err error
	var params SearchRequestParams

	limit := c.Query("limit")
	offset := c.Query("offset")
	params.MovementType = c.Query("transaction_type")
	params.Currency = c.Query("currency")

	params.Limit, err = strconv.Atoi(limit)
	if err != nil {
		params.Limit = 10
	}

	if params.Limit > 500 {
		params.Limit = 500
	}

	params.Offset, err = strconv.Atoi(offset)
	if err != nil {
		params.Offset = 0
	}
	return &params, nil
}

type SearchResponse struct {
	Paging  PageInfo      `json:"paging"`
	Results []Transaction `json:"results"`
}
type PageInfo struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
