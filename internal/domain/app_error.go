package domain

import (
	"errors"
	"fmt"
	"log"
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

func ManageError(err error, msg string) error {
	var appErr AppError

	switch {
	case errors.Is(err, ErrDuplicateKey):
		log.Println("duplicate key")
		appErr = AppError{
			Code: ErrCodeDuplicateKey,
			Msg:  "Duplicate key",
		}
	case errors.Is(err, ErrIncorrectID):
		log.Println("incorrect id error")
		appErr = AppError{
			Code: ErrCodeInvalidParams,
			Msg:  "Incorrect id",
		}
	case errors.Is(err, ErrNotFound):
		log.Println("not found error")
		appErr = AppError{
			Code: ErrCodeNotFound,
			Msg:  "Not found",
		}
	default:
		log.Println(err.Error())
		appErr = AppError{
			Code: ErrCodeInternalServerError,
			Msg:  "Server Error",
		}
	}

	// We only add the custom message if the error is not an internal server error
	if msg != "" && appErr.Code != ErrCodeInternalServerError {
		appErr.Msg = fmt.Sprintf("%s: %s", appErr.Msg, msg)
	}
	return appErr
}
