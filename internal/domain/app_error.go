package domain

import (
	"errors"
	"fmt"
)

const (
	ErrCodeDuplicateKey        = "duplicate_key"
	ErrCodeInternalServerError = "internal_server_error"
	ErrCodeInvalidParams       = "invalid_params"
	ErrCodeNotFound            = "not_found"
)

var (
	ErrDuplicateKey = errors.New("duplicate key error")
	ErrIncorrectID  = errors.New("incorrect id error")
	ErrNotFound     = errors.New("record not found error")
)

// AppError is a custom error type that implements the error interface
type AppError struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// NewAppError creates a new AppError with the given code and message.
func NewAppError(code string, msg string) AppError {
	return AppError{
		Code: code,
		Msg:  msg,
	}
}

// Error returns a string representation of the error. It is part of the error interface.
func (e AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Msg)
}
