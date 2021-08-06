package error

import "net/http"

type ApiError interface {
	Message() string
	Code() string
	Status() int
	Error() string
}

type apiError struct {
	ErrMsg    string `json:"message"`
	ErrCode   string `json:"error"`
	ErrStatus int    `json:"status"`
}


func (e apiError) Message() string {
	return e.ErrMsg
}

func (e apiError) Status() int {
	return e.ErrStatus
}

func (e apiError) Code() string {
	return e.ErrCode
}
func (e apiError) Error() string {
	return e.ErrCode
}


func NewApiError(message string, error string, status int,) ApiError {
	return apiError{message, error, status}
}
func NewNotFoundApiError(message string) ApiError {
	return apiError{message, "not_found", http.StatusNotFound}
}

func NewBadRequestApiError(message string) ApiError {
	return apiError{message, "bad_request", http.StatusBadRequest,}
}

func NewForbiddenApiError(message string) ApiError {
	return apiError{message, "forbidden", http.StatusForbidden}
}