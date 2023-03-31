package exception

import (
	"errors"
	"net/http"
	"strings"

	"auto-emails/helper"
	"auto-emails/model/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(c *gin.Context, err interface{}) {
	if validationError(c, err) {
		return
	}

	if sendToResponseError(c, err) {
		return
	}

	if permissionDeniedError(c, err) {
		return
	}

	if foreignKeyError(c, err) {
		return
	}

	if recordNotFoundError(c, err) {
		return
	}

	if unauthorizedError(c, err) {
		return
	}

	if duplicateError(c, err) {
		return
	}

	internalServerError(c, err)
}

func validationError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		webResponse := web.WebResponse{
			Success: false,
			Message: "Bad Request",
			Data:    helper.ErrorRequestMessage(exception),
		}

		c.JSON(http.StatusBadRequest, webResponse)
		return true
	} else {
		return false
	}
}

func recordNotFoundError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(error)
	if ok && exception.Error() == "record not found" {
		webResponse := web.WebResponse{
			Success: true,
			Message: "Record not found",
		}

		c.JSON(http.StatusOK, webResponse)
		return true
	}
	return false
}

func sendToResponseError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(*ErrorSendToResponse)
	if ok {
		webResponse := web.WebResponse{
			Success: false,
			Message: exception.Error(),
		}

		c.JSON(http.StatusBadRequest, webResponse)
		return true
	}
	return false
}

func unauthorizedError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(error)
	if ok && (errors.Is(exception, ErrUnauthorized) || errors.Is(exception, ErrRefreshTokenExpired)) {
		webResponse := web.WebResponse{
			Success: false,
			Message: exception.Error(),
		}

		c.JSON(http.StatusUnauthorized, webResponse)
		return true
	}
	return false
}

func permissionDeniedError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(error)
	if ok && errors.Is(exception, ErrPermissionDenied) {
		webResponse := web.WebResponse{
			Success: false,
			Message: exception.Error(),
		}

		c.JSON(http.StatusForbidden, webResponse)
		return true
	}
	return false
}

func duplicateError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(error)
	if ok {
		if strings.Contains(exception.Error(), "Error 1062: Duplicate entry") {
			webResponse := web.WebResponse{
				Success: false,
				Message: helper.ErrorDuplicateMessage(exception),
			}

			c.JSON(http.StatusBadRequest, webResponse)
			return true
		}
	}
	return false
}

func foreignKeyError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(error)
	if ok {
		if strings.Contains(exception.Error(), "Error 1452: Cannot add or update a child row") {
			webResponse := web.WebResponse{
				Success: false,
				Message: "A foreign key constraint fails",
			}

			c.JSON(http.StatusBadRequest, webResponse)
			return true
		}
	}
	return false
}

func internalServerError(c *gin.Context, err interface{}) {
	webResponse := web.WebResponse{
		Success: false,
		Message: "Internal Server Error",
	}
	c.JSON(http.StatusInternalServerError, webResponse)
}
