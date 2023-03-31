package controller

import (
	"auto-emails/auth"

	"github.com/gin-gonic/gin"
)

type EmailController interface {
	EmailProsess(context *gin.Context, auth *auth.AccessDetails)
}
