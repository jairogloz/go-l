package core

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-l/pkg/domain"
	"net/http"
)

const InternalServerErrorMessage = "Ooops! Something went wrong. Please help us by reporting this issue."

// RespondError is a helper function to respond with an error. It checks if the error is an AppError and
// responds with the appropriate status code.
// If the error is not an AppError, it responds with a 500 status code and a generic error message.
func RespondError(c *gin.Context, err error) {
	c.Header("Content-Type", "application/json")
	var appErr domain.AppError
	if errors.As(err, &appErr) {
		if status, ok := ErrCodeMapping[appErr.Code]; ok {
			c.JSON(status, appErr)
			return
		}
	}
	c.JSON(http.StatusInternalServerError, domain.AppError{Code: domain.ErrCodeInternalServerError, Msg: InternalServerErrorMessage})
}

var ErrCodeMapping map[string]int = map[string]int{
	domain.ErrCodeDuplicateKey:  http.StatusConflict,
	domain.ErrCodeNotFound:      http.StatusNotFound,
	domain.ErrCodeInvalidParams: http.StatusBadRequest,
}
