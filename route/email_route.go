package route

import (
	"auto-emails/auth"
	"auto-emails/controller"
	"auto-emails/repository"
	"auto-emails/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func EmailRoute(router *gin.Engine, dbMS *gorm.DB, dbMY *gorm.DB, validate *validator.Validate) {

	emailService := service.NewEmailService(
		repository.NewEmailRepository(),
		dbMS,
		dbMY,
		validate,
	)
	emailController := controller.NewEmailController(emailService)

	router.GET("/", auth.Auth(emailController.EmailProsess, []string{}))
}
