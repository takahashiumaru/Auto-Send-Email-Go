package controller

import (
	"net/http"

	"auto-emails/auth"
	"auto-emails/model/web"
	"auto-emails/service"

	"github.com/gin-gonic/gin"
)

type EmailControllerImpl struct {
	EmailService service.EmailService
}

func NewEmailController(emailService service.EmailService) EmailController {
	return &EmailControllerImpl{
		EmailService: emailService,
	}
}

func (controller *EmailControllerImpl) EmailProsess(c *gin.Context, auth *auth.AccessDetails) {
	emailResponses := controller.EmailService.EmailProsess(auth)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Proses Email Successfully",
		Data:    emailResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}
