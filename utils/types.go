package utils

import (
	"errors"
	"strconv"
)

func ConvertStringToInt64(idStr string) (int64, error) {
	if idStr == "" {
		return 0, errors.New("error: empty value")
	}
	return strconv.ParseInt(idStr, 10, 64)
}
