package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func New(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func NewBadRequest(message string, err error) *AppError {
	return New(http.StatusBadRequest, message, err)
}

func NewNotFound(message string, err error) *AppError {
	return New(http.StatusNotFound, message, err)
}

func NewInternalServerError(message string, err error) *AppError {
	return New(http.StatusInternalServerError, message, err)
}

func NewUnauthorized(message string, err error) *AppError {
	return New(http.StatusUnauthorized, message, err)
}